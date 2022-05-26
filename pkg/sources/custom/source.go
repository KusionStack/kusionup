package custom

import (
	"github.com/KusionStack/kusionup/pkg/sources"
	"github.com/KusionStack/kusionup/pkg/util/gitutil"
)

var _ sources.ReleaseSource = &CustomSource{}

// CustomSource is a custom source implementation.
type CustomSource struct {
	Name         string            `yaml:"name,omitempty" json:"name,omitempty"`
	RemoteURL    string            `yaml:"remoteURL,omitempty" json:"remoteURL,omitempty"`
	DownloadURLs map[string]string `yaml:"downloadURLs,omitempty" json:"downloadURLs,omitempty"`
	HiddenTags   []string          `yaml:"hiddenTags,omitempty" json:"hiddenTags,omitempty"`
}

func (s *CustomSource) GetName() string {
	return s.Name
}

func (s *CustomSource) GetVersions() []string {
	versions := []string{"latest"}

	// Get tags from remote repo
	remoteURL := s.RemoteURL

	tags, err := gitutil.GetTagListFromRemote(remoteURL, true)
	if err != nil {
		// klog.Warningf("Get tag list from remote failed, err: %v", err)
		return versions
	}

	// Hidden timeout tags
	isHiddenTag := getIsHiddenTag(s.HiddenTags)
	for _, tag := range tags {
		if _, hidden := isHiddenTag[tag]; !hidden {
			versions = append(versions, tag)
		}
	}

	return versions
}

func (s *CustomSource) GetDownloadURL(ver string) (string, error) {
	return getArchiveDownloadURL(s.DownloadURLs, ver)
}

func getIsHiddenTag(hiddenTags []string) map[string]struct{} {
	isHiddenTag := make(map[string]struct{})

	for _, tag := range hiddenTags {
		isHiddenTag[tag] = struct{}{}
	}

	return isHiddenTag
}
