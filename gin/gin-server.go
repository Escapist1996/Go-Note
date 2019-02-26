package main

import ( 
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func stest() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	s := &http.Server{
		Addr:           ":8080",
		//Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	//r.Run() // listen and serve on 0.0.0.0:8080
}