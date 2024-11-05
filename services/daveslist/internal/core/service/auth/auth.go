package auth

import (
	"daveslist/internal/core/port"
	logger "daveslist/pkg/go-logger"

	"github.com/casbin/casbin/v2"
)

type Config struct {
	ModelPath  string
	PolicyPath string
}

type Service struct {
	casbinEnforcer casbin.IEnforcer
}

func New(cfg *Config) port.AuthService {
	enforcer, err := casbin.NewEnforcer(cfg.ModelPath, cfg.PolicyPath)
	if err != nil {
		logger.PanicW("Auth Service [New]:", err.Error())
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		logger.PanicW("Auth Service [New]:", err.Error())
	}
	return &Service{
		casbinEnforcer: enforcer,
	}
}
