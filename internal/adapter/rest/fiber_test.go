package rest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestFiber_Startup(t *testing.T) {
	// Create a new instance of the Fiber struct
	f := &Fiber{}

	// Call the Startup method
	err := f.Startup()

	// Assert that there is no error
	assert.NoError(t, err)

	// Create a new HTTP request
	req := &fasthttp.RequestCtx{}

	// Create a new HTTP response
	res := &fasthttp.Response{}

	// Call the Fiber app's handler function
	f.App.Handler()(req)

	// Assert that the response status code is 200 OK
	assert.Equal(t, fasthttp.StatusOK, res.StatusCode())

}

func TestFiber_Shutdown(t *testing.T) {
	// Create a new instance of the Fiber struct
	f := &Fiber{}

	// Call the Startup method
	err := f.Startup()

	// Call the Shutdown method
	err = f.Shutdown()

	// Assert that there is no error
	assert.NoError(t, err)
}
