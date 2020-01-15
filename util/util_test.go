package util

import (
	"net/http"
	"testing"
)

var httpUnitTest = []struct {
	title string
	f     func(string) (*http.Response, error)
	out   string
}{
	{"TestHTTPGet", HTTPGet, "200"},
}

func TestHTTPGet(t *testing.T) {
	var url = "http://www.baidu.com"
	resp, err := HTTPGet(url)
	if resp.Status != "200" && err != nil {
		t.Errorf("reality: %s  expect:%d", resp.Status, 200)
	}
	for _, v := range httpUnitTest {
		if resp, err := v.f(url); resp.Status != v.out && err != nil {
			t.Errorf("reality: %s expect: %s", resp.Status, v.out)
		}
	}
}

func TestRandUserAgent(t *testing.T) {
	t.Log(RandUserAgent())
}
