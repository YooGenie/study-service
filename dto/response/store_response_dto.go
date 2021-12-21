package response

import (
	"encoding/json"
)

type StoreSummary struct {
	No                         int64           `json:"no"`
	Id                         string          `json:"id"`
	Mobile                     string          `json:"mobile"`
	BusinessRegistrationNumber string          `json:"businessRegistrationNumber"`
	Created                    json.RawMessage `json:"created"`
	Updated                    json.RawMessage `json:"updated"`
}
