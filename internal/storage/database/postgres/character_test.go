package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5"
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
		name          string
		id            int
		prepareClient func(m *tools.MockPGXClient, pool *tools.MockPGXPool)
		preparePool   func(m *tools.MockPGXPool)
		want          *entity.Character
		wantErr       bool
	}{
		{
			name: "error get slave",
			id:   1,
			prepareClient: func(m *tools.MockPGXClient, pool *tools.MockPGXPool) {
				m.EXPECT().GetSlave(ctx).Return(nil, assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "error scan",
			id:   1,
			prepareClient: func(m *tools.MockPGXClient, pool *tools.MockPGXPool) {
				m.EXPECT().GetSlave(ctx).Return(pool, nil)
			},
			preparePool: func(m *tools.MockPGXPool) {
				mockRow := tools.NewMockPGXRow(
					[]string{
						"id",
						"name",
						"created_at",
					},
					[]interface{}{
						"invalid id",
						"testname",
						time.Time{},
					},
				)
				m.QueryRowFunc = func(ctx context.Context, query string, args ...interface{}) pgx.Row {
					return mockRow
				}
			},
			wantErr: true,
		},
		{
			name: "success",
			id:   1,
			prepareClient: func(m *tools.MockPGXClient, pool *tools.MockPGXPool) {
				m.EXPECT().GetSlave(ctx).Return(pool, nil)
			},
			preparePool: func(m *tools.MockPGXPool) {
				mockRow := tools.NewMockPGXRow(
					[]string{
						"id",
						"name",
						"created_at",
					},
					[]interface{}{
						1,
						"testname",
						time.Time{},
					},
				)
				m.QueryRowFunc = func(ctx context.Context, query string, args ...interface{}) pgx.Row {
					return mockRow
				}
			},
			want: &entity.Character{
				ID:        1,
				Name:      "testname",
				CreatedAt: time.Time{},
			},
			wantErr: false,
		},
	}
	ctrl := gomock.NewController(t)
	mockPgx := tools.NewMockPGXClient(ctrl)
	mockPool := &tools.MockPGXPool{}
	defer ctrl.Finish()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prepareClient != nil {
				tt.prepareClient(mockPgx, mockPool)
			}
			s.client = mockPgx

			if tt.preparePool != nil {
				tt.preparePool(mockPool)
			}

			got, err := s.GetByID(ctx, tt.id)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
