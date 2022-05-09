package open

var downloadUrlMap map[string]string = map[string]string{
	"linux-amd64":   "http://TODO/cli/kusion-open/latest/kusion-linux.tgz",
	"darwin-amd64":  "http://TODO/cli/kusion-open/latest/kusion-darwin.tgz",
	"darwin-arm64":  "http://TODO/cli/kusion-open/latest/kusion-darwin-arm64.tgz",
	"windows-amd64": "http://TODO/cli/kusion-open/latest/kusion-windows.zip",
}

func getOsArchKey(GOOS, GOARCH string) string {
	return GOOS + "-" + GOARCH
}
