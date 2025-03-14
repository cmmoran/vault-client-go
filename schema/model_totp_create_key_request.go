// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

// TotpCreateKeyRequest struct for TotpCreateKeyRequest
type TotpCreateKeyRequest struct {
	// The name of the account associated with the key. Required if generate is true.
	AccountName string `json:"account_name,omitempty"`

	// The hashing algorithm used to generate the TOTP token. Options include SHA1, SHA256 and SHA512.
	Algorithm string `json:"algorithm,omitempty"`

	// The number of digits in the generated TOTP token. This value can either be 6 or 8.
	Digits int32 `json:"digits,omitempty"`

	// Determines if a QR code and url are returned upon generating a key. Only used if generate is true.
	Exported *bool `json:"exported,omitempty"`

	// Determines if a key should be generated by Vault or if a key is being passed from another service.
	Generate *bool `json:"generate,omitempty"`

	// The name of the key's issuing organization. Required if generate is true.
	Issuer string `json:"issuer,omitempty"`

	// The shared master key used to generate a TOTP token. Only used if generate is false.
	Key string `json:"key,omitempty"`

	// Determines the size in bytes of the generated key. Only used if generate is true.
	KeySize int32 `json:"key_size,omitempty"`

	// The length of time used to generate a counter for the TOTP token calculation.
	Period string `json:"period,omitempty"`

	// The pixel size of the generated square QR code. Only used if generate is true and exported is true. If this value is 0, a QR code will not be returned.
	QrSize int32 `json:"qr_size,omitempty"`

	// The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1. Only used if generate is true.
	Skew int32 `json:"skew,omitempty"`

	// A TOTP url string containing all of the parameters for key setup. Only used if generate is false.
	Url string `json:"url,omitempty"`
}
