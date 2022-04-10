package grpc_gw

import (
	"github.com/juicyluv/jiraya/internal/jiraya/interfaces/grpc_gw/protobuf"
	"github.com/juicyluv/jiraya/internal/jiraya/storage"
)

type Server struct {
	protobuf.UnimplementedJirayaServer
	storage storage.Storage
}

// New returns a new Server instance.
func New(st storage.Storage) *Server {
	return &Server{storage: st}
}

func (s *Server) GetStorage() storage.Storage {
	return s.storage
}
