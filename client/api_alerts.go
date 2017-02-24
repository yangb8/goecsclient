package client

import "fmt"

// Alert ...
type Alert struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	Severity     string `json:"severity"`
	UserID       string `json:"userId"`
	ServiceType  string `json:"serviceType"`
	ResourceID   string `json:"resourceId"`
	Timestamp    string `json:"timestamp"`
	Namespace    string `json:"namespace"`
	Description  string `json:"description"`
	Acknowledged bool   `json:"acknowledged"`
}

// LatestAlertsResult ...
type LatestAlertsResult struct {
	Alert        []Alert `json:"alert"`
	MaxAlerts    int64
	NextMarker   string
	Filter       string
	NextPageLink string
}

// GetLatestAlerts ...
func (e *EcsMgmtClient) GetLatestAlerts(limit int) (*LatestAlertsResult, error) {
	result := &LatestAlertsResult{}
	uri := "/vdc/alerts/latest.json"
	if limit > 0 {
		uri = fmt.Sprintf("%s?limit=%d", uri, limit)
	}
	if err := e.GetMgmtQueryResultJSON(uri, result); err != nil {
		return nil, err
	}
	return result, nil
}
