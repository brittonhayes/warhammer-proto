package server

import (
	"net"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.

	"github.com/brittonhayes/warhammer/internal/logger"
	warhammerv1 "github.com/brittonhayes/warhammer/proto/gen/warhammer/sigmar/v1"
	"github.com/diamondburned/arikawa/v3/session"
	"google.golang.org/grpc"
)

var log = logger.NewWithPrefix("role", "gRPC")

type Server struct {
	session *session.Session
	host    string
}

func New(host string) *Server {
	return &Server{
		host: host,
	}
}

func (s *Server) Listen() {
	defer s.session.Close()

	log.Info().Msgf("GRPC server is up on %s 🚀", s.host)
	listener, err := net.Listen("tcp", s.host)
	if err != nil {
		log.Fatal().Msgf("failed to listen on %s: %w", s.host, err)
	}

	server := grpc.NewServer()
	warhammerv1.RegisterArmyServiceServer(server, armyServer{})

	if err := server.Serve(listener); err != nil {
		log.Fatal().Msgf("failed to serve gRPC server: %w", err)
	}
}
