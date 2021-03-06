package github

import (
	"fmt"
	"runtime"
)

var ErrUnsupportedOsArch = fmt.Errorf("unsupported os/arch: %s/%s", runtime.GOOS, runtime.GOARCH)

func getArchiveDownloadURL(ver string) (string, error) {
	archiveDownloadURLMap := map[string]string{
		"linux-amd64":   "https://github.com/KusionStack/kusion/releases/download/%s/kusion-linux.tgz",
		"darwin-amd64":  "https://github.com/KusionStack/kusion/releases/download/%s/kusion-darwin.tgz",
		"darwin-arm64":  "https://github.com/KusionStack/kusion/releases/download/%s/kusion-darwin-arm64.tgz",
		"windows-amd64": "https://github.com/KusionStack/kusion/releases/download/%s/kusion-windows.tgz",
	}

	if urlPattern, ok := archiveDownloadURLMap[getOsArchKey(runtime.GOOS, runtime.GOARCH)]; ok {
		return fmt.Sprintf(urlPattern, ver), nil
	}

	return "", ErrUnsupportedOsArch
}

func getOsArchKey(goos, goarch string) string {
	return goos + "-" + goarch
}
