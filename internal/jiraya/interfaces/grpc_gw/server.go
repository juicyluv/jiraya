package grpc_gw

import (
	"github.com/juicyluv/jiraya/internal/jiraya/core"
	"github.com/juicyluv/jiraya/internal/jiraya/interfaces/grpc_gw/protobuf"
)

type server struct {
	protobuf.UnimplementedJirayaServer
	core *core.Core
}

// New returns a new grpc gateway server instance.
func New(core *core.Core) *server {
	return &server{core: core}
}

func (s *server) Core() *core.Core {
	return s.core
}
