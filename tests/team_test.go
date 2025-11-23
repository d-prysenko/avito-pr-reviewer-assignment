package tests

import (
	"net/http"
	"net/url"
	"revass/internal/handler/team/request"
	"revass/internal/model"
	"revass/internal/service"
	"revass/internal/storage/repository"
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
		resp.Value("error").Object().Value("message").String().IsEqual("Bad request")
	})
}

func TestTeamGet(t *testing.T) {
	db := setup(t)
	userRep := repository.NewUserRepository(db)
	teamRep := repository.NewTeamRepository(db)

	teamManager := service.NewTeamManager(userRep, teamRep)

	t.Run("/team/get Success", func(t *testing.T) {
		teamManager.AddTeamWithMembers(model.Team{
			Name: "team 1",
			Members: []*model.TeamMember{
				{
					UserID:   "u1",
					Username: "Alice",
					IsActive: true,
				},
				{
					UserID:   "u2",
					Username: "Bob",
					IsActive: true,
				},
			},
		})

		u := url.URL{
			Scheme: "http",
			Host:   host,
		}

		e := httpexpect.Default(t, u.String())

		resp := e.GET("/team/get").WithQuery("team_name", "team 1").
			Expect().Status(http.StatusOK).
			JSON().Object()

		resp.ContainsKey("team_name")
		resp.ContainsKey("members")
		resp.Value("members").Array().Length().IsEqual(2)
	})
}

func TestTeamGetNotFound(t *testing.T) {
	setup(t)

	t.Run("/team/get Not found", func(t *testing.T) {
		u := url.URL{
			Scheme: "http",
			Host:   host,
		}

		e := httpexpect.Default(t, u.String())

		resp := e.GET("/team/get").WithQuery("team_name", "team 1").
			Expect().Status(http.StatusNotFound).
			JSON().Object()

		resp.Value("error").Object().Value("code").String().IsEqual("NOT_FOUND")
		resp.Value("error").Object().Value("message").String().IsEqual("Resource not found")
	})
}
