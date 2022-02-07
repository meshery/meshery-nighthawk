package config

import (
	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/meshes"
)

func getOperations(dev adapter.Operations) adapter.Operations {
	var adapterVersions []adapter.Version

	dev[PerfOperation] = &adapter.Operation{
		Type:                 int32(meshes.OpCategory_INSTALL),
		Description:          "Install Meshery Perf",
		Versions:             adapterVersions,
		AdditionalProperties: map[string]string{},
	}

	return dev
}
