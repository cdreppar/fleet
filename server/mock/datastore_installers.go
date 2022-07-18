// Automatically generated by mockimpl. DO NOT EDIT!

package mock

import (
	"context"
	"io"

	"github.com/fleetdm/fleet/v4/server/fleet"
)

var _ fleet.InstallerStore = (*InstallerStore)(nil)

type GetFunc func(ctx context.Context, installer fleet.Installer) (io.ReadCloser, int64, error)

type PutFunc func(ctx context.Context, installer fleet.Installer) (string, error)

type InstallerStore struct {
	GetFunc        GetFunc
	GetFuncInvoked bool

	PutFunc        PutFunc
	PutFuncInvoked bool
}

func (s *InstallerStore) Get(ctx context.Context, installer fleet.Installer) (io.ReadCloser, int64, error) {
	s.GetFuncInvoked = true
	return s.GetFunc(ctx, installer)
}

func (s *InstallerStore) Put(ctx context.Context, installer fleet.Installer) (string, error) {
	s.PutFuncInvoked = true
	return s.PutFunc(ctx, installer)
}
