package auth

import "golang.org/x/oauth2"

// Faceook Auth object
var Facebook = newConfig()

func newConfig() *oauth2.Config {
	c := &oauth2.Config{
		ClientID: "963765180338498",
		ClientSecret: "2b97bf0ed00ad8bd72d7bcea773e4f74",
		RedirectURL: "http://localhost:8080/auth/callback",
		Endpoint: oauth2.Endpoint{
			AuthURL: "https://www.facebook.com/dialog/oauth",
			TokenURL: "https://graph.facebook.com/oauth/access_token",
		},
		Scopes: []string{"email"},
	}

	return c
}
