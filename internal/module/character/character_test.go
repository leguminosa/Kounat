package character

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/leguminosa/kounat/internal/entity"
	"github.com/leguminosa/kounat/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestNewModule(t *testing.T) {
	mockRepo := &repository.MockCharacterRepository{}
	assert.NotEmpty(t, NewModule(mockRepo))
}

func TestModuleImpl_GetByID(t *testing.T) {
	ctx := context.Background()
	m := &ModuleImpl{}
	tests := []struct {
		name    string
		id      int
		prepare func(m *repository.MockCharacterRepository)
		want    *entity.Character
		wantErr bool
	}{
		{
			name: "error get by id",
			id:   1,
			prepare: func(m *repository.MockCharacterRepository) {
				m.EXPECT().GetByID(ctx, 1).Return(nil, assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "success",
			id:   1,
			prepare: func(m *repository.MockCharacterRepository) {
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
	mockRepo := repository.NewMockCharacterRepository(ctrl)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prepare != nil {
				tt.prepare(mockRepo)
			}
			m.characterRepo = mockRepo

			got, err := m.GetByID(ctx, tt.id)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
