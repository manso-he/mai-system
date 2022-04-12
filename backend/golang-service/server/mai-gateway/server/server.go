package server

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"sync"

	"manso.live/backend/golang-service/pkg/constant"
	"manso.live/backend/golang-service/pkg/log"
	"manso.live/backend/golang-service/server/mai-gateway/config"
	"manso.live/backend/golang-service/server/mai-gateway/http"
)

func Run() error {
	cfg, err := config.NewConfig(constant.DefaultConfigFile)
	if err != nil {
		return errors.Wrap(err, constant.InitConfigFailed)
	}

	defer log.Sync()
	if log.IsDebug() {
		log.Warn("You are currently in DEBUG mode, please DO NOT leak sensitive data.")
		log.Debug("Configuration specification", zap.Any("configutil", *cfg))
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		http.StartHttpServer(cfg)
	}()

	wg.Wait()
	log.Info("mai-gateway terminated")

	return nil
}
