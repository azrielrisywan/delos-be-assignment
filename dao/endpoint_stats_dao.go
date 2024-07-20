package dao

import (
    "crud-app/config"
    "crud-app/dto"
    "errors"
)

func LogEndpointUsage(stats dto.EndpointStats) error {
    db := config.SetupDatabase()
    defer db.Close()

    sqlQuery := `
        INSERT INTO delos.endpointstatistics (n_endpoint, n_method, n_user_agent, d_request_time)
        VALUES ($1, $2, $3, $4)
    `
    _, err := db.Exec(sqlQuery, stats.Endpoint, stats.Method, stats.UserAgent, stats.RequestTime)
    if err != nil {
        return errors.New("failed to log request")
    }
    return nil
}

func GetEndpointStats() (map[string]map[string]interface{}, error) {
    db := config.SetupDatabase()
    defer db.Close()

    sqlQuery := `
        SELECT n_endpoint, n_method, COUNT(*) as count, COUNT(DISTINCT n_user_agent) as unique_user_agents
        FROM delos.endpointstatistics
        GROUP BY n_endpoint, n_method
    `

    rows, err := db.Queryx(sqlQuery)
    if err != nil {
        return nil, errors.New("failed to retrieve statistics")
    }
    defer rows.Close()

    result := make(map[string]map[string]interface{})
    recordsFound := false // Flag to check if any records are found

    for rows.Next() {
        recordsFound = true
        var endpoint, method string
        var count, uniqueUserAgents int
        if err := rows.Scan(&endpoint, &method, &count, &uniqueUserAgents); err != nil {
            return nil, errors.New("failed to scan row")
        }

        endpointKey := method + " " + endpoint
        result[endpointKey] = map[string]interface{}{
            "count":              count,
            "unique_user_agents": uniqueUserAgents,
        }
    }

    if !recordsFound {
        return nil, errors.New("no endpoints tracked")
    }

    return result, nil
}
