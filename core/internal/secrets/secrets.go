package secrets

import "context"

type Vault interface {
  Put(ctx context.Context, key, value string) error
  Get(ctx context.Context, key string) (string, error)
}

func New() Vault { return newVault() }

