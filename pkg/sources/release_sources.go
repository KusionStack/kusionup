package sources

type ReleaseSource interface {
	GetName() string                               // Get release source name
	GetVersions() []string                         // Get versions from release source
	GetDownloadURL(version string) (string, error) // Get download url of specify version
}
