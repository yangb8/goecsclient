package tools

import "strings"

// ParseHosts parses hosts string: ecsbeat is special that we only query one of hosts each time but at all each one as default
func ParseHosts(hosts []string) []string {
	if len(hosts) == 0 {
		return []string{}
	}
	// Currently, only support 1 hosts string
	// e.g. NOTE the position of double quotes
	// hosts: ["https://10.1.83.51:4443, https://10.1.83.52:4443, https://10.1.83.53:4443, https://10.1.83.54:4443"]
	result := make([]string, 0)
	for _, s := range strings.Split(hosts[0], ",") {
		result = append(result, strings.TrimSpace(s))
	}
	return result
}
