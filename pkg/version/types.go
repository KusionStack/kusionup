//go:generate go run gen.go
//go:generate go fmt

package version

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	"gopkg.in/yaml.v3"

	goversion "github.com/hashicorp/go-version"

	git "github.com/KusionStack/kusionup/pkg/util/gitutil"
)

var versionInfo *Info

// Info contains versioning information.
// following attributes:
//
//	   ReleaseVersion - "vX.Y.Z-00000000" used to indicate the last release version,
//			  containing GitVersion and GitCommitShort.
//	   GitVersion - "vX.Y.Z" used to indicate the last git tag.
//	   GitCommit - The git commit id corresponding to this source code.
//	   GitTreeState - "clean" indicates no changes since the git commit id
//	       "dirty" indicates source code changes after the git commit id
type Info struct {
	ReleaseVersion string      `json:"releaseVersion" yaml:"releaseVersion"`                 // Such as "v1.2.3-3836f877"
	GitVersion     string      `json:"gitVersion" yaml:"gitVersion"`                         // Such as "v1.2.3"
	GitCommit      string      `json:"gitCommit" yaml:"gitCommit"`                           // Such as "3836f8770ab8f488356b2129f42f2ae5c1134bb0"
	GitTreeState   string      `json:"gitTreeState,omitempty" yaml:"gitTreeState,omitempty"` // Such as "clean", "dirty"
	BuildTime      string      `json:"buildTime" yaml:"buildTime"`                           // Such as "2021-10-20 18:24:03"
	Runtime        RuntimeInfo `json:"runtime,omitempty" yaml:"runtime,omitempty"`
}

type RuntimeInfo struct {
	GoVersion string `json:"goVersion,omitempty" yaml:"goVersion,omitempty"`
	GOOS      string `json:"GOOS,omitempty" yaml:"GOOS,omitempty"`
	GOARCH    string `json:"GOARCH,omitempty" yaml:"GOARCH,omitempty"`
	NumCPU    int    `json:"numCPU,omitempty" yaml:"numCPU,omitempty"`
	Compiler  string `json:"compiler,omitempty" yaml:"compiler,omitempty"`
}

func NewInfo() (*Info, error) {
	var (
		isHeadAtTag    bool
		headHash       string
		headHashShort  string
		latestTag      string
		gitVersion     *goversion.Version
		releaseVersion string
		isDirty        bool
		gitTreeState   string
		err            error
	)

	// Get git info
	if headHash, err = git.GetHeadHash(); err != nil {
		return nil, err
	}

	if headHashShort, err = git.GetHeadHashShort(); err != nil {
		return nil, err
	}

	if latestTag, err = git.GetLatestTag(); err != nil {
		return nil, err
	}

	if gitVersion, err = goversion.NewVersion(latestTag); err != nil {
		return nil, err
	}

	if isHeadAtTag, err = git.IsHeadAtTag(latestTag); err != nil {
		return nil, err
	}

	if isDirty, err = git.IsDirty(); err != nil {
		return nil, err
	}

	// Get git tree state
	if isDirty {
		gitTreeState = "dirty"
	} else {
		gitTreeState = "clean"
	}

	// Get release version
	if isHeadAtTag {
		releaseVersion = gitVersion.Original()
	} else {
		releaseVersion = fmt.Sprintf("%s-%s", gitVersion.Original(), headHashShort)
	}

	return &Info{
		ReleaseVersion: releaseVersion,
		GitVersion:     gitVersion.Original(),
		GitCommit:      headHash,
		GitTreeState:   gitTreeState,
		BuildTime:      time.Now().Format("2006-01-02 15:04:05"),
		Runtime: RuntimeInfo{
			GoVersion: runtime.Version(),
			GOOS:      runtime.GOOS,
			GOARCH:    runtime.GOARCH,
			NumCPU:    runtime.NumCPU(),
			Compiler:  runtime.Compiler,
		},
	}, nil
}

func (v *Info) String() string {
	return v.YAML()
}

func (v *Info) ShortString() string {
	return fmt.Sprintf("%s; git: %s; build time: %s", v.ReleaseVersion, v.GitCommit, v.BuildTime)
}

func (v *Info) JSON() string {
	data, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return ""
	}

	return string(data)
}

func (v *Info) YAML() string {
	data, err := yaml.Marshal(v)
	if err != nil {
		return ""
	}

	return string(data)
}
