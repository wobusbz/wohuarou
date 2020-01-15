package util

import "net/http"

func HTTPGet(url string) (*http.Response, error) {
	resp, err := http.NewRequest("GET", url, nil)
	resp.Header = RandUserAgent()
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(resp)
}
