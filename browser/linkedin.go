package browser

import (
	"fmt"
	"github.com/pkg/browser"
	"net/url"
	"strings"
)

// OpenSearchInLinked opens a browser window with the search results for the given keywords on LinkedIn
func OpenSearchInLinked(keywords []string) error {
	searchTerm := strings.Join(keywords, " ")
	safeSearchTerm := url.PathEscape(searchTerm)
	urlPath := fmt.Sprintf("https://www.linkedin.com/search/results/all/?keywords=%s&origin=GLOBAL_SEARCH_HEADER&sid=0Wc", safeSearchTerm)
	err := browser.OpenURL(urlPath)
	if err != nil {
		return err
	}
	return nil
}
