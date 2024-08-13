package shadowsocks

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestSsConfParse_WhenCalled_CallsHttpClientDo(t *testing.T) {
	mockClient := newMockHttpClient(t)
	expectedRequest, _ := http.NewRequest("GET", "https://example.com:1234/some-path", nil)
	resp := &http.Response{
		StatusCode: 200,
		Header:     map[string][]string{"Content-Type": {"text/plain"}},
		Body:       io.NopCloser(bytes.NewBufferString("ss://c29tZS1lbmNyaXB0aW9uLW1ldGhvZDpzb21lLXBhc3N3b3JkCg@example.com:8080")),
	}
	mockClient.On("Do", expectedRequest).Return(resp, nil)
	HttpClient = mockClient
	uri, err := url.Parse("ssconf://example.com:1234/some-path")
	assert.Nil(t, err)
	_, err = SsConfParse(uri)
	assert.Nil(t, err)
}
