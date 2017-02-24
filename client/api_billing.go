package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// NamespaceBillingInfosResult ...
type NamespaceBillingInfosResult struct {
	NamespaceBillingInfos []struct {
		Namespace         string `json:"namespace"`
		NextMarker        string `json:"next_marker"`
		SampleTime        string `json:"sample_time"`
		TotalObjects      int64  `json:"total_objects"`
		TotalSize         int64  `json:"total_size"`
		TotalSizeUnit     string `json:"total_size_unit"`
		BucketBillingInfo []struct {
			Name          string `json:"name"`
			VpoolID       string `json:"vpool_id"`
			TotalSize     int64  `json:"total_size"`
			TotalSizeUnit string `json:"total_size_unit"`
			TotalObjects  int64  `json:"total_objects"`
			SampleTime    string `json:"sample_time"`
			TagSet        []struct {
				Key   string `json:"Key"`
				Value string `json:"Value"`
			} `json:"TagSet"`
		} `json:"bucket_billing_info"`
	} `json:"namespace_billing_infos"`
}

// GetNamespaceBillingInfos ...
func (e *EcsMgmtClient) GetNamespaceBillingInfos(ids []string) (*NamespaceBillingInfosResult, error) {
	return e.GetNamespaceBillingInfosWithOptions("", "", ids)
}

// GetNamespaceBillingInfosWithOptions ...
func (e *EcsMgmtClient) GetNamespaceBillingInfosWithOptions(includeBucketDetail, sizeunit string, ids []string) (*NamespaceBillingInfosResult, error) {
	result := &NamespaceBillingInfosResult{}
	var (
		queries []string
		uri     string
	)
	detail := strings.ToLower(includeBucketDetail)
	if detail == "true" || detail == "false" {
		queries = append(queries, fmt.Sprintf("include_bucket_detail=%s", detail))
	}
	UNIT := strings.ToUpper(sizeunit)
	if UNIT == "KB" || UNIT == "MB" || UNIT == "GB" {
		queries = append(queries, fmt.Sprintf("sizeunit=%s", UNIT))
	}
	queryStr := strings.Join(queries, "&")
	if len(queryStr) == 0 {
		uri = "/object/billing/namespace/info.json"
	} else {
		uri = fmt.Sprintf("%s?%s", "/object/billing/namespace/info.json", queryStr)
	}

	nsList := struct {
		ID []string `json:"id"`
	}{}
	for _, v := range ids {
		nsList.ID = append(nsList.ID, v)
	}

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(nsList)
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	if err := e.GetMgmtQueryResultJSONByPost(uri, body, 0, headers, result); err != nil {
		return nil, err
	}
	return result, nil
}
