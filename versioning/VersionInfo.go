package versioning

// GetCurrentVersion returns the current version of the Contracts.
func GetCurrentVersion() Version {
	return Version{
		Major:            2,
		Minor:            9,
		Patch:            1,
		PreReleaseString: "",
	}
}
