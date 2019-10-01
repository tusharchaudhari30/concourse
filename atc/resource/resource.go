package resource

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/concourse/concourse/atc/runtime"

	"github.com/concourse/concourse/atc"
)

//go:generate counterfeiter . Resource

type Resource interface {
	Get(context.Context, runtime.ProcessSpec, runtime.Runnable) (runtime.VersionResult, error)
	Put(context.Context, runtime.ProcessSpec, runtime.Runnable) (runtime.VersionResult, error)
	Check(context.Context, runtime.ProcessSpec, runtime.Runnable) ([]atc.Version, error)
	Signature() (string, error)
}

type ResourceType string

type Metadata interface {
	Env() []string
}

func ResourcesDir(suffix string) string {
	return filepath.Join("/tmp", "build", suffix)
}

//func NewResource(spec runtime.ProcessSpec, configParams ConfigParams) *resource {
//	return &resource{
//		processSpec: spec,
//		params:      configParams,
//	}
//}

func NewResource(source atc.Source, params atc.Params, version atc.Version) Resource {
	return &resource{
		source:  source,
		params:  params,
		version: version,
	}
}

type resource struct {
	source  atc.Source  `json:"source"`
	params  atc.Params  `json:"params,omitempty"`
	version atc.Version `json:"version,omitempty"`
}

func (resource *resource) Signature() (string, error) {
	taskNameJSON, err := json.Marshal(resource)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sha256.Sum256(taskNameJSON)), nil
}
