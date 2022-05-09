package version

func ReleaseVersion() string {
	return versionInfo.ReleaseVersion
}

func String() string {
	return versionInfo.String()
}

func ShortString() string {
	return versionInfo.ShortString()
}

func Json() string {
	return versionInfo.Json()
}

func Yaml() string {
	return versionInfo.Yaml()
}
