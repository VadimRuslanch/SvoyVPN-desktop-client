//go:build darwin
package controller

import (
  "context"
  prof "github.com/you/vadim-desktop/core/api/profile"  ctrl "github.com/you/vadim-desktop/core/api/control")

type wgImpl struct{}

func newWGImpl() WG { return &wgImpl{} }

func (w *wgImpl) Up(ctx context.Context, p *prof.Profile, privateKey, psk string) error {
  // TODO: embed/run wireguard-go, request utun, assign routes/DNS/MTU, apply peers
  // TODO: optional PF rules для Kill-Switch
  return nil
}

func (w *wgImpl) Down(ctx context.Context, id string) error { return nil }

func (w *wgImpl) Status(ctx context.Context, id string) (*ctrl.Status, error) {
  return &ctrl.Status{State: ctrl.Status_CONNECTED}, nil
}

