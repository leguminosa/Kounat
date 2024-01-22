package character

import (
	"context"

	"github.com/leguminosa/kounat/internal/entity"
	"github.com/leguminosa/kounat/internal/repository"
	"github.com/leguminosa/kounat/internal/storage/database"
)

type RepositoryImpl struct {
	characterDB database.CharacterDB
}

func NewRepository(characterDB database.CharacterDB) repository.CharacterRepository {
	return &RepositoryImpl{
		characterDB: characterDB,
	}
}

func (r *RepositoryImpl) GetByID(ctx context.Context, id int) (*entity.Character, error) {
	return r.characterDB.GetByID(ctx, id)
}
