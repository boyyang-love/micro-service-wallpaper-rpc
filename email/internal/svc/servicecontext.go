package svc

import "github.com/boyyang-love/micro-service-wallpaper-rpc/email/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
