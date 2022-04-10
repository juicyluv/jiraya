package jiraya

import "github.com/juicyluv/jiraya/internal/jiraya/infrastructure/config"

type App interface {
	GetConfig() *config.Config
	Start() error
}
