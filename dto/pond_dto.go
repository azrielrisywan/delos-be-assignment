package dto

import (
    "time"
    "github.com/google/uuid"
)

type CreatePond struct {
    FarmID uuid.UUID `json:"farm_id"`
    Name   string    `json:"name"`
}

type UpdatePond struct {
    ID     uuid.UUID `json:"id"`
    FarmID uuid.UUID `json:"farm_id"`
    Name   string    `json:"name"`
}

type Pond struct {
    ID        uuid.UUID `json:"id" db:"i_id"`
    FarmID    uuid.UUID `json:"farm_id" db:"i_id_farm"`
    Name      string    `json:"name" db:"n_name"`
    CreatedOn time.Time `json:"created_on" db:"d_created_on"`
    Deleted   *string   `json:"deleted,omitempty" db:"c_deleted"`
    DeletedOn *time.Time `json:"deleted_on,omitempty" db:"d_deleted_on"`
}
