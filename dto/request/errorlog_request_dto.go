package dto

import "time"

type ErrorLog struct {
	Hostname      string                 `json:"hostname"`
	Service       string                 `json:"service"`
	Code          string                 `json:"code"`
	Error         string                 `json:"error"`
	StackTrace    string                 `json:"stackTrace"`
	Timestamp     time.Time              `json:"timestamp"`
	RemoteIP      string                 `json:"remote_ip"`
	Host          string                 `json:"host"`
	Uri           string                 `json:"uri"`
	Method        string                 `json:"method"`
	Path          string                 `json:"path"`
	Referer       string                 `json:"referer"`
	UserAgent     string                 `json:"user_agent"`
	Status        int                    `json:"status"`
	RequestLength int64                  `json:"request_length"`
	BytesSent     int64                  `json:"bytes_sent"`
	Body          interface{}            `json:"body"`
	Params        map[string]interface{} `json:"params"`
	Controller    string                 `json:"controller"`
	Action        string                 `json:"action"`
	UserId        int64                  `json:"userId"`
	AuthToken     string                 `json:"-"`
}
