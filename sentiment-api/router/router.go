package router

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"sentiment-api/internal/sentiment"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var r *gin.Engine

type RouterConfig struct {
	// Add future handlers here, e.g.:
	SentimentHandler *sentiment.SentimentHandler
}

func InitRouter(cfg *RouterConfig) {
	r = gin.Default()

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:3001", "http://127.0.0.1:5500"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"Content-Type", "Authorization", "Range", "Content-Range"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		//time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	// Add Swagger only in local development mode
	if os.Getenv("IS_PRODUCTION") == "NO" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	sentimentRoutes := r.Group("/api/sentiment")
	{
		sentimentRoutes.POST("/", cfg.SentimentHandler.HandleSentimentPost)
	}

	log.Println("Server listening on http://localhost:8080")

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 1 second.

	<-ctx.Done()
	log.Println("timeout of 1 second")

	log.Println("Server exiting")
}
