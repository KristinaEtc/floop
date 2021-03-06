package handlers

import (
	"fmt"

	"github.com/d3sw/floop/types"
)

// EchoHandler implements a Handler that simply echoes back the input
type EchoHandler struct{}

func (lc *EchoHandler) Init(*types.HandlerConfig) error {
	return nil
}

// Handle echos back input data before process starts.  The config is the normalized
// config built using data from the child process.  This may be different from the one
// used in Init
func (lc *EchoHandler) Handle(event *types.Event, conf *types.HandlerConfig) (map[string]interface{}, error) {
	fmt.Printf("[Echo] phase=%s %+v\n", event.Type, event.Data)
	return nil, nil
}
