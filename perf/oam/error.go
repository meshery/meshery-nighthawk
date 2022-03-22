package oam

import "github.com/layer5io/meshkit/errors"

const (
	ErrLoadingComponentsCode = "1000"
)

func ErrLoadingComponents(err error) error {
	return errors.New(ErrLoadingComponentsCode, errors.Alert, []string{"could not load workloads or traits"}, []string{err.Error()}, []string{"the path to workloads or traits does not exist"}, []string{"make sure the path to workloads or traits exist"})
}
