package actions

import (
	"testing"

	"github/vriaan/footballmanagerapi/test"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	urlPing := "/ping"
	responseStatus, responseBody, err := test.CallAction("GET", urlPing, test.Params{}, Ping, "")
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
		_, _, err := test.CallAction("GET", urlPing, test.Params{}, Ping, "")
		if err != nil {
			b.Fatal(err)
		}
	}
}
