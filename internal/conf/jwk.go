package conf

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

type JwtKeysDecoder map[string]JwkInfo

type JwkInfo struct {
	PublicKey  jwk.Key `json:"public_key"`
	PrivateKey jwk.Key `json:"private_key"`
}

// Decode implements the Decoder interface
func (j *JwtKeysDecoder) Decode(value string) error { _ = "STUB: not implemented"; return nil }

func (j *JwtKeysDecoder) decodeKey(config JwtKeysDecoder, key []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *JwtKeysDecoder) decodePrivateKey(
	config JwtKeysDecoder,
	privJwk jwk.Key,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *JwtKeysDecoder) decodePublicKey(
	config JwtKeysDecoder,
	privJwk jwk.Key,
	pubJwk jwk.Key,
) error {
	_ = "STUB: not implemented"
	// all public keys should have the the use claim set to 'sig
	return nil
}

// all public keys should only have 'verify' set as the key_ops

func (j *JwtKeysDecoder) Validate() error {
	_ = "STUB: not implemented"
	// Validate performs _minimal_ checks if the data stored in the key are valid.
	// By minimal, we mean that it does not check if the key is valid for use in
	// cryptographic operations. For example, it does not check if an RSA key's
	// `e` field is a valid exponent, or if the `n` field is a valid modulus.
	// Instead, it checks for things such as the _presence_ of some required fields,
	// or if certain keys' values are of particular length.
	//
	// Note that depending on the underlying key type, use of this method requires
	// that multiple fields in the key are properly populated. For example, an EC
	// key's "x", "y" fields cannot be validated unless the "crv" field is populated first.
	return nil
}

// symmetric keys don't have public keys

func GetSigningJwk(config *JWTConfiguration) (jwk.Key, error) {
	_ = "STUB: not implemented"
	return *new(jwk.Key), nil
}

// the private JWK with key_ops "sign" should be used as the signing key

func GetSigningKey(k jwk.Key) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func GetSigningAlg(k jwk.Key) jwt.SigningMethod {
	_ = "STUB: not implemented"
	return *new(jwt.SigningMethod)
}

// return HS256 to preserve existing behaviour

func FindPublicKeyByKid(kid string, config *JWTConfiguration) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// don't return error, as a fallback key might be used
