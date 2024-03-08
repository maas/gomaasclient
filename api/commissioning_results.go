package api

import "github.com/maas/gomaasclient/entity"

// CommissioningResults is an interface for node commissioning results
type CommissioningResults interface {
	Get(params *entity.InstallationResultParams) ([]entity.InstallationResult, error)
}
