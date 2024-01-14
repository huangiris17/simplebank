package gapi

import (
	"fmt"

	db "github.com/huangiris17/simplebank/db/sqlc"
	"github.com/huangiris17/simplebank/pb"
	"github.com/huangiris17/simplebank/token"
	"github.com/huangiris17/simplebank/util"
	"github.com/huangiris17/simplebank/worker"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// create new gRPC server
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:          config,
		tokenMaker:      tokenMaker,
		store:           store,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
