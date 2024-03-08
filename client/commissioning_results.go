package client

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// CommissioningResults implements api.CommissioningResults
type CommissioningResults struct {
	APIClient APIClient
}

func (c *CommissioningResults) client() APIClient {
	return c.APIClient.GetSubObject("installation-results")
}

// Get installation results
func (c *CommissioningResults) Get(params *entity.InstallationResultParams) ([]entity.InstallationResult, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	results := make([]entity.InstallationResult, 0)
	err = c.client().Get("", qsp, func(data []byte) error {
		return json.Unmarshal(data, &results)
	})

	return results, err
}
