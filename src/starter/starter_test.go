package starter_test

import (
	"io"
	"net/http/httptest"
	"testing"

	starter "github.com/ragazziluis/tdd-golang/starter"
	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	greeting := starter.SayHello("William")
	assert.Equal(t, "Hello William. Welcome!", greeting)

	anotherGreeting := starter.SayHello("asdf ghjkl")
	assert.Equal(t, "Hello asdf ghjkl. Welcome!", anotherGreeting)
}

func TestOddOrEven(t *testing.T) {
	t.Run("Check Non Negative Numbers", func(t *testing.T) {
		assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
		assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
		assert.Equal(t, "0 is an even number", starter.OddOrEven(0))
	})
	t.Run("Check Negative Numbers", func(t *testing.T) {
		assert.Equal(t, "-45 is an odd number", starter.OddOrEven(-45))
		assert.Equal(t, "-42 is an even number", starter.OddOrEven(-42))
	})
}

func TestCheckHealth(t *testing.T) {
	t.Run("Check health status", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://mysite.com/example", nil)
		writer := httptest.NewRecorder()
		starter.CheckHealth(writer, req)
		response := writer.Result()
		body, err := io.ReadAll(response.Body)

		assert.Equal(t, "health check passed", string(body))
		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t,
			"text/plain; charset=utf-8",
			response.Header.Get("Content-Type"))
		assert.Equal(t, nil, err)
	})
}
