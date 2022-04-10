package app

import (
	"github.com/juicyluv/jiraya/internal/jiraya"
	"github.com/juicyluv/jiraya/internal/jiraya/infrastructure/config"
	"github.com/juicyluv/jiraya/internal/jiraya/infrastructure/logging"
	"github.com/juicyluv/jiraya/internal/jiraya/storage"
	"go.uber.org/zap"
)

type app struct {
	cfg     *config.Config
	storage storage.Storage
	logger  *zap.Logger
}

func New() jiraya.App {
	cfg := config.Get()

	logger := logging.Get()

	return &app{
		cfg:    cfg,
		logger: logger,
	}
}

func (a *app) Config() *config.Config {
	return a.cfg
}

func (a *app) SetConfig(cfg *config.Config) {
	a.cfg = cfg
}

func (a *app) Storage() storage.Storage {
	return a.storage
}
