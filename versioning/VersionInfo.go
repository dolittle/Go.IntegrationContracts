package versioning

// GetCurrentVersion returns the current version of the Contracts.
func GetCurrentVersion() Version {
	return Version{
		Major:            2,
		Minor:            13,
		Patch:            0,
		PreReleaseString: "",
	}
}
