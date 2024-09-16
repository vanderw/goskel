package htp

import (
	"io"
	"net/http"
	"time"
)

var (
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36"
)

// TODO HTTP keep-alive enhancement
func Request(method, url string, secs int, headers map[string]string, data io.Reader) (int, http.Header, []byte, error) {
	c := http.Client{
		Timeout: time.Duration(secs) * time.Second,
	}
	req, err := http.NewRequest(method, url, data)
	if nil != err {
		return 0, nil, nil, err
	}
	// headers
	req.Header.Set("User-Agent", UserAgent)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.Do(req)
	if nil != err {
		return 0, nil, nil, err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if nil != err {
		return 0, nil, nil, err
	}
	return resp.StatusCode, resp.Header, bytes, nil
}

func Get(url string, secs int, headers map[string]string) (int, http.Header, []byte, error) {
	return Request("GET", url, secs, headers, nil)
}

func Post(url string, secs int, headers map[string]string, data io.Reader) (int, http.Header, []byte, error) {
	return Request("POST", url, secs, headers, data)
}
