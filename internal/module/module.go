package module

import (
	"context"

	"github.com/leguminosa/kounat/internal/entity"
)

type CharacterModule interface {
	GetByID(ctx context.Context, id int) (*entity.Character, error)
}
