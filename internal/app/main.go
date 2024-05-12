package app

import (
	"context"
	"sync"
	"time"

	"github.com/IBM/sarama"

	_ "github.com/gvidow/go-chat-transport/docs"
	"github.com/gvidow/go-chat-transport/internal/api"
	http "github.com/gvidow/go-chat-transport/internal/pkg/delivery/http/v1"
	consumer "github.com/gvidow/go-chat-transport/internal/pkg/delivery/kafka"
	manager "github.com/gvidow/go-chat-transport/internal/pkg/kafka"
	msgRepo "github.com/gvidow/go-chat-transport/internal/pkg/repository/message"
	repository "github.com/gvidow/go-chat-transport/internal/pkg/repository/segment"
	"github.com/gvidow/go-chat-transport/internal/pkg/repository/user"
	"github.com/gvidow/go-chat-transport/internal/pkg/usecase/message"
	"github.com/gvidow/go-chat-transport/internal/pkg/usecase/segment"
	"github.com/gvidow/go-chat-transport/internal/server"
	"github.com/gvidow/go-chat-transport/pkg/errors"
	"github.com/gvidow/go-chat-transport/pkg/logger"
)

const _addr = "0.0.0.0:5500"
const _appTopic = "chat"

const _lruSize = 1 << 10

var (
	_timeoutSenderMessage        = 2 * time.Second
	_sizeSegment          uint32 = 110
)

var _kafkaAddrs = []string{"localhost:9092"}

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

	kafkaCfg := sarama.NewConfig()
	kafkaCfg.Producer.Return.Successes = true
	kafkaManager, err := manager.NewKafkaManager(_kafkaAddrs, _appTopic, kafkaCfg)
	if err != nil {
		return errors.WrapError(err, "create kafka manager")
	}
	defer kafkaManager.Close()

	userRepo, err := user.NewUsernameStory(_lruSize)
	if err != nil {
		return errors.WrapError(err, "create user repository")
	}

	segmentStackerRepo, err := repository.NewSegmentStacker(kafkaManager)
	if err != nil {
		return errors.WrapError(err, "new repository segmant stacker")
	}

	segRepo := repository.NewRepository(segmentStackerRepo, repository.NewSegmentTransfer(_apiEncodingServer))

	su := segment.NewSegmentUsecase(log, segRepo, userRepo, msgRepo.NewSenderMsg(_apiWSServer))

	if err := api.RegistryHandler(serv, http.NewHandler(
		log,
		su,
		message.NewUsecaseMessage(segRepo, userRepo, message.WithPartitionBySize(_sizeSegment)),
	)); err != nil {
		log.Error(err.Error())
		return err
	}

	consumerHub, err := consumer.NewConsumerHub(ctx, &consumer.Config{
		Manager:              kafkaManager,
		Log:                  log,
		TimeoutSenderMessage: _timeoutSenderMessage,
	}, su)
	if err != nil {
		return errors.WrapError(err, "new consumer hub")
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		wg.Done()
		log.Info("start consumer hub")
		if err := consumerHub.Serve(); err != nil {
			log.Error(err.Error())
		}
	}()

	defer func() {
		consumerHub.Shutdown()
		wg.Wait()
	}()

	log.Sugar().Infof("run server on %s", _addr)
	if err := serv.Run(_addr); err != nil {
		log.Sugar().Errorf("server stop with error: %v", err)
		return err
	}
	log.Info("server stop")

	return nil
}
