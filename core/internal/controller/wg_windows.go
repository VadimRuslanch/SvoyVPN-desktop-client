//go:build windows
package controller

import (
  "context"
  prof "github.com/you/vadim-desktop/core/api/profile"  ctrl "github.com/you/vadim-desktop/core/api/control")

type wgImpl struct{}

func newWGImpl() WG { return &wgImpl{} }

func (w *wgImpl) Up(ctx context.Context, p *prof.Profile, privateKey, psk string) error {
  // TODO: ensure Wintun, create adapter, assign IP/DNS/MTU, set keys/peers via wgctrl
  // TODO: optional WFP rules для Kill-Switch
  return nil
}

func (w *wgImpl) Down(ctx context.Context, id string) error {
  // TODO: remove routes, bring adapter down, cleanup
  return nil
}

func (w *wgImpl) Status(ctx context.Context, id string) (*ctrl.Status, error) {
  // TODO: query latest handshake, bytes via wgctrl
  return &ctrl.Status{State: ctrl.Status_CONNECTED}, nil
}

