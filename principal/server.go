package principal

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	r.GET("/truco", GetAllTrucosHandler)
	r.POST("/truco", CreateTrucoHandler)
	r.PUT("/truco/:id", UpdateTrucoHandler)
	r.DELETE("/truco/:id", DeleteTrucoHandler)

	srv := &http.Server{
		Addr: ":4000",
		Handler: r,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 5*time.Minute,
		IdleTimeout: 1*time.Hour,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Error: Server Main hasn't begin")
	}
}