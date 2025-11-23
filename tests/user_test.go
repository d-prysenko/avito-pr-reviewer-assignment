package tests

import (
	"net/http"
	"net/url"
	"revass/internal/handler/users/request"
	"revass/internal/model"
	"revass/internal/service"
	"revass/internal/storage/repository"
	"testing"

	_ "github.com/lib/pq"

	"github.com/gavv/httpexpect/v2"
)

func TestUsersSetIsActive(t *testing.T) {
	db := setup(t)

	userRep := repository.NewUserRepository(db)
	teamRep := repository.NewTeamRepository(db)

	teamManager := service.NewTeamManager(userRep, teamRep)

	t.Run("/users/setIsActive Deactivate user", func(t *testing.T) {
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

		resp := e.POST("/users/setIsActive").
			WithJSON(request.SetIsActiveRequest{
				UserID:   "u1",
				IsActive: false,
			},
			).
			Expect().Status(http.StatusOK).
			JSON().Object()

		resp.ContainsKey("user")
		resp.Value("user").Object().Value("is_active").Boolean().IsFalse()
	})
}

func TestUsersSetIsActiveNotFound(t *testing.T) {
	setup(t)

	t.Run("/users/setIsActive Try to deactivate non existing user", func(t *testing.T) {
		u := url.URL{
			Scheme: "http",
			Host:   host,
		}

		e := httpexpect.Default(t, u.String())

		resp := e.POST("/users/setIsActive").
			WithJSON(request.SetIsActiveRequest{
				UserID:   "u1",
				IsActive: false,
			},
			).
			Expect().Status(http.StatusNotFound).
			JSON().Object()

		resp.Value("error").Object().Value("code").String().IsEqual("NOT_FOUND")
		resp.Value("error").Object().Value("message").String().IsEqual("Resource not found")
	})
}
