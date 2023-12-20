package oauth_test

import (
	"fmt"
	"net/http"

	"github.com/minetest-go/oauth"
)

func ExampleOauthHandler() {
	// OAuth config
	cfg := &oauth.OAuthConfig{
		Provider:    oauth.ProviderTypeGithub,
		ClientID:    "1234",
		Secret:      "much-secret-very-secure",
		CallbackURL: "http://localhost:8080/api/oauth_callback/github",
	}

	// user callback handler
	cb := func(w http.ResponseWriter, r *http.Request, user_info *oauth.OauthUserInfo) error {
		//TODO: register / login user here
		fmt.Printf("ID: %s, Name: %s, AvatarURL: %s", user_info.ExternalID, user_info.Name, user_info.AvatarURL)
		return nil
	}

	// create a new oauth handler instance
	handler := oauth.NewHandler(cb, cfg)

	// register handler on above callback-url
	http.Handle("/api/oauth_callback/github", handler)

	// navigate the user to the login-url
	fmt.Printf("Login-URL: '%s'", handler.LoginURL())

	// profit?
}
