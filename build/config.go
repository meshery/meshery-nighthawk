package build

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshkit/utils"
	"github.com/layer5io/meshkit/utils/manifests"
)

var DefaultGenerationMethod string
var DefaultGenerationURL string
var LatestVersion string
var MeshModelPath string
var AllVersions []string

const Component = "Nighthawk"

var MeshModelConfig = adapter.MeshModelConfig{ //Move to build/config.go
	Category: "Observability and Analysis",
	Metadata: map[string]interface{}{},
}

// NewConfig creates the configuration for creating components
func NewConfig(version string) manifests.Config {
	return manifests.Config{
		Name:        "meshery-nighthawk",
		Type:        Component,
		MeshVersion: version,
		CrdFilter: manifests.NewCueCrdFilter(manifests.ExtractorPaths{
			NamePath:    "spec.names.kind",
			IdPath:      "spec.names.kind",
			VersionPath: "spec.versions[0].name",
			GroupPath:   "spec.group",
			SpecPath:    "spec.versions[0].schema.openAPIV3Schema"}, false),
		ExtractCrds: func(manifest string) []string {
			crds := strings.Split(manifest, "---")
			return crds
		},
	}
}

func init() {
	//Initialize Metadata including logo svgs
	modelMetadata, _ := os.Open("./build/meshmodel_metadata.json")
	defer func() {
		if err := modelMetadata.Close(); err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}()
	byt, _ := io.ReadAll(modelMetadata)

	_ = json.Unmarshal(byt, &MeshModelConfig.Metadata)

	wd, _ := os.Getwd()
	MeshModelPath = filepath.Join(wd, "templates", "meshmodel", "components")
	AllVersions, _ = utils.GetLatestReleaseTagsSorted("meshery", "meshery-nighthawk")
	if len(AllVersions) == 0 {
		return
	}
	LatestVersion = AllVersions[len(AllVersions)-1]
	DefaultGenerationMethod = adapter.Manifests
	DefaultGenerationURL = "https://raw.githubusercontent.com/meshery/meshery-nighthawk/" + LatestVersion + "/manifests/charts/crds.yaml"
}
