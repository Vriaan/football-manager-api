package handlers_test

import (
	"testing"

	"github/vriaan/footballmanagerapi/tests"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	responseStatus, responseBody, err := tests.DoJSONRequest("GET", "/ping", nil)
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
		_, _, err := tests.DoJSONRequest("GET", urlPing, nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}
