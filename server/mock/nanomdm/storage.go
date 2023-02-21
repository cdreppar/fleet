// Automatically generated by mockimpl. DO NOT EDIT!

package mock

import (
	"context"
	"crypto/tls"
	"sync"

	"github.com/micromdm/nanomdm/mdm"
	"github.com/micromdm/nanomdm/storage"
)

var _ storage.AllStorage = (*Storage)(nil)

type StoreAuthenticateFunc func(r *mdm.Request, msg *mdm.Authenticate) error

type StoreTokenUpdateFunc func(r *mdm.Request, msg *mdm.TokenUpdate) error

type StoreUserAuthenticateFunc func(r *mdm.Request, msg *mdm.UserAuthenticate) error

type DisableFunc func(r *mdm.Request) error

type StoreCommandReportFunc func(r *mdm.Request, report *mdm.CommandResults) error

type RetrieveNextCommandFunc func(r *mdm.Request, skipNotNow bool) (*mdm.Command, error)

type ClearQueueFunc func(r *mdm.Request) error

type StoreBootstrapTokenFunc func(r *mdm.Request, msg *mdm.SetBootstrapToken) error

type RetrieveBootstrapTokenFunc func(r *mdm.Request, msg *mdm.GetBootstrapToken) (*mdm.BootstrapToken, error)

type RetrievePushInfoFunc func(p0 context.Context, p1 []string) (map[string]*mdm.Push, error)

type IsPushCertStaleFunc func(ctx context.Context, topic string, staleToken string) (bool, error)

type RetrievePushCertFunc func(ctx context.Context, topic string) (cert *tls.Certificate, staleToken string, err error)

type StorePushCertFunc func(ctx context.Context, pemCert []byte, pemKey []byte) error

type EnqueueCommandFunc func(ctx context.Context, id []string, cmd *mdm.Command) (map[string]error, error)

type HasCertHashFunc func(r *mdm.Request, hash string) (bool, error)

type EnrollmentHasCertHashFunc func(r *mdm.Request, hash string) (bool, error)

type IsCertHashAssociatedFunc func(r *mdm.Request, hash string) (bool, error)

type AssociateCertHashFunc func(r *mdm.Request, hash string) error

type RetrieveMigrationCheckinsFunc func(p0 context.Context, p1 chan<- interface{}) error

type RetrieveTokenUpdateTallyFunc func(ctx context.Context, id string) (int, error)

type Storage struct {
	StoreAuthenticateFunc        StoreAuthenticateFunc
	StoreAuthenticateFuncInvoked bool

	StoreTokenUpdateFunc        StoreTokenUpdateFunc
	StoreTokenUpdateFuncInvoked bool

	StoreUserAuthenticateFunc        StoreUserAuthenticateFunc
	StoreUserAuthenticateFuncInvoked bool

	DisableFunc        DisableFunc
	DisableFuncInvoked bool

	StoreCommandReportFunc        StoreCommandReportFunc
	StoreCommandReportFuncInvoked bool

	RetrieveNextCommandFunc        RetrieveNextCommandFunc
	RetrieveNextCommandFuncInvoked bool

	ClearQueueFunc        ClearQueueFunc
	ClearQueueFuncInvoked bool

	StoreBootstrapTokenFunc        StoreBootstrapTokenFunc
	StoreBootstrapTokenFuncInvoked bool

	RetrieveBootstrapTokenFunc        RetrieveBootstrapTokenFunc
	RetrieveBootstrapTokenFuncInvoked bool

	RetrievePushInfoFunc        RetrievePushInfoFunc
	RetrievePushInfoFuncInvoked bool

	IsPushCertStaleFunc        IsPushCertStaleFunc
	IsPushCertStaleFuncInvoked bool

	RetrievePushCertFunc        RetrievePushCertFunc
	RetrievePushCertFuncInvoked bool

	StorePushCertFunc        StorePushCertFunc
	StorePushCertFuncInvoked bool

	EnqueueCommandFunc        EnqueueCommandFunc
	EnqueueCommandFuncInvoked bool

	HasCertHashFunc        HasCertHashFunc
	HasCertHashFuncInvoked bool

	EnrollmentHasCertHashFunc        EnrollmentHasCertHashFunc
	EnrollmentHasCertHashFuncInvoked bool

	IsCertHashAssociatedFunc        IsCertHashAssociatedFunc
	IsCertHashAssociatedFuncInvoked bool

	AssociateCertHashFunc        AssociateCertHashFunc
	AssociateCertHashFuncInvoked bool

	RetrieveMigrationCheckinsFunc        RetrieveMigrationCheckinsFunc
	RetrieveMigrationCheckinsFuncInvoked bool

	RetrieveTokenUpdateTallyFunc        RetrieveTokenUpdateTallyFunc
	RetrieveTokenUpdateTallyFuncInvoked bool

	mu sync.Mutex
}

func (s *Storage) StoreAuthenticate(r *mdm.Request, msg *mdm.Authenticate) error {
	s.mu.Lock()
	s.StoreAuthenticateFuncInvoked = true
	s.mu.Unlock()
	return s.StoreAuthenticateFunc(r, msg)
}

func (s *Storage) StoreTokenUpdate(r *mdm.Request, msg *mdm.TokenUpdate) error {
	s.mu.Lock()
	s.StoreTokenUpdateFuncInvoked = true
	s.mu.Unlock()
	return s.StoreTokenUpdateFunc(r, msg)
}

func (s *Storage) StoreUserAuthenticate(r *mdm.Request, msg *mdm.UserAuthenticate) error {
	s.mu.Lock()
	s.StoreUserAuthenticateFuncInvoked = true
	s.mu.Unlock()
	return s.StoreUserAuthenticateFunc(r, msg)
}

func (s *Storage) Disable(r *mdm.Request) error {
	s.mu.Lock()
	s.DisableFuncInvoked = true
	s.mu.Unlock()
	return s.DisableFunc(r)
}

func (s *Storage) StoreCommandReport(r *mdm.Request, report *mdm.CommandResults) error {
	s.mu.Lock()
	s.StoreCommandReportFuncInvoked = true
	s.mu.Unlock()
	return s.StoreCommandReportFunc(r, report)
}

func (s *Storage) RetrieveNextCommand(r *mdm.Request, skipNotNow bool) (*mdm.Command, error) {
	s.mu.Lock()
	s.RetrieveNextCommandFuncInvoked = true
	s.mu.Unlock()
	return s.RetrieveNextCommandFunc(r, skipNotNow)
}

func (s *Storage) ClearQueue(r *mdm.Request) error {
	s.mu.Lock()
	s.ClearQueueFuncInvoked = true
	s.mu.Unlock()
	return s.ClearQueueFunc(r)
}

func (s *Storage) StoreBootstrapToken(r *mdm.Request, msg *mdm.SetBootstrapToken) error {
	s.mu.Lock()
	s.StoreBootstrapTokenFuncInvoked = true
	s.mu.Unlock()
	return s.StoreBootstrapTokenFunc(r, msg)
}

func (s *Storage) RetrieveBootstrapToken(r *mdm.Request, msg *mdm.GetBootstrapToken) (*mdm.BootstrapToken, error) {
	s.mu.Lock()
	s.RetrieveBootstrapTokenFuncInvoked = true
	s.mu.Unlock()
	return s.RetrieveBootstrapTokenFunc(r, msg)
}

func (s *Storage) RetrievePushInfo(p0 context.Context, p1 []string) (map[string]*mdm.Push, error) {
	s.mu.Lock()
	s.RetrievePushInfoFuncInvoked = true
	s.mu.Unlock()
	return s.RetrievePushInfoFunc(p0, p1)
}

func (s *Storage) IsPushCertStale(ctx context.Context, topic string, staleToken string) (bool, error) {
	s.mu.Lock()
	s.IsPushCertStaleFuncInvoked = true
	s.mu.Unlock()
	return s.IsPushCertStaleFunc(ctx, topic, staleToken)
}

func (s *Storage) RetrievePushCert(ctx context.Context, topic string) (cert *tls.Certificate, staleToken string, err error) {
	s.mu.Lock()
	s.RetrievePushCertFuncInvoked = true
	s.mu.Unlock()
	return s.RetrievePushCertFunc(ctx, topic)
}

func (s *Storage) StorePushCert(ctx context.Context, pemCert []byte, pemKey []byte) error {
	s.mu.Lock()
	s.StorePushCertFuncInvoked = true
	s.mu.Unlock()
	return s.StorePushCertFunc(ctx, pemCert, pemKey)
}

func (s *Storage) EnqueueCommand(ctx context.Context, id []string, cmd *mdm.Command) (map[string]error, error) {
	s.mu.Lock()
	s.EnqueueCommandFuncInvoked = true
	s.mu.Unlock()
	return s.EnqueueCommandFunc(ctx, id, cmd)
}

func (s *Storage) HasCertHash(r *mdm.Request, hash string) (bool, error) {
	s.mu.Lock()
	s.HasCertHashFuncInvoked = true
	s.mu.Unlock()
	return s.HasCertHashFunc(r, hash)
}

func (s *Storage) EnrollmentHasCertHash(r *mdm.Request, hash string) (bool, error) {
	s.mu.Lock()
	s.EnrollmentHasCertHashFuncInvoked = true
	s.mu.Unlock()
	return s.EnrollmentHasCertHashFunc(r, hash)
}

func (s *Storage) IsCertHashAssociated(r *mdm.Request, hash string) (bool, error) {
	s.mu.Lock()
	s.IsCertHashAssociatedFuncInvoked = true
	s.mu.Unlock()
	return s.IsCertHashAssociatedFunc(r, hash)
}

func (s *Storage) AssociateCertHash(r *mdm.Request, hash string) error {
	s.mu.Lock()
	s.AssociateCertHashFuncInvoked = true
	s.mu.Unlock()
	return s.AssociateCertHashFunc(r, hash)
}

func (s *Storage) RetrieveMigrationCheckins(p0 context.Context, p1 chan<- interface{}) error {
	s.mu.Lock()
	s.RetrieveMigrationCheckinsFuncInvoked = true
	s.mu.Unlock()
	return s.RetrieveMigrationCheckinsFunc(p0, p1)
}

func (s *Storage) RetrieveTokenUpdateTally(ctx context.Context, id string) (int, error) {
	s.mu.Lock()
	s.RetrieveTokenUpdateTallyFuncInvoked = true
	s.mu.Unlock()
	return s.RetrieveTokenUpdateTallyFunc(ctx, id)
}
