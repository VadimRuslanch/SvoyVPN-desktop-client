package storage

import (
  "context"
  prof "github.com/you/vadim-desktop/core/api/profile"    )

type Store interface {
  SaveProfile(ctx context.Context, p *prof.Profile) (string, error)
  GetProfile(ctx context.Context, id string) (*prof.Profile, error)
}

func New() Store { return &fsStore{} }

type fsStore struct{}

func (s *fsStore) SaveProfile(ctx context.Context, p *prof.Profile) (string, error) {
  // TODO: encrypt+persist (без приватных ключей)
  if p.Id == "" { p.Id = "demo-profile" } // временно
  return p.Id, nil
}

func (s *fsStore) GetProfile(ctx context.Context, id string) (*prof.Profile, error) {
  return &prof.Profile{Id: id, Name: "Demo"}, nil
}

type SecretRefs struct {
  PrivateKey   string
  PresharedKey string
}

func FromWGQuick(raw string) (*prof.Profile, *SecretRefs, error) {
  // TODO: parse INI, validate, split secrets
  return &prof.Profile{Name: "Imported", Iface: &prof.Interface{Address: "10.0.0.2/32", Dns: "1.1.1.1", Mtu: 1420}}, &SecretRefs{PrivateKey: "PRIVATE_KEY"}, nil
}

