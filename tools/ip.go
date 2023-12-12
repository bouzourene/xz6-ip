package tools

import "strings"

func IpVersion(ip string) string {
	if strings.Contains(ip, ":") {
		return "v6"
	}

	return "v4"
}
