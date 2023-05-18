package createharvest

import (
	"errors"
	"fmt"

	"github.com/jendrusha/harvester/pkg/model"
)

var UnableToPersistHarvest = errors.New("createHarvestHandler - unable to persist harvest")

type createHarvestStorage interface {
	PersistHarvest(model.Harvest) error
}

type createHarvestHandler struct {
	storage createHarvestStorage
}

func NewHandler(storage createHarvestStorage) *createHarvestHandler {
	return &createHarvestHandler{
		storage: storage,
	}
}

func (uc createHarvestHandler) Handle(req CreateHarvestRequest) (any, error) {
	if err := uc.storage.PersistHarvest(model.Harvest{}); err != nil {
		return nil, fmt.Errorf("%w: %w", UnableToPersistHarvest, err)
	}

	return map[string]string{
		"test": "done",
	}, nil
}
