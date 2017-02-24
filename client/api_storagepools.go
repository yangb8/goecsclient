package client

// LocalzoneStoragepoolsResult ...
type LocalzoneStoragepoolsResult struct {
	Title string `json:"title"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Embedded struct {
		Instances []struct {
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Nodes struct {
					Href string `json:"href"`
				} `json:"nodes"`
			} `json:"_links"`
			NumBadDisks                     string `json:"numBadDisks"`
			ChunksL1JournalTotalSize        string `json:"chunksL1JournalTotalSize"`
			ChunksL0BtreeTotalSize          string `json:"chunksL0BtreeTotalSize"`
			ChunksRepoTotalSealSize         string `json:"chunksRepoTotalSealSize"`
			ChunksL1BtreeNumber             string `json:"chunksL1BtreeNumber"`
			ChunksXorNumber                 string `json:"chunksXorNumber"`
			ChunksGeoCopyNumber             string `json:"chunksGeoCopyNumber"`
			ChunksL1JournalAvgSize          string `json:"chunksL1JournalAvgSize"`
			ChunksL0JournalNumber           string `json:"chunksL0JournalNumber"`
			NumNodesWithSufficientDiskSpace string `json:"numNodesWithSufficientDiskSpace"`
			ChunksL0BtreeNumber             string `json:"chunksL0BtreeNumber"`
			ID                              string `json:"id"`
			ChunksL1JournalNumber           string `json:"chunksL1JournalNumber"`
			ChunksGeoCacheTotalSize         string `json:"chunksGeoCacheTotalSize"`
			ChunksEcCompleteTimeEstimate    string `json:"chunksEcCompleteTimeEstimate"`
			ChunksRepoNumber                string `json:"chunksRepoNumber"`
			ChunksL0JournalTotalSize        string `json:"chunksL0JournalTotalSize"`
			RecoveryCompleteTimeEstimate    string `json:"recoveryCompleteTimeEstimate"`
			APIChange                       string `json:"apiChange"`
			NumNodes                        string `json:"numNodes"`
			ChunksL1BtreeTotalSize          string `json:"chunksL1BtreeTotalSize"`
			ChunksL0JournalAvgSize          string `json:"chunksL0JournalAvgSize"`
			NumBadNodes                     string `json:"numBadNodes"`
			ChunksXorTotalSize              string `json:"chunksXorTotalSize"`
			NumGoodDisks                    string `json:"numGoodDisks"`
			ChunksL1BtreeAvgSize            string `json:"chunksL1BtreeAvgSize"`
			Name                            string `json:"name"`
			ChunksL0BtreeAvgSize            string `json:"chunksL0BtreeAvgSize"`
			NumGoodNodes                    string `json:"numGoodNodes"`
			NumDisks                        string `json:"numDisks"`
			ChunksGeoCopyTotalSize          string `json:"chunksGeoCopyTotalSize"`
			ChunksGeoCacheCount             string `json:"chunksGeoCacheCount"`
			ChunksRepoAvgSealSize           string `json:"chunksRepoAvgSealSize"`
			DiskSpaceOfflineTotalCurrent    []struct {
				T     string `json:"t"`
				Space string `json:"Space"`
			} `json:"diskSpaceOfflineTotalCurrent"`
			DiskSpaceAllocatedPercentage []struct {
				T       string `json:"t"`
				Percent string `json:"Percent"`
			} `json:"diskSpaceAllocatedPercentage"`
			RecoveryRateCurrent []struct {
				T    string `json:"t"`
				Rate string `json:"Rate"`
			} `json:"recoveryRateCurrent"`
			DiskSpaceFreeCurrent []struct {
				T     string `json:"t"`
				Space string `json:"Space"`
			} `json:"diskSpaceFreeCurrent"`
			DiskSpaceAllocatedPercentageCurrent []struct {
				T       string `json:"t"`
				Percent string `json:"Percent"`
			} `json:"diskSpaceAllocatedPercentageCurrent"`
			ChunksEcCodedTotalSealSize []struct {
				T     string `json:"t"`
				Space string `json:"Space"`
			} `json:"chunksEcCodedTotalSealSize"`
			DiskSpaceTotalCurrent []struct {
				T     string `json:"t"`
				Space string `json:"Space"`
			} `json:"diskSpaceTotalCurrent"`
			ChunksEcRateCurrent []struct {
				T    string `json:"t"`
				Rate string `json:"Rate"`
			} `json:"chunksEcRateCurrent"`
			ChunksEcCodedRatioCurrent []struct {
				T       string `json:"t"`
				Percent string `json:"Percent"`
			} `json:"chunksEcCodedRatioCurrent"`
			ChunksEcApplicableTotalSealSize []struct {
				T     string `json:"t"`
				Space string `json:"Space"`
			} `json:"chunksEcApplicableTotalSealSize"`
			ChunksEcRate []struct {
				T    string `json:"t"`
				Rate string `json:"Rate"`
			} `json:"chunksEcRate"`
			DiskSpaceAllocatedCurrent []struct {
				T     string `json:"t"`
				Space string `json:"Space"`
			} `json:"diskSpaceAllocatedCurrent"`
			RecoveryRate []struct {
				T    string `json:"t"`
				Rate string `json:"Rate"`
			} `json:"recoveryRate"`
			ChunksEcApplicableTotalSealSizeCurrent []struct {
				T     string `json:"t"`
				Space string `json:"Space"`
			} `json:"chunksEcApplicableTotalSealSizeCurrent"`
			DiskSpaceTotal []struct {
				T     string `json:"t"`
				Space string `json:"Space"`
			} `json:"diskSpaceTotal"`
			RecoveryBadChunksTotalSize []struct {
				T     string `json:"t"`
				Space string `json:"Space"`
			} `json:"recoveryBadChunksTotalSize"`
			RecoveryBadChunksTotalSizeCurrent []struct {
				T     string `json:"t"`
				Space string `json:"Space"`
			} `json:"recoveryBadChunksTotalSizeCurrent"`
			DiskSpaceAllocated []struct {
				T     string `json:"t"`
				Space string `json:"Space"`
			} `json:"diskSpaceAllocated"`
			ChunksEcCodedTotalSealSizeCurrent []struct {
				T     string `json:"t"`
				Space string `json:"Space"`
			} `json:"chunksEcCodedTotalSealSizeCurrent"`
			ChunksEcCodedRatio []struct {
				T       string `json:"t"`
				Percent string `json:"Percent"`
			} `json:"chunksEcCodedRatio"`
			DiskSpaceFree []struct {
				T     string `json:"t"`
				Space string `json:"Space"`
			} `json:"diskSpaceFree"`
			DiskSpaceTotalSummary struct {
				Min []struct {
					T     string `json:"t"`
					Space string `json:"Space"`
				} `json:"Min"`
				Max []struct {
					T     string `json:"t"`
					Space string `json:"Space"`
				} `json:"Max"`
				Avg string `json:"Avg"`
			} `json:"diskSpaceTotalSummary"`
			ChunksEcCodedRatioSummary struct {
				Min []struct {
					T       string `json:"t"`
					Percent string `json:"Percent"`
				} `json:"Min"`
				Max []struct {
					T       string `json:"t"`
					Percent string `json:"Percent"`
				} `json:"Max"`
				Avg string `json:"Avg"`
			} `json:"chunksEcCodedRatioSummary"`
			ChunksEcCodedTotalSealSizeSummary struct {
				Min []struct {
					T     string `json:"t"`
					Space string `json:"Space"`
				} `json:"Min"`
				Max []struct {
					T     string `json:"t"`
					Space string `json:"Space"`
				} `json:"Max"`
				Avg string `json:"Avg"`
			} `json:"chunksEcCodedTotalSealSizeSummary"`
			DiskSpaceFreeSummary struct {
				Min []struct {
					T     string `json:"t"`
					Space string `json:"Space"`
				} `json:"Min"`
				Max []struct {
					T     string `json:"t"`
					Space string `json:"Space"`
				} `json:"Max"`
				Avg string `json:"Avg"`
			} `json:"diskSpaceFreeSummary"`
			DiskSpaceAllocatedPercentageSummary struct {
				Min []struct {
					T       string `json:"t"`
					Percent string `json:"Percent"`
				} `json:"Min"`
				Max []struct {
					T       string `json:"t"`
					Percent string `json:"Percent"`
				} `json:"Max"`
				Avg string `json:"Avg"`
			} `json:"diskSpaceAllocatedPercentageSummary"`
			RecoveryBadChunksTotalSizeSummary struct {
				Min []struct {
					T     string `json:"t"`
					Space string `json:"Space"`
				} `json:"Min"`
				Max []struct {
					T     string `json:"t"`
					Space string `json:"Space"`
				} `json:"Max"`
				Avg string `json:"Avg"`
			} `json:"recoveryBadChunksTotalSizeSummary"`
			DiskSpaceAllocatedSummary struct {
				Min []struct {
					T     string `json:"t"`
					Space string `json:"Space"`
				} `json:"Min"`
				Max []struct {
					T     string `json:"t"`
					Space string `json:"Space"`
				} `json:"Max"`
				Avg string `json:"Avg"`
			} `json:"diskSpaceAllocatedSummary"`
			ChunksEcRateSummary struct {
				Min []struct {
					T    string `json:"t"`
					Rate string `json:"Rate"`
				} `json:"Min"`
				Max []struct {
					T    string `json:"t"`
					Rate string `json:"Rate"`
				} `json:"Max"`
				Avg string `json:"Avg"`
			} `json:"chunksEcRateSummary"`
			ChunksEcApplicableTotalSealSizeSummary struct {
				Min []struct {
					T     string `json:"t"`
					Space string `json:"Space"`
				} `json:"Min"`
				Max []struct {
					T     string `json:"t"`
					Space string `json:"Space"`
				} `json:"Max"`
				Avg string `json:"Avg"`
			} `json:"chunksEcApplicableTotalSealSizeSummary"`
			RecoveryRateSummary struct {
				Min []struct {
					T    string `json:"t"`
					Rate string `json:"Rate"`
				} `json:"Min"`
				Max []struct {
					T    string `json:"t"`
					Rate string `json:"Rate"`
				} `json:"Max"`
				Avg string `json:"Avg"`
			} `json:"recoveryRateSummary"`
		} `json:"_instances"`
	} `json:"_embedded"`
}

// GetLocalzoneStoragepools ...
func (e *EcsMgmtClient) GetLocalzoneStoragepools() (*LocalzoneStoragepoolsResult, error) {
	result := &LocalzoneStoragepoolsResult{}
	if err := e.GetMgmtQueryResultJSON("/dashboard/zones/localzone/storagepools", result); err != nil {
		return nil, err
	}
	return result, nil
}
