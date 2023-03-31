package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

	"testing"

	"github.com/golang/mock/gomock"
)

func TestAddTodo(t *testing.T) {

}

func TestGetTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	server := newTestServer(t)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/todo/%d", 1)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)
	server.router.ServeHTTP(recorder, request)
	checkResponse(t, recorder)
}

func checkResponse(t *testing.T, recorder *httptest.ResponseRecorder) {
	require.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}

func newTestServer(t *testing.T) *Server {
	server, err := NewServer()
	require.NoError(t, err)

	return server
}
