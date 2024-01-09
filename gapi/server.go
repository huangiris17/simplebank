package gapi

import (
	"fmt"

	db "github.com/huangiris17/simplebank/db/sqlc"
	"github.com/huangiris17/simplebank/pb"
	"github.com/huangiris17/simplebank/token"
	"github.com/huangiris17/simplebank/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// create new gRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
		store:      store}

	return server, nil
}
