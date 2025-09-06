package oauth

import (
	"context"
	"encoding/json"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOauthConfig *oauth2.Config

func Init(clientID, clientSecret, redirectURL string) {
	GoogleOauthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
	}
}

// BuildAuthURL returns URL to redirect user to Google consent page.
// state should be a cryptographically-random string stored in user session (to prevent CSRF).
func BuildAuthURL(state string) string {
	return GoogleOauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

// FetchUserInfo uses token to call Google userinfo endpoint
func FetchUserInfo(ctx context.Context, token *oauth2.Token) (map[string]interface{}, error) {
	client := GoogleOauthConfig.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userinfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userinfo); err != nil {
		return nil, err
	}
	return userinfo, nil
}

// Exchange code -> token
func Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	return GoogleOauthConfig.Exchange(ctx, code)
}
