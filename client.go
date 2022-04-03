package NovelUpdatesClient

import (
	"crypto/tls"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

var client = &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{MaxVersion: tls.VersionTLS12}}}

func createDocumentFromURL(url string) (*goquery.Document, error) {

	response, err := doGetRequest(url)
	if err != nil {
		return nil, err
	}

	// NewDocumentFromReader does not
	// close body itself.
	body := response.Body
	defer body.Close()

	// Create GoQuery Document
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func doGetRequest(url string) (*http.Response, error) {

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

	return response, nil
}
