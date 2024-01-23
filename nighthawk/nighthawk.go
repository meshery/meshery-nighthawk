package nighthawk

import (
	"github.com/layer5io/meshery-adapter-library/adapter"
	meshkitCfg "github.com/layer5io/meshkit/config"
	"github.com/layer5io/meshkit/logger"
	"github.com/layer5io/meshkit/utils/events"
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
