package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/BalamutDiana/crud_audit/internal/config"
	"github.com/BalamutDiana/crud_audit/internal/repository"
	"github.com/BalamutDiana/crud_audit/internal/server"
	"github.com/BalamutDiana/crud_audit/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client()
	opts.SetAuth(options.Credential{
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
	})
	opts.ApplyURI(cfg.DB.URI)

	dbClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	if err := dbClient.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	db := dbClient.Database(cfg.DB.Database)

	auditRepo := repository.NewAudit(db)
	auditService := service.NewAudit(auditRepo)
	auditSrv := server.NewAuditServer(auditService)
	srv := server.New(auditSrv, cfg.Server.Port)
	defer srv.CloseConnection()

	fmt.Println("SERVER STARTED", time.Now())

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
