package versioning

// GetCurrentVersion returns the current version of the Contracts.
func GetCurrentVersion() Version {
	return Version{
		Major:            1,
		Minor:            0,
		Patch:            3,
		PreReleaseString: "",
	}
}
