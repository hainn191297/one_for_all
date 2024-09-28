package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"one_for_all/services/account/api/internal/config"
	"one_for_all/services/account/api/internal/middleware"
)

type ServiceContext struct {
	Config              config.Config
	UserTokenMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		UserTokenMiddleware: middleware.NewUserTokenMiddleware().Handle,
	}
}
