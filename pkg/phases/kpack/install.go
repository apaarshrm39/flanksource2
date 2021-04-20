package kpack

import (
	"github.com/flanksource/karina/pkg/platform"
)

const (
	Namespace = "kpack"
)

func Deploy(p *platform.Platform) error {
	if p.Kpack.Disabled || &p.Kpack.ImageVersions == nil {
		return p.DeleteSpecs(Namespace, "kpack.yaml")
	}
	if err := p.CreateOrUpdateNamespace(Namespace, nil, nil); err != nil {
		return err
	}

	return p.ApplySpecs(Namespace, "kpack.yaml")
}
