package dto

import (
    "time"
    "github.com/google/uuid"
)

type CreatePond struct {
    FarmID uuid.UUID `json:"farm_id" example:"c48e3c9d-50a8-400c-b63f-f72b67c6fe5b"`
    Name   string    `json:"name" example:"Main Pond"`
}

type UpdatePond struct {
    ID     uuid.UUID `json:"id" example:"d01629f8-c708-4c92-b950-848eefd26b9a"`
    FarmID uuid.UUID `json:"farm_id" example:"c48e3c9d-50a8-400c-b63f-f72b67c6fe5b"`
    Name   string    `json:"name" example:"Updated Pond"`
}

type Pond struct {
    ID        uuid.UUID  `json:"id" db:"i_id" example:"c48e3c9d-50a8-400c-b63f-f72b67c6fe5b"`
    FarmID    uuid.UUID  `json:"farm_id" db:"i_id_farm" example:"d01629f8-c708-4c92-b950-848eefd26b9a"`
    Name      string     `json:"name" db:"n_name" example:"Main Pond"`
    CreatedOn time.Time  `json:"created_on" db:"d_created_on" example:"2024-07-19T13:41:42.770296Z"`
    Deleted   *string    `json:"deleted,omitempty" db:"c_deleted" example:"0"`
    DeletedOn *time.Time `json:"deleted_on,omitempty" db:"d_deleted_on" example:"null"`
    Farm      *Farm       `json:"farm,omitempty"`
}

type DeletePondResponse struct {
    Message string `json:"message" example:"Pond deleted successfully"`
}

// ErrorResponse represents the error response structure
type PondListErrorResponse struct {
    Error string `json:"error" example:"No ponds found"`
}

type PondListByIdErrorResponse struct {
    Error string `json:"error" example:"Invalid ID or Pond not found"`
}

type CreatePondErrorResponse struct {
    Error string `json:"error" example:"Failed to create pond"`
}

type UpdatePondErrorResponse struct {
    Error string `json:"error" example:"Failed to update pond"`
}

type DeletePondErrorResponse struct {
    Error string `json:"error" example:"Failed to delete pond"`
}
