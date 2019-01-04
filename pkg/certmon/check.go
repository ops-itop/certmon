package certmon

import (
	"crypto/tls"
	"net"
	"time"
)

//Run runs the tests and generates result
func (c Config) Run() Results {
	results := Results{Results: make([]Result, len(c.Checks)), Timestamp: time.Now().UTC()}
	for i, check := range c.Checks {
		results.Results[i] = check.Run()
	}
	return results
}

//Run runs the tests and generates result
func (c Check) Run() Result {
	result := Result{Hostname: c.Hostname, Endpoints: make(map[string]IndividualResult), Timestamp: time.Now()}
	for _, endpoint := range c.Endpoints {
		result.Endpoints[endpoint] = checkEndpoint(c.Hostname, endpoint, c.Warning, c.Critical)
	}
	return result
}

func checkEndpoint(hostname, endpoint string, warning int, critical int) (result IndividualResult) {
	//Can we dial?
	result.Timestamp = time.Now()
	dialer := &net.Dialer{Timeout: time.Minute}
	conn, err := tls.DialWithDialer(dialer, "tcp", net.JoinHostPort(endpoint, "443"), &tls.Config{ServerName: hostname})
	if err != nil {
		result.Err = err.Error()
		return
	}
	defer conn.Close()
	//Can we handshake?
	err = conn.Handshake()
	if err != nil {
		result.Err = err.Error()
		return
	}
	//Stamp Connection state
	cs := conn.ConnectionState()
	result.ConnectionState = &cs
	//Find the earliest expiring cert
	for _, crt := range cs.PeerCertificates {
		if result.Expiry.IsZero() || result.Expiry.After(crt.NotAfter) {
			result.Expiry = crt.NotAfter
		}
	}
	if result.Expiry.After(time.Now()) {
		result.OK = true
	}

	if critical == 0 {
		critical = 20
	}

	if warning == 0 {
		warning = 45
	}

	result.Days = int(result.Expiry.Sub(time.Now()).Hours()) / 24
	if result.Days < critical {
		result.Status = "critical"
	} else if result.Days < warning {
		result.Status = "warning"
	} else {
		result.Status = "ok"
	}

	result.Warn = warning
	result.Crit = critical

	return result
}
