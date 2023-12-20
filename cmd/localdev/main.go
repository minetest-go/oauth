package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/minetest-go/oauth"
)

const BaseURL = "http://localhost:8080"

type Model struct {
	Logins map[oauth.ProviderType]string
}

var model = &Model{
	Logins: make(map[oauth.ProviderType]string),
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("cmd/localdev/index.html"))
	tmpl.ExecuteTemplate(w, "index.html", model)
}

func callback(w http.ResponseWriter, r *http.Request, user *oauth.OauthUserInfo) error {
	tmpl := template.Must(template.ParseFiles("cmd/localdev/callback.html"))
	tmpl.ExecuteTemplate(w, "callback.html", user)
	return nil
}

func setup_provider(path string, cfg *oauth.OAuthConfig) {
	fmt.Printf("Registering oauth provider %s on path '%s'\n", cfg.Provider, path)
	cfg.CallbackURL = fmt.Sprintf("%s%s", BaseURL, path)
	handler := oauth.NewHandler(callback, cfg)
	model.Logins[cfg.Provider] = handler.LoginURL()
	http.Handle(path, handler)
}

func main() {
	if os.Getenv("DISCORD_APP_ID") != "" {
		setup_provider("/api/oauth_callback/discord", &oauth.OAuthConfig{
			Provider: oauth.ProviderTypeDiscord,
			ClientID: os.Getenv("DISCORD_APP_ID"),
			Secret:   os.Getenv("DISCORD_APP_SECRET"),
		})
	}

	if os.Getenv("CDB_APP_ID") != "" {
		setup_provider("/api/oauth_callback/cdb", &oauth.OAuthConfig{
			Provider: oauth.ProviderTypeCDB,
			ClientID: os.Getenv("CDB_APP_ID"),
			Secret:   os.Getenv("CDB_APP_SECRET"),
		})
	}

	if os.Getenv("GITHUB_APP_ID") != "" {
		setup_provider("/api/oauth_callback/github", &oauth.OAuthConfig{
			Provider: oauth.ProviderTypeGithub,
			ClientID: os.Getenv("GITHUB_APP_ID"),
			Secret:   os.Getenv("GITHUB_APP_SECRET"),
		})
	}

	if os.Getenv("MESEHUB_APP_ID") != "" {
		setup_provider("/api/oauth_callback/mesehub", &oauth.OAuthConfig{
			Provider: oauth.ProviderTypeMesehub,
			ClientID: os.Getenv("MESEHUB_APP_ID"),
			Secret:   os.Getenv("MESEHUB_APP_SECRET"),
		})
	}

	http.HandleFunc("/", index)
	fmt.Printf("Starting local oauth dev env on %s\n", BaseURL)

	server := &http.Server{Addr: ":8080", Handler: nil}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
