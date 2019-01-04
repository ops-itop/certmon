package certmon

import (
	"crypto/tls"
	"time"
)

//Results is the final output
type Results struct {
	Results   []Result  `json:"results"`
	Timestamp time.Time `json:"timestamp"`
}

//Result is the output for individual check
type Result struct {
	Hostname  string                      `json:"hostname"` //The value in SNI and certificate will be validated against it.
	Endpoints map[string]IndividualResult `json:"endpoints"`
	Timestamp time.Time                   `json:"timestamp"`
}

//IndividualResult is the output of checking individual endpoint
type IndividualResult struct {
	ConnectionState *tls.ConnectionState `json:"connectionstate"`
	Expiry          time.Time            `json:"expiry"`
	Warn            int                  `json: "threshwarn"`
	Crit            int                  `json: "threshcrit"`
	Days            int                  `json:"days"`
	OK              bool                 `json:"ok"`
	Status          string               `json:"status"`
	Err             string               `json:"err"`
	Timestamp       time.Time            `json:"timestamp"`
}
