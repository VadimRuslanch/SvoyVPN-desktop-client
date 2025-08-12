//go:build windows
package secrets

import "context"

type vault struct{}

func newVault() Vault { return &vault{} }
func (v *vault) Put(ctx context.Context, key, value string) error { /* TODO: DPAPI/CredMan */ return nil }
func (v *vault) Get(ctx context.Context, key string) (string, error) { return "", nil }

