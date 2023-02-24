package utils

import (
	"context"
	"io"
	"net/http"
	"time"
)

func NewResponse(url, ip string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url+ip, nil)
	if err != nil {
		return nil, err
	}
	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	b, err := io.ReadAll(resp.Body)
	return b, err
}
