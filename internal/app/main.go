package app

import (
	"context"

	_ "github.com/gvidow/go-chat-transport/docs"
	"github.com/gvidow/go-chat-transport/internal/api"
	http "github.com/gvidow/go-chat-transport/internal/pkg/delivery/http/v1"
	"github.com/gvidow/go-chat-transport/internal/server"
	"github.com/gvidow/go-chat-transport/pkg/logger"
)

const _addr = "0.0.0.0:5500"

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server chat server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      pinspire.site:5500
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func Main(ctx context.Context, log *logger.Logger) error {
	serv := server.NewServer()

	if err := api.RegistryHandler(serv, http.NewHandler()); err != nil {
		log.Error(err.Error())
		return err
	}

	log.Sugar().Infof("run server on %s", _addr)
	if err := serv.Run(_addr); err != nil {
		log.Sugar().Errorf("server stop with error: %v", err)
		return err
	}
	log.Info("server stop")

	return nil
}
