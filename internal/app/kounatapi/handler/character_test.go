package handler

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/leguminosa/kounat/internal/entity"
	"github.com/leguminosa/kounat/internal/module"
	"github.com/stretchr/testify/assert"
)

func TestNewCharacter(t *testing.T) {
	mockModule := &module.MockCharacterModule{}
	assert.NotEmpty(t, NewCharacter(mockModule))
}

func TestCharacterHandler_GetByID(t *testing.T) {
	h := &CharacterHandler{}
	tests := []struct {
		name    string
		mockCtx *mockEchoContext
		prepare func(m *module.MockCharacterModule)
		want    string
		wantErr bool
	}{
		{
			name: "error get by id",
			prepare: func(m *module.MockCharacterModule) {
				m.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(nil, assert.AnError)
			},
			want:    "{\"message\":\"something went wrong\"}\n",
			wantErr: false,
		},
		{
			name: "success",
			prepare: func(m *module.MockCharacterModule) {
				m.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(&entity.Character{
					ID:   1,
					Name: "Elesis",
				}, nil)
			},
			want:    "{\"id\":1,\"name\":\"Elesis\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\"}\n",
			wantErr: false,
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockModule := module.NewMockCharacterModule(ctrl)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newMockEchoContext(nil)

			if tt.prepare != nil {
				tt.prepare(mockModule)
			}
			h.module = mockModule

			err := h.GetByID(c)
			if !assert.Equal(t, tt.wantErr, err != nil) {
				return
			}

			got := c.getResponseBody()
			assert.Equal(t, tt.want, got)
		})
	}
}
