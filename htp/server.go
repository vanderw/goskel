package htp

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ServerCallback func(r *gin.Engine) error

type Server struct {
	// *fasthttp.Server
	*http.Server
	addr string // listen addr
}

func NewServer(name, addr string, cb ServerCallback) *Server {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// r.Use(cors.Default())

	if cb != nil {
		if err := cb(r); err != nil {
			panic(err)
		}
	}

	// more
	s := &Server{
		Server: &http.Server{
			Addr:    addr,
			Handler: r.Handler(),
		},
		addr: addr,
	}
	return s
}

func (s *Server) Start() {
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		return nil
	}
}
