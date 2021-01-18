package httpz_test

import (
	"context"
	"encoding/json"
	"github.com/absurdlab/x/v2/httpz"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	req, err := httpz.Get(context.Background(), "http://httpbin.org/get")
	assert.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestPost(t *testing.T) {
	req, err := httpz.Post(context.Background(), "http://httpbin.org/post")
	assert.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestPut(t *testing.T) {
	req, err := httpz.Put(context.Background(), "http://httpbin.org/put")
	assert.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDelete(t *testing.T) {
	req, err := httpz.Delete(context.Background(), "http://httpbin.org/delete")
	assert.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestPostJSON(t *testing.T) {
	type obj struct {
		Foo string `json:"foo"`
	}

	req, err := httpz.Post(context.Background(), "http://httpbin.org/post", httpz.JSON(obj{Foo: "Bar"}))
	assert.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var payload = struct {
		JSON obj `json:"json"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&payload)
	assert.NoError(t, err)
	assert.Equal(t, "Bar", payload.JSON.Foo)
}

func TestBasicAuth(t *testing.T) {
	req, err := httpz.Get(context.Background(), "http://httpbin.org/basic-auth/foo/bar",
		httpz.BasicAuth("foo", "bar"))
	assert.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestBearerAuth(t *testing.T) {
	req, err := httpz.Get(context.Background(), "http://httpbin.org/bearer",
		httpz.BearerAuth("dummy"))
	assert.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
