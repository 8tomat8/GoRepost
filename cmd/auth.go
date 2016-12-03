package main

// Thanks @qwertmax for this example
// (http://stackoverflow.com/questions/29359907/social-network-vk-auth-with-martini)

import (
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/yanple/vk_api"
)

var api vk_api.Api

func prepareMartini() *martini.ClassicMartini {
	m := martini.Classic()

	m.Get("/vk/auth", func(w http.ResponseWriter, r *http.Request) {
		authUrl, err := api.GetAuthUrl(
			"http://localhost:3000/vk/token",
			"5759292",
			"wall,offline")

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, authUrl, http.StatusFound)
	})

	m.Get("/vk/token", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		err := api.OAuth(
			"http://localhost:3000/vk/token", // redirect uri
			"00Noj1YI6BaXs7zkA5Cn",
			"5759292",
			code)
		if err != nil {
			panic(err)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	})

	m.Get("/", func(w http.ResponseWriter, r *http.Request) string {
		if api.AccessToken == "" {
			return "<a href='/vk/auth'>Авторизоваться</a>"
		}

		// Api have: AccessToken, UserId, ExpiresIn
		log.Println("[LOG] martini.go:48 ->", api.AccessToken)

		// Make query
		params := make(map[string]string)
		params["domain"] = "yanple"
		params["count"] = "1"

		strResp, err := api.Request("wall.get", params)
		if err != nil {
			panic(err)
		}
		return strResp
	})
	return m
}

func main() {
	prepareMartini().Run()
}
