package healtcheck_api_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const baseUrl = "http://localhost:8080/"

func Test(t *testing.T) {
	t.Run("health check", func(t *testing.T) {
		response, exc := http.Get(baseUrl)
		defer response.Body.Close()
		assert.Empty(t, exc)
		assert.Equal(t, response.StatusCode, 200)
	})
}
