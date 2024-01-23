package character

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/leguminosa/kounat/internal/entity"
	"github.com/leguminosa/kounat/internal/storage/database"
	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	mockDB := &database.MockCharacterDB{}
	assert.NotEmpty(t, NewRepository(mockDB))
}

func TestRepositoryImpl_GetByID(t *testing.T) {
	ctx := context.Background()
	r := &RepositoryImpl{}
	tests := []struct {
		name    string
		id      int
		prepare func(m *database.MockCharacterDB)
		want    *entity.Character
		wantErr bool
	}{
		{
			name: "error get by id",
			id:   1,
			prepare: func(m *database.MockCharacterDB) {
				m.EXPECT().GetByID(ctx, 1).Return(nil, assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "success",
			id:   1,
			prepare: func(m *database.MockCharacterDB) {
				m.EXPECT().GetByID(ctx, 1).Return(&entity.Character{
					ID: 1,
				}, nil)
			},
			want: &entity.Character{
				ID: 1,
			},
			wantErr: false,
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := database.NewMockCharacterDB(ctrl)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prepare != nil {
				tt.prepare(mockDB)
			}
			r.characterDB = mockDB

			got, err := r.GetByID(ctx, tt.id)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
