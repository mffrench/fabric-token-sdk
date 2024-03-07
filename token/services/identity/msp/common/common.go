/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package common

import (
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/view"
	"github.com/hyperledger-labs/fabric-token-sdk/token/driver"
)

// Resolver contains information about an identity and how to retrieve it.
type Resolver struct {
	Name         string `yaml:"name,omitempty"`
	EnrollmentID string
	Default      bool
	GetIdentity  GetIdentityFunc
	Remote       bool
}

type SigService interface {
	IsMe(view.Identity) bool
	RegisterSigner(identity view.Identity, signer driver.Signer, verifier driver.Verifier) error
	RegisterVerifier(identity view.Identity, v driver.Verifier) error
	RegisterAuditInfo(identity view.Identity, info []byte) error
	GetAuditInfo(id view.Identity) ([]byte, error)
}

type BinderService interface {
	Bind(longTerm view.Identity, ephemeral view.Identity) error
}
