package NovelUpdatesClient

import (
	"crypto/tls"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strings"
)

var client = &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{MaxVersion: tls.VersionTLS12}}}

type BasicRequester struct{}

func (s *BasicRequester) Request(url string) (io.ReadCloser, error) {

	// Create GET Request
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "NovelUpdatesClient")

	// Request NovelUpdates Page
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("err: %s", response.Status)
	}

	return response.Body, nil
}

func createDocumentFromRequest(req Requester, url string) (*goquery.Document, error) {

	response, err := req.Request(url)
	if err != nil {
		return nil, err
	}

	// NewDocumentFromReader does not
	// close body itself.
	defer response.Close()

	// Create GoQuery Document
	doc, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

type closerWrapper struct{ io.Reader }

func (s closerWrapper) Close() error { return nil }

func DummyRequester(content string) RequesterFunc {
	return func(url string) (io.ReadCloser, error) {
		reader := strings.NewReader(content)
		readCloser := closerWrapper{reader}
		return readCloser, nil
	}
}
