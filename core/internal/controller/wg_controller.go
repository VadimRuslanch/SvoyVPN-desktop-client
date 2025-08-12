package controller

import (
  "context"
  prof "github.com/you/vadim-desktop/core/api/profile"      ctrl "github.com/you/vadim-desktop/core/api/control"    )

type WG interface {
  Up(ctx context.Context, p *prof.Profile, privateKey string, psk string) error
  Down(ctx context.Context, profileID string) error
  Status(ctx context.Context, profileID string) (*ctrl.Status, error)
}

func NewWG() WG { return newWGImpl() }

