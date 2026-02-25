package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/sigsignv/bfconv"
)

type Request = events.LambdaFunctionURLRequest
type Response = events.LambdaFunctionURLResponse

type Handler struct {
	client *http.Client
	conv   *bfconv.Converter
}

func NewHandler() *Handler {
	return &Handler{
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
		conv: bfconv.NewConverter(),
	}
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	method := req.RequestContext.HTTP.Method
	if method != http.MethodGet && method != http.MethodHead {
		return h.errorResponse(http.StatusMethodNotAllowed, "Method Not Allowed"), nil
	}

	path := strings.TrimSpace(req.RawPath)
	if path == "" {
		return h.errorResponse(http.StatusBadRequest, "Path is required"), nil
	}
	if path == "/" {
		return h.descriptionResponse(), nil
	}

	url, err := h.buildURL(path, req.RawQueryString)
	if err != nil {
		return h.errorResponse(http.StatusInternalServerError, "failed to build URL"), nil
	}

	feed, err := h.request(ctx, url)
	if err != nil {
		return h.errorResponse(
			http.StatusInternalServerError,
			fmt.Sprintf("failed to fetch feed: %v", err)), nil
	}

	payload, err := json.Marshal(feed)
	if err != nil {
		return h.errorResponse(http.StatusInternalServerError, "failed to encode feed"), nil
	}

	return h.jsonResponse(http.StatusOK, string(payload)), nil

}

func (h *Handler) buildURL(path string, query string) (string, error) {
	url, err := url.Parse("https://b.hatena.ne.jp/")
	if err != nil {
		return "", err
	}

	url.Path = path
	url.RawQuery = query

	return url.String(), nil
}

func (h *Handler) request(ctx context.Context, url string) (*bfconv.Feed, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "application/xml" {
		return nil, fmt.Errorf("unexpected content type: %s", resp.Header.Get("Content-Type"))
	}

	feed, err := h.conv.Convert(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to convert feed: %w", err)
	}

	return feed, nil
}

func (h *Handler) descriptionResponse() Response {
	payload, err := json.Marshal(map[string]string{
		"name":        "bfconv-lambda",
		"description": "bfconv-lambda is a converter that transforms Hatena Bookmark RSS feeds into JSON Feed format",
		"author":      "Sigsign <sig@signote.cc>",
		"license":     "Apache-2.0",
		"example":     "https://bfconv.signote.cc/entrylist.json",
	})
	if err != nil {
		return h.errorResponse(http.StatusInternalServerError, "failed to description encode")
	}

	return h.jsonResponse(http.StatusOK, string(payload))
}

func (h *Handler) errorResponse(status int, message string) Response {
	payload, err := json.Marshal(map[string]string{"error": message})
	if err != nil {
		return h.jsonResponse(
			http.StatusInternalServerError,
			`{"error": "Internal Server Error"}`,
		)
	}

	return h.jsonResponse(status, string(payload))
}

func (h *Handler) jsonResponse(status int, payload string) Response {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	return Response{
		StatusCode: status,
		Headers:    headers,
		Body:       payload,
	}
}
