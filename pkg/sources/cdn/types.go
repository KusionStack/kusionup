package cdn

import (
	"fmt"
	"runtime"
)

var ErrUnsupportedOsArch = fmt.Errorf("unsupported os/arch: %s/%s", runtime.GOOS, runtime.GOARCH)

func getArchiveDownloadURL(ver string) (string, error) {
	archiveDownloadURLMap := map[string]string{
		"linux-amd64":   "https://ghproxy.com/https://github.com/KusionStack/kusion/releases/download/%s/kusion_%s_linux_amd64.tar.gz",
		"darwin-amd64":  "https://ghproxy.com/https://github.com/KusionStack/kusion/releases/download/%s/kusion_%s_darwin_amd64.tar.gz",
		"darwin-arm64":  "https://ghproxy.com/https://github.com/KusionStack/kusion/releases/download/%s/kusion_%s_darwin_arm64.tar.gz",
		"windows-amd64": "https://ghproxy.com/https://github.com/KusionStack/kusion/releases/download/%s/kusion_%s_windows_amd64.zip",
	}

	if urlPattern, ok := archiveDownloadURLMap[getOsArchKey(runtime.GOOS, runtime.GOARCH)]; ok {
		return fmt.Sprintf(urlPattern, ver, ver[1:]), nil
	}

	return "", ErrUnsupportedOsArch
}

func getOsArchKey(goos, goarch string) string {
	return goos + "-" + goarch
}
