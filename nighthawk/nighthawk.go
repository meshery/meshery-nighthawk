package nighthawk

import (
	"context"

	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/meshes"
	"github.com/layer5io/meshery-adapter-library/status"
	meshkitCfg "github.com/layer5io/meshkit/config"
	"github.com/layer5io/meshkit/logger"
	"github.com/layer5io/meshkit/models"
	"github.com/layer5io/meshkit/utils/events"
	"gopkg.in/yaml.v2"
)

type Nighthawk struct {
	adapter.Adapter
}

// New initializes Nighthawk handler.
func New(c meshkitCfg.Handler, l logger.Handler, kc meshkitCfg.Handler, ev *events.EventStreamer) adapter.Handler {
	return &Nighthawk{
		Adapter: adapter.Adapter{
			Config:            c,
			Log:               l,
			KubeconfigHandler: kc,
			EventStreamer:     ev,
		},
	}
}

// ApplyOperation applies the operation on nighthawk
func (nighthawk *Nighthawk) ApplyOperation(ctx context.Context, opReq adapter.OperationRequest) error {
	err := nighthawk.CreateKubeconfigs(opReq.K8sConfigs)
	if err != nil {
		return err
	}
	kubeConfigs := opReq.K8sConfigs
	operations := make(adapter.Operations)
	requestedVersion := adapter.Version(opReq.Version)
	err = nighthawk.Config.GetObject(adapter.OperationsKey, &operations)
	if err != nil {
		return err
	}

	e := &meshes.EventsResponse{
		OperationId:   opReq.OperationID,
		Summary:       status.Deploying,
		Details:       "Operation is not supported",
		Component:     internalconfig.ServerConfig["type"],
		ComponentName: internalconfig.ServerConfig["name"],
	}
	return nil
}

// CreateKubeconfigs creates and writes passed kubeconfig onto the filesystem
func (nighthawk *Nighthawk) CreateKubeconfigs(kubeconfigs []string) error {
	var errs = make([]error, 0)
	for _, kubeconfig := range kubeconfigs {
		kconfig := models.Kubeconfig{}
		err := yaml.Unmarshal([]byte(kubeconfig), &kconfig)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		// To have control over what exactly to take in on kubeconfig
		nighthawk.KubeconfigHandler.SetKey("kind", kconfig.Kind)
		nighthawk.KubeconfigHandler.SetKey("apiVersion", kconfig.APIVersion)
		nighthawk.KubeconfigHandler.SetKey("current-context", kconfig.CurrentContext)
		err = nighthawk.KubeconfigHandler.SetObject("preferences", kconfig.Preferences)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		err = nighthawk.KubeconfigHandler.SetObject("clusters", kconfig.Clusters)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		err = 
// CreateKubeconfigs creates and writes passed kubeconfig onto the filesystem
func (nighthawk *Nighthawk) CreateKubeconfigs(kubeconfigs []string) error {
	var errs = make([]error, 0)
	for _, kubeconfig := range kubeconfigs {
		kconfig := models.Kubeconfig{}
		err := yaml.Unmarshal([]byte(kubeconfig), &kconfig)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		// To have control over what exactly to take in on kubeconfig
		nighthawk.KubeconfigHandler.SetKey("kind", kconfig.Kind)
		nighthawk.KubeconfigHandler.SetKey("apiVersion", kconfig.APIVersion)
		nighthawk.KubeconfigHandler.SetKey("current-context", kconfig.CurrentContext)
		err = nighthawk.KubeconfigHandler.SetObject("preferences", kconfig.Preferences)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		err = nighthawk.KubeconfigHandler.SetObject("clusters", kconfig.Clusters)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		err = nighthawk.KubeconfigHandler.SetObject("users", kconfig.Users)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		err = nighthawk.KubeconfigHandler.SetObject("contexts", kconfig.Contexts)
		if err != nil {
			errs = append(errs, err)
			continue
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return mergeErrors(errs)
}
.KubeconfigHandler.SetObject("users", kconfig.Users)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		err = nighthawk.KubeconfigHandler.SetObject("contexts", kconfig.Contexts)
		if err != nil {
			errs = append(errs, err)
			continue
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return mergeErrors(errs)
}
