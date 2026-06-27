package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/swaroopt14/sbi-gff-hackthon/backend/internal/config"
	"github.com/swaroopt14/sbi-gff-hackthon/backend/internal/logger"
	"github.com/swaroopt14/sbi-gff-hackthon/backend/internal/middleware"
	"github.com/swaroopt14/sbi-gff-hackthon/backend/pkg/database"
)

const version = "0.1.0"

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	if err := logger.Init(cfg.App.LogLevel); err != nil {
		panic("failed to init logger: " + err.Error())
	}
	defer logger.Sync()

	log := logger.Get()
	log.Info("starting Aperture backend",
		zap.String("version", version),
		zap.String("env", cfg.App.Env),
		zap.String("port", cfg.App.Port),
	)

	db, err := database.NewPostgres(cfg.Database)
	if err != nil {
		log.Warn("database not available — starting without DB (dev mode)", zap.Error(err))
	} else {
		log.Info("database connected")
		defer db.Close()
	}

	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestID())
	router.Use(middleware.Logger())

	// Public routes
	router.GET("/health", func(c *gin.Context) {
		dbStatus := "unavailable"
		if db != nil {
			if err := db.PingContext(c.Request.Context()); err == nil {
				dbStatus = "ok"
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"version":   version,
			"env":       cfg.App.Env,
			"database":  dbStatus,
			"request_id": c.GetString(middleware.ContextKeyRequestID),
		})
	})

	router.GET("/ready", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ready"})
	})

	// API v1 group — handlers wired here in subsequent phases
	v1 := router.Group("/api/v1")
	_ = v1 // handlers registered in Phase 15+

	srv := &http.Server{
		Addr:         ":" + cfg.App.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Info("server listening", zap.String("addr", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("server failed", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("server forced to shutdown", zap.Error(err))
	}
	log.Info("server exited")
}
