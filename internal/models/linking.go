package models

import (
	"github.com/supabase/auth/internal/api/provider"
	"github.com/supabase/auth/internal/conf"
	"github.com/supabase/auth/internal/storage"
)

// GetAccountLinkingDomain returns a string that describes the account linking
// domain. An account linking domain describes a set of Identity entities that
// _should_ generally fall under the same User entity. It's just a runtime
// string, and is not typically persisted in the database. This value can vary
// across time.
func GetAccountLinkingDomain(provider string, ownLinkingDomains []string) string {
	_ = "STUB: not implemented"
	return ""
}

// when the provider ID is a SSO provider, then the linking
// domain is the provider itself i.e. there can only be one
// user + identity per identity provider

// otherwise, the linking domain is the default linking domain that
// links all accounts

type AccountLinkingDecision = int

const (
	AccountExists AccountLinkingDecision = iota
	CreateAccount
	LinkAccount
	MultipleAccounts
)

type AccountLinkingResult struct {
	Decision       AccountLinkingDecision
	User           *User
	Identities     []*Identity
	LinkingDomain  string
	CandidateEmail provider.Email
}

// DetermineAccountLinking uses the provided data and database state to compute a decision on whether:
// - A new User should be created (CreateAccount)
// - A new Identity should be created (LinkAccount) with a UserID pointing to an existing user account
// - Nothing should be done (AccountExists)
// - It's not possible to decide due to data inconsistency (MultipleAccounts) and the caller should decide
//
// Errors signal failure in processing only, like database access errors.
func DetermineAccountLinking(tx *storage.Connection, config *conf.GlobalConfiguration, emails []provider.Email, aud, providerName, sub string) (AccountLinkingResult, error) {
	_ = "STUB: not implemented"
	return *new(AccountLinkingResult), nil
}

// this is the linking domain for the new identity

// account exists

// we overwrite the email with the existing user's email since the user
// could have an empty email

// the identity does not exist, so we need to check if we should create a new account
// or link to an existing one

// if there are no verified emails, we always decide to create a new account

// look for similar identities and users based on email

// there can be multiple user accounts with the same email when is_sso_user is true
// so we just do not consider those similar user accounts

// Need to check if the new identity should be assigned to an
// existing user or to create a new user, according to the automatic
// linking rules

// now let's see if there are any existing and similar identities in
// the same linking domain

// no similarIdentities but a user with the same email exists
// so we link this new identity to the user
// TODO: Backfill the missing identity for the user

// this shouldn't happen since there is a partial unique index on (email and is_sso_user = false)

// there are no identities in the linking domain, we have to
// create a new identity and new user

// there is at least one identity in the linking domain let's do a
// sanity check to see if all of the identities in the domain share the
// same user ID

// ok this linking domain has more than one user account
// caller should decide what to do

// there's only one user ID in this linking domain, we can go on and
// create a new identity and link it to the existing account
