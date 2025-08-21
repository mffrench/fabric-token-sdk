/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package validator

type ValidatorOptions struct {
	ValidateSBContextFunc []ValidateSBContextFunc
	ValidateTransferFunc  []ValidateTransferFunc
	// ValidateIssueFunc     []ValidateIssueFunc
}

func CompileOpts(opts ...ValidatorOption) (*ValidatorOptions, error) {
	txOptions := &ValidatorOptions{}
	for _, opt := range opts {
		if err := opt(txOptions); err != nil {
			return nil, err
		}
	}
	return txOptions, nil
}

type ValidatorOption func(*ValidatorOptions) error

func WithValidateSBContextFuncs(sbCtxFuncs []ValidateSBContextFunc) ValidatorOption {
	return func(o *ValidatorOptions) error {
		o.ValidateSBContextFunc = sbCtxFuncs
		return nil
	}
}

func WithValidateTransferFuncs(trFuncs []ValidateTransferFunc) ValidatorOption {
	return func(o *ValidatorOptions) error {
		o.ValidateTransferFunc = trFuncs
		return nil
	}
}

// func WithIssueTransferFuncs(isFuncs []ValidateIssueFunc) ValidatorOption {
// 	return func(o *ValidatorOptions) error {
// 		o.ValidateIssueFunc = isFuncs
// 		return nil
// 	}
// }
