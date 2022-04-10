package jiraya

import "github.com/juicyluv/jiraya/internal/jiraya/infrastructure/config"

type App interface {
	Config() *config.Config
	Start() error
}
