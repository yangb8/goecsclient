package client

// CapacityResult ...
type CapacityResult struct {
	TotalProvisionedGB int64 `json:"totalProvisioned_gb"`
	TotalFreeGB        int64 `json:"totalFree_gb"`
}

// GetCapacity ...
func (e *EcsMgmtClient) GetCapacity() (*CapacityResult, error) {
	result := &CapacityResult{}
	if err := e.GetMgmtQueryResultJSON("/object/capacity.json", result); err != nil {
		return nil, err
	}
	return result, nil
}
