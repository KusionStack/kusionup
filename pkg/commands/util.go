package commands

import (
	"fmt"
	"path/filepath"
	"strings"
)

// e.g. open@v0.2.10, open@latest
func GetSourceVersionTitle(source, version string) string {
	return fmt.Sprintf("%s@%s", source, version)
}

// e.g. open@v0.2.10 => "open", "v0.2.10"
func ParseSourceVersion(title string) (source, version string, err error) {
	s := strings.Split(title, "@")
	if len(s) != 2 {
		return "", "", fmt.Errorf("invalid source version key: %s", title)
	} else {
		if _, ok := registedReleaseSources[s[0]]; ok {
			return s[0], s[1], nil
		} else {
			return "", "", fmt.Errorf("unsupported source: %s", s[0])
		}
	}
}

// e.g. /Users/user/.kusionup/bin
func KusionupBinDir() string {
	return KusionupDir("bin")
}

// e.g. /Users/user/.kusionup/current
func KusionupCurrentDir() string {
	return KusionupDir("current")
}

// e.g. /Users/user/.kusionup/env
func KusionupEnvFile() string {
	return KusionupDir("env")
}

// e.g. /Users/user/.kusionup/current/bin and /Users/user/.kusionup/current/kclvm/bin
func KusionupCurrentBinDir() string {
	return KusionupDir("current", "bin") + " and " + KusionupDir("current", "kclvm", "bin")
}

// e.g. /Users/user/.kusionup/kusion-open-v0.2.10
func KusionupVersionDir(ver string) string {
	if !strings.HasPrefix(ver, "kusion-") {
		ver = "kusion-" + ver
	}
	return KusionupDir(ver)
}

// e.g. /Users/user/.kusionup/kusion-open-v0.2.10
func KusionupSourceVersionDir(source, ver string) string {
	return KusionupDir(GetSourceVersionDirName(source, ver))
}

// e.g. /Users/user/.kusionup
func KusionupDir(paths ...string) string {
	elem := []string{homedir, ".kusionup"}
	elem = append(elem, paths...)

	return filepath.Join(elem...)
}

// e.g. kusion-open-v0.2.10, kusion-open-latest
func GetSourceVersionDirName(source, ver string) string {
	return fmt.Sprintf("kusion-%s", GetSourceVersionTitle(source, ver))
}
