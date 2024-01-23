package postgres

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/leguminosa/kounat/internal/entity"
	"github.com/leguminosa/kounat/internal/tools"
	"github.com/stretchr/testify/assert"
)

func TestNewCharacterDB(t *testing.T) {
	mockPgx := &tools.MockPGXClient{}
	assert.NotEmpty(t, NewCharacterDB(mockPgx))
}

func TestCharacterDB_GetByID(t *testing.T) {
	ctx := context.Background()
	s := &CharacterDB{}
	tests := []struct {
		name    string
		id      int
		prepare func(m *tools.MockPGXClient)
		want    *entity.Character
		wantErr bool
	}{
		{
			name: "error get slave",
			id:   1,
			prepare: func(m *tools.MockPGXClient) {
				m.EXPECT().GetSlave(ctx).Return(nil, assert.AnError)
			},
			wantErr: true,
		},
	}
	ctrl := gomock.NewController(t)
	mockPgx := tools.NewMockPGXClient(ctrl)
	defer ctrl.Finish()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prepare != nil {
				tt.prepare(mockPgx)
			}
			s.client = mockPgx

			got, err := s.GetByID(ctx, tt.id)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
