package api

import (
	"context"
	"net/http"

	"github.com/supabase/auth/internal/conf"
	"github.com/supabase/auth/internal/models"
)

func sendJSON(w http.ResponseWriter, status int, obj interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func isAdmin(u *models.User, config *conf.GlobalConfiguration) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *API) requestAud(ctx context.Context, r *http.Request) string {
	_ = "STUB: not implemented"

	// First check for an audience in the header
	return ""
}

// Then check the token

// ignore the JWT's aud claim if the role is admin
// this is because anon, service_role never had an aud claim to begin with

// Finally, return the default if none of the above methods are successful

type RequestParams interface {
	AdminUserParams |
		AdminCustomOAuthProviderParams |
		CreateSSOProviderParams |
		EnrollFactorParams |
		GenerateLinkParams |
		IdTokenGrantParams |
		InviteParams |
		OtpParams |
		PKCEGrantParams |
		PasswordGrantParams |
		RecoverParams |
		RefreshTokenGrantParams |
		ResendConfirmationParams |
		SignupParams |
		SingleSignOnParams |
		SmsParams |
		Web3GrantParams |
		UserUpdateParams |
		VerifyFactorParams |
		VerifyParams |
		adminUserUpdateFactorParams |
		adminUserDeleteParams |
		captchaRequest |
		ChallengeFactorParams |

		struct {
			Email string `json:"email"`
			Phone string `json:"phone"`
		} |
		struct {
			Email string `json:"email"`
		}
}

// retrieveRequestParams is a generic method that unmarshals the request body into the params struct provided
func retrieveRequestParams[A RequestParams](r *http.Request, params *A) error {
	_ = "STUB: not implemented"
	return nil
}
