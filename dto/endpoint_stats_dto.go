package dto

import "time"

type EndpointStats struct {
    Endpoint    string    `json:"endpoint"`
    Method      string    `json:"method"`
    UserAgent   string    `json:"user_agent"`
    RequestTime time.Time `json:"request_time"`
}

type EndpointStatsResponse map[string]struct {
    Count             int `json:"count"`
    UniqueUserAgents  int `json:"unique_user_agents"`
}

type EndpointStatsErrorResponse struct {
    Error string `json:"error" example:"no endpoints tracked"`
}
