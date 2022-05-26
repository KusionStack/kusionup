package custom

import (
	"fmt"
	"runtime"
)

var (
	ErrUnsupportedOsArch   = fmt.Errorf("unsupported os/arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	ErrEmptyDownloadURLMap = fmt.Errorf("archiveDownloadURLMap is nil")
)

func getArchiveDownloadURL(archiveDownloadURLMap map[string]string, ver string) (string, error) {
	if archiveDownloadURLMap == nil {
		return "", ErrEmptyDownloadURLMap
	}

	if urlPattern, ok := archiveDownloadURLMap[getOsArchKey(runtime.GOOS, runtime.GOARCH)]; ok {
		return fmt.Sprintf(urlPattern, ver), nil
	}

	return "", ErrUnsupportedOsArch
}

func getOsArchKey(goos, goarch string) string {
	return goos + "-" + goarch
}
