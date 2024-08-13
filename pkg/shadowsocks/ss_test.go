package shadowsocks

import (
	"net/url"
	"reflect"
	"testing"
)

func TestSsParse(t *testing.T) {
	type args struct {
		uri *url.URL
	}
	tests := []struct {
		name    string
		args    args
		want    *Info
		wantErr bool
	}{
		{"success", args{
			uri: &url.URL{Scheme: "ss", User: url.User("c29tZS1lbmNyaXB0aW9uLW1ldGhvZDpzb21lLXBhc3N3b3JkCg"),
				Host: "example.com:8080", Path: "/some-path", RawQuery: "param1=value1&param2=value2"}},
			&Info{EncryptionMethod: "some-encription-method", Password: "some-password",
				Host: "example.com:8080", Params: map[string][]string{"param1": {"value1"}, "param2": {"value2"}}}, false},
		{
			"invalid scheme", args{uri: &url.URL{Scheme: "http", User: url.User("c29tZS1lbmNyaXB0aW9uLW1ldGhvZDpzb21lLXBhc3N3b3JkCg=="),
				Host: "example.com:8080", Path: "/some-path", RawQuery: "param1=value1&param2=value2"}},
			nil, true,
		},
		{
			"empty username", args{uri: &url.URL{Scheme: "ss", User: url.User(""),
				Host: "example.com:8080", Path: "/some-path", RawQuery: "param1=value1&param2=value2"}},
			nil, true,
		},
		{
			"invalid username", args{uri: &url.URL{Scheme: "ss", User: url.User("aW52YWxpZAo="),
				Host: "example.com:8080", Path: "/some-path", RawQuery: "param1=value1&param2=value2"}},
			nil, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SsParse(tt.args.uri)
			if (err != nil) != tt.wantErr {
				t.Errorf("SsParse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SsParse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
