package models

type NwClient struct {
	ClientId              string            `json:"clientId"`
	ClientSecret          string            `json:"clientSecret"`
	DisplayName           string            `json:"displayName"`
	Scopes                []string          `json:"scopes"`
	AuthorizedGrantTypes  []string          `json:"authorizedGrantTypes"`
	WebServerRedirectUris []string          `json:"webServerRedirectUris"`
	AccessTokenValidity   interface{}       `json:"accessTokenValidity"`
	RefreshTokenValidity  interface{}       `json:"refreshTokenValidity"`
	AutoApproveScopes     []string          `json:"autoApproveScopes"`
	AutoApprove           bool              `json:"autoApprove"`
	AdditionalInformation map[string]string `json:"additionalInformation"`
}
