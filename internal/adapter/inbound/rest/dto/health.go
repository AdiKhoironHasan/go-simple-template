package dto

import "github.com/adikhoironhasan/go-simple-template/internal/core/domain/entity"

type CheckHealthRequest struct {
	MongoDB bool `query:"mongodb"`
	Redis   bool `query:"redis"`
}

type CheckHealthResponse struct {
	MongoDB bool `json:"mongodb"`
	Redis   bool `json:"redis"`
}

func (d *CheckHealthRequest) ToEntity() entity.CheckHealth {
	return entity.CheckHealth{
		MongoDB: d.MongoDB,
		Redis:   d.Redis,
	}
}

func ToCheckHealthResponse(req *entity.CheckHealth) *CheckHealthResponse {
	return &CheckHealthResponse{
		MongoDB: req.MongoDB,
		Redis:   req.Redis,
	}
}
