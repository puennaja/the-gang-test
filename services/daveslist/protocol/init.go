package protocol

import (
	"daveslist/infrastructure"
	"daveslist/internal/core/port"

	"daveslist/config"
	"daveslist/internal/repository/mongorepo"

	authSrv "daveslist/internal/core/service/auth"
	categorySrv "daveslist/internal/core/service/category"
	listingSrv "daveslist/internal/core/service/listing"
	messageSrv "daveslist/internal/core/service/message"

	logger "daveslist/pkg/go-logger"
	configLog "daveslist/pkg/go-logger/config"
	coreLog "daveslist/pkg/go-logger/core"

	errors "daveslist/pkg/go-errors"

	"daveslist/pkg/validator"
)

var app *application

type application struct {
	svc services
	pkg packages
}

type services struct {
	authSvc     port.AuthService
	categorySvc port.CategoryService
	listingSvc  port.ListingService
	messageSvc  port.MessageService
}

type packages struct {
	validator validator.Validator
}

func init() {
	startApp()
}

func startApp() {
	// Setup
	config.Init()
	cfg := config.GetConfig()

	errors.Init()

	initiateFeild := map[string]interface{}{
		"service": cfg.AppConfig.Service,
		"version": cfg.AppConfig.Version,
	}
	logCfg := configLog.NewDefaultConfig()
	logCfg.SetDisableStacktrace(true)
	logCfg.SetDisableCaller(true)
	logCfg.SetInitialFields(initiateFeild)
	baseLog := coreLog.BuildDefaultBaseLog(logCfg)
	logger.InitLogger(baseLog)

	// infrastructure
	var (
		mongoConn = infrastructure.InitMongo()
	)

	// packages
	var (
		trans     = validator.NewTranslator()
		validator = validator.New(trans)
		packages  = packages{
			validator: validator,
		}
	)

	// adapters
	var ()

	// repositories
	var (
		categoryRepo     = mongorepo.NewCategoryRepository(mongoConn, cfg.MongoDB.Database)
		listingRepo      = mongorepo.NewListingRepository(mongoConn, cfg.MongoDB.Database)
		replyListingRepo = mongorepo.NewReplyListingRepository(mongoConn, cfg.MongoDB.Database)
		messageRepo      = mongorepo.NewMessageRepository(mongoConn, cfg.MongoDB.Database)
	)

	// service
	var (
		authSvc = authSrv.New(&authSrv.Config{
			ModelPath:  cfg.Service.AuthSvc.ModelPath,
			PolicyPath: cfg.Service.AuthSvc.PolicyPath,
		})
		listingSvc = listingSrv.New(&listingSrv.Config{
			AuthService:      authSvc,
			ListingRepo:      listingRepo,
			ReplyListingRepo: replyListingRepo,
		})
		categorySvc = categorySrv.New(&categorySrv.Config{
			CategoryRepo:   categoryRepo,
			ListingService: listingSvc,
		})
		messageSvc = messageSrv.New(&messageSrv.Config{
			MessageRepo: messageRepo,
		})
	)

	app = &application{
		svc: services{
			authSvc:     authSvc,
			listingSvc:  listingSvc,
			categorySvc: categorySvc,
			messageSvc:  messageSvc,
		},
		pkg: packages,
	}
}
