package dao

import (
    "crud-app/config"
    "crud-app/dto"
    "errors"
    "github.com/google/uuid"
	"time"
)

func PondList() ([]dto.Pond, error) {
    db := config.SetupDatabase()
    defer db.Close()

    sqlQuery := "SELECT i_id, i_id_farm, n_name, d_created_on, c_deleted, d_deleted_on FROM delos.pond"

    rows, err := db.Queryx(sqlQuery)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var ponds []dto.Pond
    for rows.Next() {
        var pond dto.Pond
        if err := rows.StructScan(&pond); err != nil {
            return nil, err
        }
        ponds = append(ponds, pond)
    }

    if len(ponds) == 0 {
        return nil, errors.New("PONDS_NOT_FOUND")
    }

    return ponds, nil
}

func PondListById(id uuid.UUID) (dto.Pond, error) {
    db := config.SetupDatabase()
    defer db.Close()

    sqlQuery := "SELECT i_id, i_id_farm, n_name, d_created_on, c_deleted, d_deleted_on FROM delos.pond WHERE i_id = $1"

    var pond dto.Pond
    err := db.QueryRowx(sqlQuery, id).StructScan(&pond)
    if err != nil {
        if err.Error() == "sql: no rows in result set" {
            return dto.Pond{}, errors.New("POND_NOT_FOUND")
        }
        return dto.Pond{}, err
    }

    return pond, nil
}

func CreatePond(createPondDto dto.CreatePond) (dto.Pond, error) {
    db := config.SetupDatabase()
    defer db.Close()

    // Check if the Farm ID exists
    var farmExists bool
    farmCheckQuery := "SELECT EXISTS (SELECT 1 FROM delos.farm WHERE i_id = $1 AND c_deleted = '0')"
    err := db.QueryRow(farmCheckQuery, createPondDto.FarmID).Scan(&farmExists)
    if err != nil {
        return dto.Pond{}, err
    }
    if !farmExists {
        return dto.Pond{}, errors.New("FARM_NOT_FOUND")
    }

    // Check for duplicate pond name
    var existingPond dto.Pond
    duplicateCheckQuery := "SELECT i_id, i_id_farm, n_name, d_created_on FROM delos.pond WHERE n_name = $1 AND c_deleted = '0'"
    err = db.QueryRowx(duplicateCheckQuery, createPondDto.Name).StructScan(&existingPond)
    if err == nil {
        return dto.Pond{}, errors.New("DUPLICATE_POND_NAME")
    }

    // Proceed to create a new pond
    sqlQuery := "INSERT INTO delos.pond (i_id, i_id_farm, n_name) VALUES (uuid_generate_v4(), $1, $2) RETURNING i_id, i_id_farm, n_name, d_created_on"
    var pond dto.Pond
    err = db.QueryRowx(sqlQuery, createPondDto.FarmID, createPondDto.Name).StructScan(&pond)
    if err != nil {
        return dto.Pond{}, err
    }

    return pond, nil
}

func UpdatePond(updatePondDto dto.UpdatePond) (int64, error) {
    db := config.SetupDatabase()
    defer db.Close()

    sqlQuery := "UPDATE delos.pond SET n_name = $1 WHERE i_id = $2"

    result, err := db.Exec(sqlQuery, updatePondDto.Name, updatePondDto.ID)
    if err != nil {
        return 0, err
    }

    rowsAffected, err := result.RowsAffected()
    return rowsAffected, err
}

func DeletePond(id uuid.UUID) error {
    db := config.SetupDatabase()
    defer db.Close()

    sqlQuery := `
        UPDATE delos.pond
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
        return errors.New("POND_NOT_FOUND_OR_ALREADY_DELETED")
    }

    return nil
}
