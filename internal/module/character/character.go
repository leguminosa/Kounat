package character

import (
	"context"

	"github.com/leguminosa/kounat/internal/entity"
	"github.com/leguminosa/kounat/internal/module"
	"github.com/leguminosa/kounat/internal/repository"
)

type ModuleImpl struct {
	characterRepo repository.CharacterRepository
}

func NewModule(characterRepo repository.CharacterRepository) module.CharacterModule {
	return &ModuleImpl{
		characterRepo: characterRepo,
	}
}

func (m *ModuleImpl) GetByID(ctx context.Context, id int) (*entity.Character, error) {
	return m.characterRepo.GetByID(ctx, id)
}
