package protocol

import (
	"context"
	"daveslist/config"
	"daveslist/internal/handler/httphdl"
	"daveslist/middleware"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

	logger "daveslist/pkg/go-logger"
)

func ServeREST() error {
	e := echo.New()
	e.HTTPErrorHandler = middleware.ErrorHandler

	httpHdl := httphdl.NewHTTP(httphdl.Config{
		Validator:       app.pkg.validator,
		CategoryService: app.svc.categorySvc,
		ListingService:  app.svc.listingSvc,
		MessageService:  app.svc.messageSvc,
	})
	e.GET("/healthcheck", httpHdl.HealthCheck)
	apiGroup := e.Group("/api/v1")
	{
		apiGroup.Use(
			echomiddleware.CORS(),
			middleware.Logger(logger.GetLogger()),
			middleware.Auth(app.svc.authSvc),
		)

		categoryGroup := apiGroup.Group("/category")
		{
			categoryGroup.POST("", httpHdl.CreateCategory)
			categoryGroup.GET("", httpHdl.GetCategoryList)
			categoryGroup.DELETE("", httpHdl.DeleteCategory)
		}

		listingGroup := apiGroup.Group("/listing")
		{
			listingGroup.POST("", httpHdl.CreateListing)
			listingGroup.GET("", httpHdl.GetListingList)
			listingGroup.PUT("", httpHdl.UpdateListing)
			listingGroup.DELETE("", httpHdl.DeleteListing)
			listingGroup.PATCH("/:listing_id/hide", httpHdl.HideListing)
			listingGroup.POST("/:listing_id/reply", httpHdl.CreateReplyListing)
			listingGroup.GET("/:listing_id/reply", httpHdl.GetReplyListingList)
		}

		messageGroup := apiGroup.Group("/message")
		{
			messageGroup.POST("", httpHdl.CreateMessage)
			messageGroup.GET("", httpHdl.GetMessageList)
		}
	}

	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf("%s:%s", config.GetConfig().Server.Host, config.GetConfig().Server.Port)); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("Gracefully shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
