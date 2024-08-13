package shadowsocks

import (
	"encoding/base64"
	"errors"
	"log/slog"
	"net/url"
	"strings"
)

func SsParse(uri *url.URL) (*Info, error) {
	if uri.Scheme != "ss" {
		return nil, errors.New("only support ss://")
	}
	slog.Debug("parsing", slog.String("url", uri.String()))
	data := uri.User.Username()
	if data == "" {
		return nil, errors.New("username is empty")
	}
	decoded, err := base64.RawURLEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	decodedStr := strings.Trim(string(decoded), "\n")
	sepIdx := strings.Index(decodedStr, ":")
	if sepIdx < 0 {
		return nil, errors.New("invalid username: " + decodedStr)
	}
	return &Info{
		EncryptionMethod: decodedStr[:sepIdx],
		Password:         decodedStr[sepIdx+1:],
		Host:             uri.Host,
		Params:           uri.Query(),
	}, nil
}
