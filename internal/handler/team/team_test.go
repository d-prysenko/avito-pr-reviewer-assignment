package team_test

import (
// "bytes"
// "encoding/json"
// "net/http"
// "net/http/httptest"
// "testing"

// "revass/internal/handler/team"
// "revass/internal/handler/team/request"

// "github.com/stretchr/testify/require"
)

// func TestTeamAddHandler(t *testing.T) {
// 	cases := []struct {
// 		name        string
// 		requestData request.TeamAddRequest
// 		respError   string
// 		mockError   error
// 	}{
// 		{
// 			name: "Success 1",
// 			requestData: request.TeamAddRequest{
// 				TeamName: "team 1",
// 				Members: []request.TeamMember{
// 					{
// 						UserID:   "u1",
// 						Username: "Alice",
// 						IsActive: true,
// 					},
// 					{
// 						UserID:   "u2",
// 						Username: "Bob",
// 						IsActive: true,
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "Success 2",
// 			requestData: request.TeamAddRequest{
// 				TeamName: "team 1",
// 				Members: []request.TeamMember{
// 					{
// 						UserID:   "u1",
// 						Username: "Alice",
// 						IsActive: true,
// 					},
// 					{
// 						UserID:   "u2",
// 						Username: "Bob",
// 						IsActive: false,
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "Success 3",
// 			requestData: request.TeamAddRequest{
// 				TeamName: "team 1",
// 				Members:  []request.TeamMember{},
// 			},
// 		},
// 	}

// 	for _, tc := range cases {
// 		t.Run(tc.name, func(t *testing.T) {

// 			handler := team.Add()

// 			input, err := json.Marshal(tc.requestData)
// 			require.Nil(t, err)

// 			req, err := http.NewRequest(http.MethodPost, "/team/add", bytes.NewReader(input))
// 			require.NoError(t, err)

// 			rr := httptest.NewRecorder()
// 			handler.ServeHTTP(rr, req)

// 			res := rr.Result()

// 			require.Equal(t, rr.Code, http.StatusCreated)
// 			require.Equal(t, res.Header.Get("Content-Type"), "application/json")
// 		})
// 	}
// }

// func TestTeamAddTeamAlreadyExists(t *testing.T) {
// 	cases := []struct {
// 		name        string
// 		requestData request.TeamAddRequest
// 		respError   string
// 		mockError   error
// 	}{
// 		{
// 			name: "Team already existst",
// 			requestData: request.TeamAddRequest{
// 				TeamName: "team 1",
// 				Members: []request.TeamMember{
// 					{
// 						UserID:   "u1",
// 						Username: "Alice",
// 						IsActive: true,
// 					},
// 					{
// 						UserID:   "u2",
// 						Username: "Bob",
// 						IsActive: true,
// 					},
// 				},
// 			},
// 		},
// 	}

// 	for _, tc := range cases {
// 		t.Run(tc.name, func(t *testing.T) {

// 			handler := team.Add()

// 			input, err := json.Marshal(tc.requestData)
// 			require.Nil(t, err)

// 			req, err := http.NewRequest(http.MethodPost, "/team/add", bytes.NewReader(input))
// 			require.NoError(t, err)

// 			rr := httptest.NewRecorder()
// 			handler.ServeHTTP(rr, req)

// 			res := rr.Result()

// 			require.Equal(t, rr.Code, http.StatusCreated)
// 			require.Equal(t, res.Header.Get("Content-Type"), "application/json")
// 		})
// 	}
// }
