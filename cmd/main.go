package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"github.com/codebyaadi/rss-feed-agg/config"
	"github.com/codebyaadi/rss-feed-agg/internal/database"
	"github.com/codebyaadi/rss-feed-agg/internal/handlers"
	"github.com/codebyaadi/rss-feed-agg/internal/redis"
	"github.com/codebyaadi/rss-feed-agg/internal/rss"
	"github.com/codebyaadi/rss-feed-agg/internal/utils"
	"github.com/codebyaadi/rss-feed-agg/pkg/logger"
)

func main() {
	if err := logger.New(); err != nil {
		log.Fatal("Error initializing logger", err)
	}
	defer logger.Logger.Sync()
	logger.Logger.Info("starting server...")

	if err := godotenv.Load(); err != nil {
		logger.Logger.Warn("Error loading .env file", zap.Error(err))
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		logger.Logger.Warn("defaulting to port", zap.String("port", port))
	}

	dbUrl := os.Getenv("POSTGRES_URL")
	if dbUrl == "" {
		logger.Logger.Fatal("POSTGRES_URL must be set")
	}

	if err := redis.InitRedis(); err != nil {
		logger.Logger.Fatal("Can't connect to Redis", zap.Error(err))
	}
	defer redis.CloseRedis()

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		logger.Logger.Fatal("can't connect to postgres database %v", zap.Error(err))
	}
	defer conn.Close()

	db := database.New(conn)
	apiCfg := &config.ApiConfig{
		DB: db,
	}

	go rss.RSSFeedScrapper(db, 10, time.Minute)

	handler := &handlers.Handler{ApiConfig: apiCfg}

	if err := conn.Ping(); err != nil {
		logger.Logger.Fatal("can't reach postgres database %v", zap.Error(err))
	}

	logger.Logger.Info("successfully connected to database")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlerHealth)
	mux.HandleFunc("GET /error", handlerErr)
	mux.HandleFunc(("POST /users/create"), handler.CreateUser)
	mux.HandleFunc("POST /users/login", handler.LoginUser)
	mux.HandleFunc(("GET /users"), handler.AuthMiddleware(handler.GetUserByAPIKey))
	mux.HandleFunc(("POST /feeds/create"), handler.AuthMiddleware(handler.CreateFeed))
	mux.HandleFunc(("GET /feeds"), handler.GetAllFeeds)
	mux.HandleFunc(("POST /feeds/follow"), handler.AuthMiddleware(handler.CreateFeedFollow))
	mux.HandleFunc(("GET /feeds/follow"), handler.AuthMiddleware(handler.GetAllFeedFollows))
	mux.HandleFunc(("DELETE /feeds/follow/{feedFollowID}"), handler.AuthMiddleware(handler.DeleteFeedFollow))
	mux.HandleFunc(("GET /posts"), handler.AuthMiddleware(handler.GetPostsForUser))

	addr := ":" + port
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		logger.Logger.Info("listening on port %s", zap.String("port", port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Fatal("could not listen on port", zap.String("port", port), zap.Error(err))
		}
	}()

	<-shutdownCh
	logger.Logger.Info("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Logger.Fatal("could not gracefully shutdown the server", zap.Error(err))
	}
	logger.Logger.Info("Server stopped")
}

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}

func handlerErr(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusInternalServerError, "something went wrong")
}
