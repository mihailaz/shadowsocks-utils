package shadowsocks

import (
	"encoding/base64"
	"fmt"
	"net/url"
)

type (
	Info struct {
		EncryptionMethod string
		Password         string
		Host             string
		Params           map[string][]string
	}
)

func Parse(str string) (*Info, error) {
	if str == "" {
		return nil, fmt.Errorf("empty url")
	}
	uri, err := url.Parse(str)
	if err != nil {
		return nil, err
	}
	switch uri.Scheme {
	case "ss":
		return SsParse(uri)
	case "ssconf":
		return SsConfParse(uri)
	default:
		return nil, fmt.Errorf("invalid scheme, expect ss or ssconf, got: %s", uri.Scheme)
	}
}

func (s *Info) String() string {
	usr := fmt.Sprintf("%s:%s", s.EncryptionMethod, s.Password)
	encoded := base64.URLEncoding.EncodeToString([]byte(usr))
	return fmt.Sprintf("ss://%s@%s", encoded, s.Host)
}
