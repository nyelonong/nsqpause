package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

func (c *Config) nsqAction(path string, action interface{}) error {
	u, err := url.Parse(c.NSQAdmin)
	if err != nil {
		return err
	}

	u.Path += path
	link := u.String()

	body, err := json.Marshal(action)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, link, bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Close = true

	hc := &http.Client{
		Timeout: c.Timeout,
	}

	resp, err := hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var msg NSQResponse
	if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(msg.Message)
	}

	return nil
}

func (c *Config) nsqInfo(path string) (*NSQChannel, error) {
	u, err := url.Parse(c.NSQAdmin)
	if err != nil {
		return nil, err
	}

	u.Path += path
	link := u.String()

	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		return nil, err
	}
	req.Close = true

	hc := &http.Client{
		Timeout: c.Timeout,
	}

	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var msg NSQResponse
		if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
			return nil, err
		}
		return nil, errors.New(msg.Message)
	}

	var msg NSQChannel
	if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
		return nil, err
	}

	return &msg, nil
}
