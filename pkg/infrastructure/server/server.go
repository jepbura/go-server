package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/jepbura/go-server/pkg/config"
	"golang.org/x/crypto/acme/autocert"

	"go.uber.org/zap"
)

type ServerHTTP struct {
	Environment string `name:"env"`
	Port        string `name:"port"`
	Logger      *zap.Logger
	Server      *http.Server
	Manager     *autocert.Manager
}

func RunServer(cnf config.Env, Logger *zap.Logger) *ServerHTTP {
	fmt.Print("*********************************************\n")
	fmt.Print("RunServer\n")
	fmt.Print("*********************************************\n")
	var man *autocert.Manager
	engine := gin.New()

	// zap.Logger integration with gin
	engine.Use(ginzap.Ginzap(Logger, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(Logger, true))

	server := &ServerHTTP{
		Environment: cnf.Environment,
		Port:        cnf.Port,
		Logger:      Logger,
	}

	if cnf.Environment != "local" {
		host := ""
		man = &autocert.Manager{
			Prompt: autocert.AcceptTOS,
			Cache:  autocert.DirCache("certs"),
		}

		server.Server = &http.Server{
			Addr:    host + ":443",
			Handler: engine,
			TLSConfig: &tls.Config{
				GetCertificate: man.GetCertificate,
			},
		}
		server.Manager = man
	} else {
		host := "localhost"
		server.Server = &http.Server{
			Addr:    host + ":" + cnf.Port,
			Handler: engine,
		}
	}

	return server
}

func (s *ServerHTTP) StartServer() {
	if s.Environment != "local" {
		s.Logger.Info("Starting HTTPS server at " + s.Server.Addr)
		go s.Server.ListenAndServeTLS("", "")
		go http.ListenAndServe(":80", s.Manager.HTTPHandler(nil))
	} else {
		s.Logger.Info("Starting HTTP server at " + s.Server.Addr)
		go s.Server.ListenAndServe()
	}
}

func (s *ServerHTTP) StopServer(ctx context.Context) error {
	s.Logger.Info("Stopping HTTPS server.")
	return s.Server.Shutdown(ctx)
}
