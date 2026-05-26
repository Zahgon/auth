package crypto

import (
	"context"
	"errors"
	"regexp"

	"github.com/supabase/auth/internal/observability"
)

type HashCost = int

const (
	// DefaultHashCost represents the default
	// hashing cost for any hashing algorithm.
	DefaultHashCost HashCost = iota

	// QuickHashCosts represents the quickest
	// hashing cost for any hashing algorithm,
	// useful for tests only.
	QuickHashCost HashCost = iota

	Argon2Prefix         = "$argon2"
	FirebaseScryptPrefix = "$fbscrypt"
	FirebaseScryptKeyLen = 32 // Firebase uses AES-256 which requires 32 byte keys: https://pkg.go.dev/golang.org/x/crypto/scrypt#Key
)

// PasswordHashCost is the current pasword hashing cost
// for all new hashes generated with
// GenerateHashFromPassword.
var PasswordHashCost = DefaultHashCost

var (
	generateFromPasswordSubmittedCounter = observability.ObtainMetricCounter("gotrue_generate_from_password_submitted", "Number of submitted GenerateFromPassword hashing attempts")
	generateFromPasswordCompletedCounter = observability.ObtainMetricCounter("gotrue_generate_from_password_completed", "Number of completed GenerateFromPassword hashing attempts")
)

var (
	compareHashAndPasswordSubmittedCounter = observability.ObtainMetricCounter("gotrue_compare_hash_and_password_submitted", "Number of submitted CompareHashAndPassword hashing attempts")
	compareHashAndPasswordCompletedCounter = observability.ObtainMetricCounter("gotrue_compare_hash_and_password_completed", "Number of completed CompareHashAndPassword hashing attempts")
)

var ErrArgon2MismatchedHashAndPassword = errors.New("crypto: argon2 hash and password mismatch")
var ErrScryptMismatchedHashAndPassword = errors.New("crypto: fbscrypt hash and password mismatch")

// argon2HashRegexp https://github.com/P-H-C/phc-string-format/blob/master/phc-sf-spec.md#argon2-encoding
var argon2HashRegexp = regexp.MustCompile("^[$](?P<alg>argon2(d|i|id))[$]v=(?P<v>(16|19))[$]m=(?P<m>[0-9]+),t=(?P<t>[0-9]+),p=(?P<p>[0-9]+)(,keyid=(?P<keyid>[^,$]+))?(,data=(?P<data>[^$]+))?[$](?P<salt>[^$]*)[$](?P<hash>.*)$")
var fbscryptHashRegexp = regexp.MustCompile(`^\$fbscrypt\$v=(?P<v>[0-9]+),n=(?P<n>[0-9]+),r=(?P<r>[0-9]+),p=(?P<p>[0-9]+)(?:,ss=(?P<ss>[^,]+))?(?:,sk=(?P<sk>[^$]+))?\$(?P<salt>[^$]+)\$(?P<hash>.+)$`)

type Argon2HashInput struct {
	alg     string
	v       string
	memory  uint64
	time    uint64
	threads uint64
	keyid   string
	data    string
	salt    []byte
	rawHash []byte
}

type FirebaseScryptHashInput struct {
	v             string
	memory        uint64
	rounds        uint64
	threads       uint64
	saltSeparator []byte
	signerKey     []byte
	salt          []byte
	rawHash       []byte
}

// See: https://github.com/firebase/scrypt for implementation
func ParseFirebaseScryptHash(hash string) (*FirebaseScryptHashInput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseArgon2Hash(hash string) (*Argon2HashInput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 1 GiB in KiB

func compareHashAndPasswordArgon2(ctx context.Context, hash, password string) error {
	_ = "STUB: not implemented"
	return nil
}

// #nosec G115

// #nosec G115

// #nosec G115

func compareHashAndPasswordFirebaseScrypt(ctx context.Context, hash, password string) error {
	_ = "STUB: not implemented"
	return nil
}

// #nosec G115

func firebaseScrypt(password, salt, signerKey, saltSeparator []byte, memCost, rounds, p uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// #nosec G115

// #nosec G407 -- Firebase scrypt requires deterministic IV for consistent results. See: JaakkoL/firebase-scrypt-python@master/firebasescrypt/firebasescrypt.py#L58

// CompareHashAndPassword compares the hash and
// password, returns nil if equal otherwise an error. Context can be used to
// cancel the hashing if the algorithm supports it.
func CompareHashAndPassword(ctx context.Context, hash, password string) error {
	_ = "STUB: not implemented"
	return nil
}

// assume bcrypt

// GenerateFromPassword generates a password hash from a
// password, using PasswordHashCost. Context can be used to cancel the hashing
// if the algorithm supports it.
func GenerateFromPassword(ctx context.Context, password string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GeneratePassword(requiredChars []string, length int) string {
	_ = "STUB: not implemented"
	return ""
}

// Add required characters

// Define a default character set for random generation (if needed)

// Fill the rest of the password

// Convert to byte slice for shuffling

// Secure shuffling
