// Code generated by sysl DO NOT EDIT.
package deps

import (
	"context"
	"time"
)

// DefaultDepsImpl ...
type DefaultDepsImpl struct {
}

// NewDefaultDepsImpl for Deps
func NewDefaultDepsImpl() *DefaultDepsImpl {
	return &DefaultDepsImpl{}
}

// GetApiDocsList Client
type GetApiDocsListClient struct {
}

// ServiceInterface for Deps
type ServiceInterface struct {
	GetApiDocsList func(ctx context.Context, req *GetApiDocsListRequest, client GetApiDocsListClient) (*ApiDoc, error)
}

// DownstreamConfig for Deps
type DownstreamConfig struct {
	ContextTimeout time.Duration `yaml:"contextTimeout"`
}
