package server

import (
	"context"

	warhammerv1 "github.com/brittonhayes/warhammer/proto/gen/warhammer/sigmar/v1"
)

type armyServer struct {
	warhammerv1.ArmyServiceServer
}

func (s armyServer) CreateUnit(ctx context.Context, req *warhammerv1.CreateUnitRequest) (*warhammerv1.CreateUnitResponse, error) {
	log.Info().Msg("message received")
	return &warhammerv1.CreateUnitResponse{}, nil
}
