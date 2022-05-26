package github

import (
	"github.com/KusionStack/kusionup/pkg/sources"
	"github.com/KusionStack/kusionup/pkg/util/gitutil"
)

var GithubReleaseSource sources.ReleaseSource = &releaseSource{}

type releaseSource struct {
	name string // release source name
}

func (s *releaseSource) GetName() string {
	s.name = "github"

	return s.name
}

func (s *releaseSource) GetVersions() []string {
	versions := []string{"latest"}

	// Get tags from github repo of kusion
	remoteURL := "https://github.com/KusionStack/kusion"

	tags, err := gitutil.GetTagListFromRemote(remoteURL, true)
	if err != nil {
		// klog.Warningf("Get tag list from remote failed, err: %v", err)
		return versions
	}

	// Hidden timeout tags
	isHiddenTag := getIsHiddenTag()
	for _, tag := range tags {
		if _, hidden := isHiddenTag[tag]; !hidden {
			versions = append(versions, tag)
		}
	}

	return versions
}

func (s *releaseSource) GetDownloadURL(ver string) (string, error) {
	vers := s.GetVersions()
	if ver == "latest" && len(vers) > 0 {
		return getArchiveDownloadURL(vers[1])
	}

	return getArchiveDownloadURL(ver)
}

func getIsHiddenTag() map[string]struct{} {
	hiddenTags := []string{"v0.4.0"}
	isHiddenTag := make(map[string]struct{})

	for _, tag := range hiddenTags {
		isHiddenTag[tag] = struct{}{}
	}

	return isHiddenTag
}
