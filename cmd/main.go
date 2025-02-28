package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	ginRounter := gin.Default()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	ginRounter.GET("/ping", func(ginCtx *gin.Context) {
		ginCtx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: ginRounter,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server cannot start", err)
	}

	logger.Info(fmt.Sprintf("server is starting %s", server.Addr))
}
