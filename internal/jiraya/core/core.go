package core

import (
	"github.com/juicyluv/jiraya/internal/jiraya/infrastructure/logging"
	"github.com/juicyluv/jiraya/internal/jiraya/storage"
	"go.uber.org/zap"
)

type Core struct {
	storage storage.Storage
	logger  *zap.Logger
}

func New(st storage.Storage) *Core {
	return &Core{
		storage: st,
		logger:  logging.Get(),
	}
}
