package request

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	"parallel.downloader.com/entity"
)

type HttpMethod string

const (
	HTTP_METHOD_GET  HttpMethod = "GET"
	HTTP_METHOD_POST HttpMethod = "POST"
)

const (
	HEADER_RANGE = "Range"
)

func HttpRequest(ctx *context.Context, d *entity.Downloader, method HttpMethod, rangeStr string) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Second * time.Duration(d.Timeout),
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	if d.Option.ProxyHost != "" {
		proxyStr, err := url.Parse(d.Option.ProxyHost)
		if err != nil {
			return nil, err
		}
		transport.Proxy = http.ProxyURL(proxyStr)
	}
	if d.ProxyEnv == "on" || d.ProxyEnv == "ON" {
		transport.Proxy = http.ProxyFromEnvironment
	}
	client.Transport = transport

	request, err := http.NewRequest(string(method), d.Target, nil)
	if err != nil {
		return nil, err
	}

	if len(d.Headers) > 0 {
		for k, v := range d.Headers {
			request.Header.Set(k, v)
		}
	}

	if rangeStr != "" {
		request.Header.Set(HEADER_RANGE, rangeStr)
	}

	request = request.WithContext(*ctx)
	return nil, nil
}
