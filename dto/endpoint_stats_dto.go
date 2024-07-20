package dto

import "time"

type EndpointStats struct {
    Endpoint    string    `json:"endpoint"`
    Method      string    `json:"method"`
    UserAgent   string    `json:"user_agent"`
    RequestTime time.Time `json:"request_time"`
}
