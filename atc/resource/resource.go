package resource

import (
	"context"
	"path/filepath"

	"github.com/concourse/concourse/atc/runtime"

	"github.com/concourse/concourse/atc"
)

//go:generate counterfeiter . Resource

type Resource interface {
	Get(context.Context, runtime.ProcessSpec, runtime.Runnable) (runtime.VersionResult, error)
	Put(context.Context, runtime.ProcessSpec, runtime.Runnable) (runtime.VersionResult, error)
	Check(context.Context, runtime.ProcessSpec, runtime.Runnable) ([]atc.Version, error)
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
		Source:  source,
		Params:  params,
		Version: version,
	}
}

type resource struct {
	Source  atc.Source  `json:"source"`
	Params  atc.Params  `json:"params,omitempty"`
	Version atc.Version `json:"version,omitempty"`
}
