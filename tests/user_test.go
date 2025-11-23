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
	"github.com/stretchr/testify/require"

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

func TestUsersGetReview(t *testing.T) {
	db := setup(t)

	userRep := repository.NewUserRepository(db)
	teamRep := repository.NewTeamRepository(db)
	prRep := repository.NewPRRepository(db)

	teamManager := service.NewTeamManager(userRep, teamRep)

	t.Run("/users/getReview Get users assigned prs", func(t *testing.T) {
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
				{
					UserID:   "u3",
					Username: "Margo",
					IsActive: true,
				},
			},
		})

		err := prRep.CreatePR("pr-1001", "Add search", "u1")
		require.NoError(t, err)
		err = prRep.CreatePR("pr-1002", "Add ping", "u1")
		require.NoError(t, err)
		err = prRep.CreatePR("pr-1003", "Add pong", "u1")
		require.NoError(t, err)

		err = prRep.AssignReviewer("pr-1001", "u2")
		require.NoError(t, err)
		err = prRep.AssignReviewer("pr-1002", "u2")
		require.NoError(t, err)

		err = prRep.AssignReviewer("pr-1003", "u3")
		require.NoError(t, err)

		u := url.URL{
			Scheme: "http",
			Host:   host,
		}

		e := httpexpect.Default(t, u.String())

		resp := e.GET("/users/getReview").WithQuery("user_id", "u2").
			Expect().Status(http.StatusOK).
			JSON().Object()

		resp.ContainsKey("pull_requests")
		resp.Value("pull_requests").Array().Length().IsEqual(2)

		resp = e.GET("/users/getReview").WithQuery("user_id", "u3").
			Expect().Status(http.StatusOK).
			JSON().Object()

		resp.ContainsKey("pull_requests")
		resp.Value("pull_requests").Array().Length().IsEqual(1)
	})
}
