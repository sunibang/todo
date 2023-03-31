package main

import (
	"net/http/httptest"

	"github.com/golang/mock/gomock"
	"testing"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func TestAddTodo(t *testing.T) {

}

func TestGetTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	server := newTestServer(t, store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/accounts/%d", tc.accountID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)
	server.router.ServeHTTP(recorder, request)
	checkResponse(t, recorder)
}

func checkResponse(t *testing.T, recorder *httptest.ResponseRecorder) {
	require.Equal(t, http.StatusBadRequest, recorder.Code)
}
