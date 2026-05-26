package api

import "net/http"

type ProviderSettings struct {
	AnonymousUsers bool `json:"anonymous_users"`
	Apple          bool `json:"apple"`
	Azure          bool `json:"azure"`
	Bitbucket      bool `json:"bitbucket"`
	Discord        bool `json:"discord"`
	Facebook       bool `json:"facebook"`
	Snapchat       bool `json:"snapchat"`
	Figma          bool `json:"figma"`
	Fly            bool `json:"fly"`
	GitHub         bool `json:"github"`
	GitLab         bool `json:"gitlab"`
	Google         bool `json:"google"`
	Keycloak       bool `json:"keycloak"`
	Kakao          bool `json:"kakao"`
	Linkedin       bool `json:"linkedin"`
	LinkedinOIDC   bool `json:"linkedin_oidc"`
	Notion         bool `json:"notion"`
	Spotify        bool `json:"spotify"`
	Slack          bool `json:"slack"`
	SlackOIDC      bool `json:"slack_oidc"`
	WorkOS         bool `json:"workos"`
	Twitch         bool `json:"twitch"`
	Twitter        bool `json:"twitter"`
	Email          bool `json:"email"`
	Phone          bool `json:"phone"`
	Zoom           bool `json:"zoom"`
}

type Settings struct {
	ExternalProviders ProviderSettings `json:"external"`
	DisableSignup     bool             `json:"disable_signup"`
	MailerAutoconfirm bool             `json:"mailer_autoconfirm"`
	PhoneAutoconfirm  bool             `json:"phone_autoconfirm"`
	SmsProvider       string           `json:"sms_provider"`
	SAMLEnabled       bool             `json:"saml_enabled"`
	PasskeysEnabled   bool             `json:"passkeys_enabled"`
}

func (a *API) Settings(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}
