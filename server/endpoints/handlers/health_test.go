package handlers

import (
	"testing"

	"github/vriaan/footballmanagerapi/tests"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	urlPing := "/ping"
	responseStatus, responseBody, err := tests.TestHTTPHandler("GET", urlPing, tests.TestParams{}, Ping)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, responseStatus)
	assert.Equal(t, "{\"message\":\"pong\"}", string(responseBody))
}

func BenchmarkPing(b *testing.B) {
	b.ReportAllocs()
	urlPing := "/ping"

	for n := 0; n < b.N; n++ {
		_, _, err := tests.TestHTTPHandler("GET", urlPing, tests.TestParams{}, Ping)
		if err != nil {
			b.Fatal(err)
		}
	}
}
