package rest

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	server http.Server
	
}

var readTimeout,writeTimeout = time.Second*10,time.Second*10
func NewServer(port string, router *gin.Engine) *Server {
	return &Server{
		server: http.Server{
			Addr:                         fmt.Sprintf(":%s",port),
			Handler:                      router,
			ReadTimeout:                  readTimeout,
			WriteTimeout:                 writeTimeout,
			MaxHeaderBytes:               1<<20,
		},
	}
}

func (s *Server) Run()  error{
	return s.server.ListenAndServe()

}
func (s *Server) ShutDown(ctx context.Context)  error{
	return s.server.Shutdown(ctx)

}