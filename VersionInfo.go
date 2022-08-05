package integrationcontracts

import (
	"dolittle.io/integrations/v1/contracts/versioning"
)

// GetCurrentVersion returns the current version of the Contracts.
func GetCurrentVersion() versioning.Version {
	return versioning.Version{
		Major:            1,
		Minor:            0,
		Patch:            1,
		PreReleaseString: "",
	}
}
