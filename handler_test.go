package oauth_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/minetest-go/oauth"
	"github.com/stretchr/testify/assert"
)

func TestLoginURL(t *testing.T) {

	cfg := &oauth.OAuthConfig{
		Provider:    oauth.ProviderTypeGithub,
		ClientID:    "1234",
		Secret:      "much-secret-very-secure",
		CallbackURL: "http://localhost:8080/api/oauth_callback/github",
	}

	cb := func(w http.ResponseWriter, r *http.Request, user_info *oauth.OauthUserInfo) error {
		// not testable (yet)
		return nil
	}

	handler := oauth.NewHandler(cb, cfg)
	assert.NotNil(t, handler)

	url := handler.LoginURL()
	assert.Contains(t, url, fmt.Sprintf("client_id=%s", cfg.ClientID))
}
