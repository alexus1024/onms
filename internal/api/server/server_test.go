package server_test

import (
	"context"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/alexus1024/onms/internal/api/server"
	"github.com/alexus1024/onms/internal/storage"
	"github.com/alexus1024/onms/internal/storage/memory"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

const json1 = `{
    "machineId": 12345,
    "stats": {
        "cpuTemp": "90c",
        "fanSpeed": 400,
        "HDDSpace": 800,
		"internalTemp": 123.456
    },
    "lastLoggedIn": "admin/Paul",
    "sysTime": "2022-04-23T18:25:43.511Z"
}`
const contentTypeJson = "application/json"

func startTestServer(t *testing.T) (*httptest.Server, http.Client) {
	log := logrus.New().WithField("test", t.Name())
	memStorage := memory.NewMemoryRepo()
	handler := server.GetMux(&server.AppContext{Log: log, Repo: memStorage})

	testServer := httptest.NewServer(handler)
	testServer.Config.BaseContext = func(l net.Listener) context.Context {
		return storage.WithStorage(context.Background(), memStorage)
	}

	return testServer, *testServer.Client()
}

func TestMainScenario(t *testing.T) {
	testServer, client := startTestServer(t)
	defer testServer.Close()

	// Add first entry
	resp, err := client.Post(testServer.URL, contentTypeJson, strings.NewReader(json1))
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// add second entry
	resp, err = client.Post(testServer.URL, contentTypeJson, strings.NewReader(json1))
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// read entries

	resp, err = client.Get(testServer.URL)
	require.NoError(t, err)

	responseBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	assert.Equal(t, contentTypeJson, resp.Header.Get("content-type"))
	assert.True(t, gjson.ValidBytes(responseBody))
	assert.Equal(t, int64(2), gjson.GetBytes(responseBody, "#").Int())
	assert.Equal(t, "[12345,12345]", gjson.GetBytes(responseBody, "#.machineId").Raw)
	assert.Equal(t, "2022-04-23T18:25:43.511Z", gjson.GetBytes(responseBody, "1.sysTime").String())
	assert.Equal(t, float64(90), gjson.GetBytes(responseBody, "0.stats.cpuTemp").Num)
	assert.Equal(t, 123.456, gjson.GetBytes(responseBody, "1.stats.internalTemp").Num)

	if t.Failed() {
		assert.Fail(t, "full response", string(responseBody))
	}
}

func TestContentTypeRequired(t *testing.T) {
	testServer, client := startTestServer(t)
	defer testServer.Close()

	resp, err := client.Post(testServer.URL, "", strings.NewReader(json1))
	require.NoError(t, err)
	assert.Equal(t, http.StatusUnsupportedMediaType, resp.StatusCode)

}
