package database

import (
	"context"

	"github.com/leguminosa/kounat/internal/entity"
)

type CharacterDB interface {
	GetByID(ctx context.Context, id int) (*entity.Character, error)
}
