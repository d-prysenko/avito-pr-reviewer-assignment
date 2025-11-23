package tests

import (
	"net/http"
	"net/url"
	"revass/internal/handler/team/request"
	"testing"

	_ "github.com/lib/pq"

	"github.com/gavv/httpexpect/v2"
)

const (
	host = "app-test:8080"
)

func TestTeamAddTeamAlreadyExists(t *testing.T) {
	setup(t)

	t.Run("/team/add Team already exists", func(t *testing.T) {
		u := url.URL{
			Scheme: "http",
			Host:   host,
		}

		e := httpexpect.Default(t, u.String())

		resp := e.POST("/team/add").
			WithJSON(request.TeamAddRequest{
				Name: "team 1",
				Members: []request.TeamMember{
					{
						UserID:   "u1",
						Username: "Alice",
						IsActive: true,
					},
					{
						UserID:   "u2",
						Username: "Bob",
						IsActive: false,
					},
				},
			}).
			Expect().Status(http.StatusCreated).
			JSON().Object()

		resp.ContainsKey("team")

		resp = e.POST("/team/add").
			WithJSON(request.TeamAddRequest{
				Name: "team 1",
				Members: []request.TeamMember{
					{
						UserID:   "u1",
						Username: "Alice",
						IsActive: true,
					},
					{
						UserID:   "u2",
						Username: "Bob",
						IsActive: false,
					},
				},
			}).
			Expect().Status(http.StatusBadRequest).
			JSON().Object()

		resp.ContainsKey("error")
		resp.Value("error").Object().Value("code").String().IsEqual("TEAM_EXISTS")
		resp.Value("error").Object().Value("message").String().IsEqual("Team 'team 1' already exists")

		resp = e.POST("/team/add").
			WithJSON(request.TeamAddRequest{
				Name: "team 2",
				Members: []request.TeamMember{
					{
						UserID:   "u1",
						Username: "Alice",
						IsActive: true,
					},
				},
			}).
			Expect().Status(http.StatusBadRequest).
			JSON().Object()

		resp.ContainsKey("error")
		resp.Value("error").Object().Value("code").String().IsEqual("USER_EXISTS")
		resp.Value("error").Object().Value("message").String().IsEqual("User 'u1' already exists")
	})
}

func TestTeamAddIncorrectInput(t *testing.T) {
	setup(t)

	t.Run("/team/add Incorrect input", func(t *testing.T) {
		u := url.URL{
			Scheme: "http",
			Host:   host,
		}

		e := httpexpect.Default(t, u.String())

		resp := e.POST("/team/add").
			WithJSON(request.TeamAddRequest{
				Name: "",
			}).
			Expect().Status(http.StatusBadRequest).
			JSON().Object()

		resp.ContainsKey("error")
		resp.Value("error").Object().Value("code").String().IsEqual("BAD_REQUEST")
		resp.Value("error").Object().Value("message").String().IsEqual("Bad request.")
	})
}
