package ingest

import (
	"code.riskrancher.com/RiskRancher/core/pkg/domain"
)

type Handler struct {
	Store domain.Store
}

func NewHandler(store domain.Store) *Handler {
	return &Handler{Store: store}
}
