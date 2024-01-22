package repository

import (
	"context"

	"github.com/leguminosa/kounat/internal/entity"
)

type CharacterRepository interface {
	GetByID(ctx context.Context, id int) (*entity.Character, error)
}
