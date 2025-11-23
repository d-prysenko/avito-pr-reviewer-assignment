package tests

import (
	"net/http"
	"net/url"
	"revass/internal/handler/pullrequest/request"
	"revass/internal/model"
	"revass/internal/service"
	"revass/internal/storage/repository"
	"testing"

	_ "github.com/lib/pq"

	"github.com/gavv/httpexpect/v2"
)

func TestPullRequestCreate(t *testing.T) {
	db := setup(t)

	userRep := repository.NewUserRepository(db)
	teamRep := repository.NewTeamRepository(db)

	teamManager := service.NewTeamManager(userRep, teamRep)

	t.Run("/pullRequest/create Create pull request", func(t *testing.T) {
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

		resp := e.POST("/pullRequest/create").
			WithJSON(request.PullRequestCreateRequest{
				PullRequestID:   "pr-1001",
				PullRequestName: "Add search",
				AuthorID:        "u1",
			},
			).
			Expect().Status(http.StatusCreated).
			JSON().Object()

		resp.ContainsKey("pr")
		resp.Value("pr").Object().Value("assigned_reviewers").Array().Length().IsEqual(1)

		resp = e.POST("/pullRequest/create").
			WithJSON(request.PullRequestCreateRequest{
				PullRequestID:   "pr-1001",
				PullRequestName: "Add search",
				AuthorID:        "u1",
			},
			).
			Expect().Status(http.StatusConflict).
			JSON().Object()

		resp.ContainsKey("error")
		resp.Value("error").Object().Value("code").String().IsEqual("PR_EXISTS")
		resp.Value("error").Object().Value("message").String().IsEqual("PR 'pr-1001' already exists")
	})
}

func TestPullRequestCreateWithInactiveUser(t *testing.T) {
	db := setup(t)

	userRep := repository.NewUserRepository(db)
	teamRep := repository.NewTeamRepository(db)

	teamManager := service.NewTeamManager(userRep, teamRep)

	t.Run("/pullRequest/create Create pull request with inactive user", func(t *testing.T) {
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
					IsActive: false,
				},
			},
		})

		u := url.URL{
			Scheme: "http",
			Host:   host,
		}

		e := httpexpect.Default(t, u.String())

		resp := e.POST("/pullRequest/create").
			WithJSON(request.PullRequestCreateRequest{
				PullRequestID:   "pr-1001",
				PullRequestName: "Add search",
				AuthorID:        "u1",
			},
			).
			Expect().Status(http.StatusCreated).
			JSON().Object()

		resp.ContainsKey("pr")
		resp.Value("pr").Object().Value("assigned_reviewers").IsNull()
	})
}

func TestPullRequestCreateWithNotExistingUser(t *testing.T) {
	setup(t)

	t.Run("/pullRequest/create Create pull request with not existing user", func(t *testing.T) {
		u := url.URL{
			Scheme: "http",
			Host:   host,
		}

		e := httpexpect.Default(t, u.String())

		resp := e.POST("/pullRequest/create").
			WithJSON(request.PullRequestCreateRequest{
				PullRequestID:   "pr-1001",
				PullRequestName: "Add search",
				AuthorID:        "u1",
			},
			).
			Expect().Status(http.StatusNotFound).
			JSON().Object()

			resp.ContainsKey("error")
			resp.Value("error").Object().Value("code").String().IsEqual("NOT_FOUND")
			resp.Value("error").Object().Value("message").String().IsEqual("Resource not found")
	})
}
