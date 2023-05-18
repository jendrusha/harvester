package storage

import (
	"fmt"

	"github.com/jendrusha/harvester/pkg/core"
	"github.com/jendrusha/harvester/pkg/model"
)

type storageMySQL struct{}

var _ core.Storage = &storageMySQL{}

func NewMySQL() (*storageMySQL, error) {
	return &storageMySQL{}, nil
}

func (s *storageMySQL) PersistHarvest(harvest model.Harvest) error {
	fmt.Println("harvest added to storage")

	return nil
}
