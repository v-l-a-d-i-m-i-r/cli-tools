// Package dockerhubimagetags fetches image tags from the Docker Hub API.
package dockerhubimagetags

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

// Result represents a single tag entry returned by the Docker Hub API.
type Result struct {
	Name string
}

// Page represents a paginated response from the Docker Hub API.
type Page struct {
	Count   int
	Results []Result
}

func fetchPage(namespace, repository string, page, pageSize int) (Page, error) {
	base, err := url.Parse("https://hub.docker.com")
	if err != nil {
		return Page{}, fmt.Errorf("parse base URL: %w", err)
	}

	base.Path += "/v2/namespaces/" + namespace + "/repositories/" + repository + "/tags"
	params := url.Values{}
	params.Add("page", strconv.Itoa(page))
	params.Add("page_size", strconv.Itoa(pageSize))
	base.RawQuery = params.Encode()

	res, err := http.Get(base.String()) //nolint:noctx
	if err != nil {
		return Page{}, fmt.Errorf("fetch page %d: %w", page, err)
	}
	defer func() { _ = res.Body.Close() }()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Page{}, fmt.Errorf("read response body: %w", err)
	}

	var p Page
	if err := json.Unmarshal(body, &p); err != nil {
		return Page{}, fmt.Errorf("unmarshal response: %w", err)
	}

	return p, nil
}

func tagsFromPage(page Page) []string {
	tags := make([]string, len(page.Results))
	for i, result := range page.Results {
		tags[i] = result.Name
	}
	return tags
}

// FetchTags returns all image tags for the given namespace and repository.
func FetchTags(namespace, repository string) ([]string, error) {
	const pageSize = 100

	firstPage, err := fetchPage(namespace, repository, 1, pageSize)
	if err != nil {
		return nil, err
	}

	tags := tagsFromPage(firstPage)
	pagesCount := int(math.Ceil(float64(firstPage.Count) / float64(pageSize)))

	if pagesCount <= 1 {
		return tags, nil
	}

	var (
		wg   sync.WaitGroup
		mu   sync.Mutex
		errs []error
	)

	wg.Add(pagesCount - 1)

	for i := 2; i <= pagesCount; i++ {
		go func(pageNum int) {
			defer wg.Done()

			page, err := fetchPage(namespace, repository, pageNum, pageSize)

			mu.Lock()
			defer mu.Unlock()

			if err != nil {
				errs = append(errs, err)
				return
			}

			tags = append(tags, tagsFromPage(page)...)
		}(i)
	}

	wg.Wait()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	return tags, nil
}
