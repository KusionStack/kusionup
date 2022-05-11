package github

var downloadURLMap = map[string]string{
	"linux-amd64":   "http://TODO/cli/kusion-open/latest/kusion-linux.tgz",
	"darwin-amd64":  "http://TODO/cli/kusion-open/latest/kusion-darwin.tgz",
	"darwin-arm64":  "http://TODO/cli/kusion-open/latest/kusion-darwin-arm64.tgz",
	"windows-amd64": "http://TODO/cli/kusion-open/latest/kusion-windows.zip",
}

func getOsArchKey(goos, goarch string) string {
	return goos + "-" + goarch
}
