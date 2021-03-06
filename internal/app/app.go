package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DePavelPo/go_training_web_with_gin/internal/handler"
	"github.com/DePavelPo/go_training_web_with_gin/internal/repository"
	"github.com/DePavelPo/go_training_web_with_gin/internal/service"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/pressly/goose"
)

func Run() {

	var router *gin.Engine

	db, err := goose.OpenDBWithDriver("postgres", "postgres://pahan:pahan123@localhost:5432/training_gin?sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	if err := goose.Run("up", db, "migrations", ""); err != nil {
		log.Fatal(err)
	}

	client := repository.NewClient()

	repository := repository.NewRepository(client)

	service := service.NewService(repository)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("github.com/DePavelPo/go_training_web_with_gin/templates/*")

	handler.NewHandler(service, router)

	// Define the route for the index page and display the index.html template
	// To start with, we'll use an inline route handler. Later on, we'll create
	// standalone functions that will be used as route handlers.

	// Start serving the application
	err = router.Run()
	if err != nil {
		log.Fatal(err.Error())
	}

	srv := &http.Server{
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

}
