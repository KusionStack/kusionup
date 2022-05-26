package commands

import "github.com/KusionStack/kusionup/pkg/sources/custom"

const DefaultCustomSourceConfigFilename = "custom_sources.yaml"

// CustomSources is a array of custom sources.
type CustomSources struct {
	Sources []custom.CustomSource `yaml:"sources,omitempty" json:"sources,omitempty"`
}
