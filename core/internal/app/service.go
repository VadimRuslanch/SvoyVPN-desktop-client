package app

import (
    "context"
    ctrl "github.com/you/vadim-desktop/core/api/control"
    prof "github.com/you/vadim-desktop/core/api/profile"
    "github.com/you/vadim-desktop/core/internal/storage"        
    "github.com/you/vadim-desktop/core/internal/secrets"
    "github.com/you/vadim-desktop/core/internal/controller"
)

type Service struct {
  ctrl.UnimplementedVadimControlServer
  store storage.Store
  sec   secrets.Vault
  wg    controller.WG
}

func NewService() *Service {
  return &Service{
    store: storage.New(),
    sec:   secrets.New(),
    wg:    controller.NewWG(),
  }
}

func (s *Service) ImportProfile(ctx context.Context, r *ctrl.ImportRequest) (*ctrl.ImportResponse, error) {
  profObj, secretRefs, err := storage.FromWGQuick(r.RawConf)
  if err != nil { return nil, err }
  id, err := s.store.SaveProfile(ctx, profObj)
  if err != nil { return nil, err }
  if err := s.sec.Put(ctx, "wg:"+id+":private", secretRefs.PrivateKey); err != nil { return nil, err }
  if secretRefs.PresharedKey != "" { _ = s.sec.Put(ctx, "wg:"+id+":psk", secretRefs.PresharedKey) }
  return &ctrl.ImportResponse{ProfileId: id}, nil
}

func (s *Service) StartTunnel(ctx context.Context, r *ctrl.StartRequest) (*ctrl.Status, error) {
  profObj, err := s.store.GetProfile(ctx, r.ProfileId)
  if err != nil { return nil, err }
  priv, _ := s.sec.Get(ctx, "wg:"+r.ProfileId+":private")
  psk, _  := s.sec.Get(ctx, "wg:"+r.ProfileId+":psk")
  if err := s.wg.Up(ctx, profObj, priv, psk); err != nil {
    return &ctrl.Status{State: ctrl.Status_ERROR, Message: err.Error()}, nil
  }
  return &ctrl.Status{State: ctrl.Status_CONNECTING}, nil
}

func (s *Service) StopTunnel(ctx context.Context, r *ctrl.StopRequest) (*ctrl.Status, error) {
  if err := s.wg.Down(ctx, r.ProfileId); err != nil {
    return &ctrl.Status{State: ctrl.Status_ERROR, Message: err.Error()}, nil
  }
  return &ctrl.Status{State: ctrl.Status_DISCONNECTED}, nil
}

func (s *Service) GetStatus(ctx context.Context, id *ctrl.ProfileId) (*ctrl.Status, error) {
  st, err := s.wg.Status(ctx, id.Id)
  if err != nil { return &ctrl.Status{State: ctrl.Status_ERROR, Message: err.Error()}, nil }
  return st, nil
}

// dummy to avoid unused import; remove when implemented fully
var _ = prof.Profile{}

