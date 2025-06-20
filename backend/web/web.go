package web

import (
	"context"
	"net"
	"net/http"
	"toy-duman/web/controller"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
	listener   net.Listener

	index *controller.IndexController

	ctx    context.Context
	cancel context.CancelFunc
}

func NewServer() *Server {
	ctx, cancel := context.WithCancel(context.Background())
	return &Server{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *Server) initRouter() (*gin.Engine, error) {
	gin.SetMode(gin.DebugMode)

	// TODO
	// if mode is release
	// gin.DefaultWriter = io.Discard
	// gin.DefaultErrorWriter = io.Discard
	// gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()

	engine.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/api/"})))

	g := engine.Group("/")

	s.index = controller.NewIndexController(g)

	return engine, nil
}

func (s *Server) Start() (err error) {
	defer func() {
		if err != nil {
			s.Stop()
		}
	}()

	engine, err := s.initRouter()
	if err != nil {
		return err
	}

	listenAddr := net.JoinHostPort("localhost", "8080")
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	s.listener = listener

	s.httpServer = &http.Server{
		Handler: engine,
	}

	go func() {
		s.httpServer.Serve(listener)
	}()

	return nil
}

func (s *Server) Stop() error {
	s.cancel()

	var err error

	if s.httpServer != nil {
		err = s.httpServer.Shutdown(s.ctx)
	}

	if s.listener != nil {
		err = s.listener.Close()
	}
	return err
}
