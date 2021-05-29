package tim

import (
	tenSig "github.com/tencentyun/tls-sig-api-golang"
)

type SigVersion int

const (
	VerNew SigVersion = 0
	VerOld SigVersion = 1
)

type ServerOption interface {
	SetOption(*IMServer) error
}

type SignOption struct {
	sigVersion SigVersion
	privateKey string
}

func NewSignOption() *SignOption {
	return &SignOption{}
}

func (o *SignOption) SetSigVersion(ver SigVersion) *SignOption {
	o.sigVersion = ver
	return o
}

func (o *SignOption) SetPrivateKey(privateKey string) *SignOption {
	o.privateKey = privateKey
	return o
}

func (o *SignOption) SetOption(s *IMServer) error {
	switch o.sigVersion {
	case VerOld:
		if o.privateKey == "" {
			return ErrInvalidPriKey
		}
		sig, err := tenSig.GenerateUsersigWithExpire(o.privateKey, s.AppId, s.Identifier, int64(s.Expire))
		if err != nil {
			return err
		}
		s.Sig = sig
	case VerNew:
		var err error
		if s.Sig, err = s.userSig(); err != nil {
			return err
		}
	default:
		return ErrInvalidVer
	}

	return nil
}
