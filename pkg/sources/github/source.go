package github

import (
	"fmt"
	"runtime"

	"github.com/KusionStack/kusionup/pkg/sources"
)

var GithubReleaseSource sources.ReleaseSource = &releaseSource{}

type releaseSource struct {
	name string // release source name
}

func (s *releaseSource) GetName() string {
	s.name = "open"

	return s.name
}

func (s *releaseSource) GetVersions() []string {
	return []string{"latest"}
}

func (s *releaseSource) GetDownloadURL(ver string) (string, error) {
	if url, ok := downloadURLMap[getOsArchKey(runtime.GOOS, runtime.GOARCH)]; ok {
		return url, nil
	}

	return "", fmt.Errorf("unsupported os/arch: %s/%s", runtime.GOOS, runtime.GOARCH)
}
