package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"
)

type Server struct {
	srv  *http.Server
	conf Config
}

func NewServer(conf Config, handler http.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Handler:      handler,
			ReadTimeout:  time.Duration(conf.ReadTimeoutMs) * time.Millisecond,
			WriteTimeout: time.Duration(conf.WriteTimeoutMs) * time.Millisecond,
			Addr:         fmt.Sprintf("0.0.0.0:%d", conf.Port),
		},
		conf: conf,
	}
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.srv.Addr)
	if err != nil {
		return err
	}

	go func() {
		fmt.Printf("Starting http server on port %d\n", s.conf.Port)

		err := s.srv.Serve(lis)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Failed to start http server")
		}
	}()

	return nil
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1000)*time.Millisecond)
	defer cancel()
	fmt.Println("Shutting down http server")

	return s.srv.Shutdown(ctx)
}
