package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	router := gin.Default()
	m := melody.New()
	m.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	router.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "app/index.html")
	})

	v1 := router.Group("/v1")
	{
		v1.GET("/ws", func(c *gin.Context) {
			m.HandleRequest(c.Writer, c.Request)
		})

		m.HandleMessage(func(s *melody.Session, msg []byte) {
			m.Broadcast(msg)
		})
	}

	http.Handle("/", router)
	port := "8080"
	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
