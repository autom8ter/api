package api

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func (h HTTPMethod) Normalize() string {
	if h == HTTPMethod_POST {
		return "POST"
	} else {
		return "GET"
	}
}

func (h *HTTPRequest) Do() (*Bytes, error) {
	u, err := url.Parse(h.Url)
	if err != nil {
		return nil, err
	}
	r := &http.Request{
		Method: h.Method.Normalize(),
		URL:    u,
	}
	if h.Token != "" {
		r.Header.Set("Authorization", "Bearer "+h.Token)
	}
	if len(h.Headers) != 0 {
	}
	for k, v := range h.Headers {
		r.Header.Set(k, v)
	}

	if len(h.Headers) != 0 {
		for k, v := range h.Form {
			r.Form.Set(k, v)
		}
	}

	if len(h.Cookies) != 0 {
		for k, v := range h.Cookies {
			r.AddCookie(&http.Cookie{
				Name:  k,
				Value: v,
			})
		}
	}

	resp, err := httpClient.Do(r)
	if err != nil {
		return nil, Util.WrapErr(err, resp.Status)
	}
	bits, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, Util.WrapErr(err, "reading http response body")
	}
	return &Bytes{
		Bits: bits,
	}, nil

}
