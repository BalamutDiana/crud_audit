package server

import (
	"context"

	audit "github.com/BalamutDiana/crud_audit/pkg/domain"
)

type AuditService interface {
	Insert(ctx context.Context, req audit.LogItem) error
}

type AuditServer struct {
	service AuditService
}

func NewAuditServer(service AuditService) *AuditServer {
	return &AuditServer{
		service: service,
	}
}

func (h *AuditServer) Log(ctx context.Context, req audit.LogItem) error {
	err := h.service.Insert(ctx, req)
	return err
}
