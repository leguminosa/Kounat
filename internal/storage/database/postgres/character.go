package postgres

import (
	"context"

	"github.com/leguminosa/kounat/internal/entity"
	"github.com/leguminosa/kounat/internal/storage/database"
	"github.com/leguminosa/kounat/internal/tools"
)

type CharacterDB struct {
	client tools.PGXClient
}

func NewCharacterDB(client tools.PGXClient) database.CharacterDB {
	return &CharacterDB{
		client: client,
	}
}

func (s *CharacterDB) GetByID(ctx context.Context, id int) (*entity.Character, error) {
	db, err := s.client.GetSlave(ctx)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT id, name, created_at
		FROM character
		WHERE id = $1
	`

	var result entity.Character
	err = db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&result.ID,
		&result.Name,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
