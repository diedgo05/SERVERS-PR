package replica

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllTrucosReplica(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, bdReplica)
}

func Run() {
	s := gin.Default()

	s.GET("/truco", GetAllTrucosReplica)

	srv := &http.Server{
		Addr:         ":4001",
		Handler:      s,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 5 * time.Minute,
		IdleTimeout:  1 * time.Hour,
	}

	go Sincronizador()

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Error: Servidor de Replicación no inició")
	}
}