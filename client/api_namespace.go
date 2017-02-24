package client

import (
	"fmt"
	"net/url"
	"strings"
)

// NamespacesResult ...
type NamespacesResult struct {
	Namespace []struct {
		Name string `json:"name"`
		ID   string `json:"id"`
		Link struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"link"`
	} `json:"namespace"`
	Filter       string `json:"Filter"`
	NextMarker   string `json:"NextMarker"`
	NextPageLink string `json:"NextPageLink"`
}

// GetNamespaces ...
func (e *EcsMgmtClient) GetNamespaces() (*NamespacesResult, error) {
	return e.GetNamespacesWithOptions(0, "", "")
}

// GetNamespacesWithOptions ...
func (e *EcsMgmtClient) GetNamespacesWithOptions(limit int, marker, name string) (*NamespacesResult, error) {
	result := &NamespacesResult{}
	var (
		queries []string
		uri     string
	)
	if limit > 0 {
		queries = append(queries, fmt.Sprintf("limit=%d", limit))
	}
	if len(marker) > 0 {
		queries = append(queries, fmt.Sprintf("marker=%s", url.QueryEscape(marker)))
	}
	if len(name) > 0 {
		queries = append(queries, fmt.Sprintf("name=%s", url.QueryEscape(name)))
	}
	queryStr := strings.Join(queries, "&")
	if len(queryStr) == 0 {
		uri = "/object/namespaces.json"
	} else {
		uri = fmt.Sprintf("%s?%s", "/object/namespaces.json", queryStr)
	}

	if err := e.GetMgmtQueryResultJSON(uri, result); err != nil {
		return nil, err
	}
	return result, nil
}
