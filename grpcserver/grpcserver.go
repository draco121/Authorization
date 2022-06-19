package server

import (
	"authorization/authorization"
	"authorization/core"
	"context"
)

type Server struct {
	authorization.UnimplementedAuthorizationServer
}

func (s *Server) CreateJWT(ctx context.Context, in *authorization.Createjwtinput) (*authorization.Createjwtoutput, error) {
	repo := core.Core{}
	res, err := repo.CreateJwt(in.Userid)
	if err != nil {
		return nil, err
	} else {
		return &authorization.Createjwtoutput{
			Token: res,
		}, nil
	}
}

func (s *Server) DeleteJWT(ctx context.Context, in *authorization.Deletejwtinput) (*authorization.Deletejwtoutput, error) {
	repo := core.Core{}
	res, err := repo.DeleteJwt(in.Token)
	if err != nil {
		return nil, err
	} else {
		return &authorization.Deletejwtoutput{
			Res: res,
		}, nil
	}
}

func (s *Server) ValidateJWT(ctx context.Context, in *authorization.Validatejwtinput) (*authorization.Validatejwtoutput, error) {
	repo := core.Core{}
	res, err := repo.ValidateJwt(in.Token)
	if err != nil {
		return nil, err
	} else {
		return &authorization.Validatejwtoutput{
			Userid: res,
		}, nil
	}
}
