package api

import (
	"regexp"

	"github.com/gofrs/uuid"
	"github.com/supabase/auth/internal/models"
	"github.com/supabase/auth/internal/storage"
)

const (
	PKCEPrefix                    = "pkce_"
	MinCodeChallengeLength        = 43
	MaxCodeChallengeLength        = 128
	InvalidPKCEParamsErrorMessage = "PKCE flow requires code_challenge_method and code_challenge"
)

var codeChallengePattern = regexp.MustCompile("^[a-zA-Z._~0-9-]+$")

func isValidCodeChallenge(codeChallenge string) (bool, error) {
	_ = "STUB: not implemented"
	// See RFC 7636 Section 4.2: https://www.rfc-editor.org/rfc/rfc7636#section-4.2
	return false, nil
}

func addFlowPrefixToToken(token string, flowType models.FlowType) string {
	_ = "STUB: not implemented"
	return ""
}

func issueAuthCode(tx *storage.Connection, user *models.User, authenticationMethod models.AuthenticationMethod) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func isPKCEFlow(flowType models.FlowType) bool { _ = "STUB: not implemented"; return false }

func isImplicitFlow(flowType models.FlowType) bool { _ = "STUB: not implemented"; return false }

func validatePKCEParams(codeChallengeMethod, codeChallenge string) error {
	_ = "STUB: not implemented"
	return nil
}

// if both params are empty, just return nil

func getFlowFromChallenge(codeChallenge string) models.FlowType {
	_ = "STUB: not implemented"
	return *new(models.FlowType)
}

func generateFlowState(tx *storage.Connection, providerType string, authenticationMethod models.AuthenticationMethod, codeChallengeMethod string, codeChallenge string, userID *uuid.UUID) (*models.FlowState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
