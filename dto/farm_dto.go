package dto

import (
	"time"
	"github.com/google/uuid"
)

type CreateFarm struct {
	Name string `json:"name"`
}

type UpdateFarm struct {
	ID   uuid.UUID `db:"i_id" json:"id" binding:"required" example:"c48e3c9d-50a8-400c-b63f-f72b67c6fe5b"`
	Name string `json:"name" binding:"required" example:"Bero Farm Updated"`
}

type Farm struct {
    ID        string     `db:"i_id" json:"id" example:"c48e3c9d-50a8-400c-b63f-f72b67c6fe5b"`
    Name      string     `db:"n_name" json:"name" example:"Bero Farm"`
    CreatedOn time.Time  `db:"d_created_on" json:"created_on" example:"2024-07-19T13:41:42.770296Z"`
    Deleted   string     `db:"c_deleted" json:"deleted" example:"0"`
    DeletedOn *time.Time `db:"d_deleted_on" json:"deleted_on" example:"null"`
}

type FarmListResponse []Farm

type FarmResponse Farm

type CreateFarmResponse Farm

type UpdateFarmResponse UpdateFarm

type DeleteFarmResponse struct {
    Message string `json:"message" example:"Farm deleted successfully"`
}

type FarmListErrorResponse struct {
    Error string `json:"error" example:"No farms found"`
}

type CreateFarmErrorResponse struct {
    Error string `json:"error" example:"Duplicate farm name"`
}

type UpdateFarmErrorResponse struct {
    Error string `json:"error" example:"Invalid ID"`
}

type DeleteFarmErrorResponse struct {
	Error string `json:"error" example:"Invalid ID or Farm not found"`
}

