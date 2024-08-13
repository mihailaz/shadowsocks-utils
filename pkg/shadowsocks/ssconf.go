package shadowsocks

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
)

const (
	contentType = "Content-Type"
	textPlain   = "text/plain"
)

var (
	HttpClient httpClient = http.DefaultClient
)

//go:generate mockery --name httpClient --inpackage
type (
	httpClient interface {
		Do(req *http.Request) (*http.Response, error)
	}
)

func SsConfParse(uri *url.URL) (*Info, error) {
	if uri.Scheme != "ssconf" {
		return nil, fmt.Errorf("invalid scheme, expect ssconf, got: %s", uri.Scheme)
	}
	slog.Debug("parsing", slog.String("url", uri.String()))
	uri.Scheme = "https"
	reqUrl := uri.String()
	slog.Debug("requesting", slog.String("url", reqUrl))
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, err
	}
	resp, err := HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	contType := resp.Header.Get(contentType)
	if contType != "" && !strings.Contains(strings.ToLower(contType), textPlain) {
		return nil, fmt.Errorf("invalid content type: %s", contType)
	}
	buf := new(strings.Builder)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, err
	}
	body := buf.String()
	if body == "" {
		return nil, fmt.Errorf("empty body")
	}
	ssUrl, err := url.Parse(body)
	if err != nil {
		return nil, err
	}
	return SsParse(ssUrl)
}
