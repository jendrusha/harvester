package core

import "github.com/jendrusha/harvester/pkg/model"

type Storage interface {
	PersistHarvest(model.Harvest) error
}
