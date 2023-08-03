package custom

import (
	"bytes"
	"fmt"
	"html/template"
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
	osArchKey := getOsArchKey(runtime.GOOS, runtime.GOARCH)
	if urlPattern, ok := archiveDownloadURLMap[osArchKey]; ok {
		tmpl, err := template.New("").Parse(urlPattern)
		if err != nil {
			return "", err
		}
		var buf bytes.Buffer
		data := struct{ VERSION string }{ver}
		if err := tmpl.Execute(&buf, data); err != nil {
			return "", err
		}
		fmt.Println(buf.String())
		return buf.String(), nil
	}
	return "", ErrUnsupportedOsArch
}

func getOsArchKey(goos, goarch string) string {
	return goos + "-" + goarch
}
