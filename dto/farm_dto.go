package dto

import (
	"time"
	"github.com/google/uuid"
)

type CreateFarm struct {
	Name string `json:"name"`
}

type UpdateFarm struct {
	ID   uuid.UUID `db:"i_id" json:"id"`
	Name string `json:"name"`
}

type Farm struct {
    ID        uuid.UUID  `db:"i_id" json:"id"`
    Name      string     `db:"n_name" json:"name"`
    CreatedOn *time.Time `db:"d_created_on" json:"created_on"`
	Deleted   string     `db:"c_deleted" json:"deleted"`
	DeletedOn *time.Time `db:"d_deleted_on" json:"deleted_on"`
}

type FarmResponse struct {
	Message string `json:"message"`
}

