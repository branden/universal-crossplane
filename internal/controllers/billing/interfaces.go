package billing

import (
	"context"

	v1 "k8s.io/api/core/v1"
)

// Registerer can register usage of universal-crossplane with idempotent calls.
type Registerer interface {
	Register(ctx context.Context, secret *v1.Secret, uid string) (string, error)
	Verify(token, uid string) (bool, error)
}

// NewNopRegisterer returns a Registerer that does nothing.
func NewNopRegisterer() NopRegisterer {
	return NopRegisterer{}
}

// NopRegisterer implements Registerer and does nothing.
type NopRegisterer struct{}

// Register does nothing.
func (np NopRegisterer) Register(_ context.Context, _ *v1.Secret, _ string) (string, error) {
	return "", nil
}

// Verify does nothing.
func (np NopRegisterer) Verify(_, _ string) (bool, error) {
	return true, nil
}