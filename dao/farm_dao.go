package dao

import (
    "crud-app/config"
    "crud-app/dto"
    "errors"
    "time"
    "github.com/google/uuid"
)

func FarmList() ([]dto.Farm, error) {
    db := config.SetupDatabase()
    defer db.Close()

    sqlQuery := "SELECT i_id, n_name, d_created_on, c_deleted, d_deleted_on FROM delos.farm"

    rows, err := db.Queryx(sqlQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var farms []dto.Farm
    for rows.Next() {
        var farm dto.Farm
        if err := rows.StructScan(&farm); err != nil {
            return nil, err
        }
        farms = append(farms, farm)
    }

    if len(farms) == 0 {
        return nil, errors.New("FARMS_NOT_FOUND")
    }

    return farms, nil
}

func FarmListById(id uuid.UUID) (dto.Farm, error) {
    db := config.SetupDatabase()
    defer db.Close()

    sqlQuery := "SELECT i_id, n_name, d_created_on, c_deleted, d_deleted_on FROM delos.farm WHERE i_id = $1"

    var farm dto.Farm
    err := db.QueryRowx(sqlQuery, id).StructScan(&farm)
    if err != nil {
        if err.Error() == "sql: no rows in result set" {
            return dto.Farm{}, errors.New("FARM_NOT_FOUND")
        }
        return dto.Farm{}, err
    }

    return farm, nil
}

func CreateFarm(createFarmDto dto.CreateFarm) (dto.Farm, error) {
    db := config.SetupDatabase()
    defer db.Close()

    // Check for duplicate farm name
    var existingFarm dto.Farm
    duplicateCheckQuery := "SELECT i_id, n_name, d_created_on FROM delos.farm WHERE n_name = $1 and c_deleted = '0'"
    err := db.QueryRowx(duplicateCheckQuery, createFarmDto.Name).StructScan(&existingFarm)
    if err == nil {
        return dto.Farm{}, errors.New("DUPLICATE_FARM_NAME")
    }

    // Proceed to create a new farm
    sqlQuery := "INSERT INTO delos.farm (i_id, n_name) VALUES (uuid_generate_v4(), $1) RETURNING i_id, n_name, d_created_on"
    var farm dto.Farm
    err = db.QueryRowx(sqlQuery, createFarmDto.Name).StructScan(&farm)
    if err != nil {
        return dto.Farm{}, err
    }

    return farm, nil
}

func UpdateFarm(updateFarmDto dto.UpdateFarm) (int64, error) {
    db := config.SetupDatabase()
    defer db.Close()

    sqlQuery := "UPDATE delos.farm SET n_name = $1 WHERE i_id = $2"

    result, err := db.Exec(sqlQuery, updateFarmDto.Name, updateFarmDto.ID)
    if err != nil {
        return 0, err
    }

    rowsAffected, err := result.RowsAffected()
    return rowsAffected, err
}

func DeleteFarm(id uuid.UUID) error {
    db := config.SetupDatabase()
    defer db.Close()

    sqlQuery := `
        UPDATE delos.farm
        SET c_deleted = '1', d_deleted_on = $1
        WHERE i_id = $2 AND c_deleted = '0'
    `

    result, err := db.Exec(sqlQuery, time.Now().UTC(), id)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        return errors.New("FARM_NOT_FOUND_OR_ALREADY_DELETED")
    }

    return nil
}
