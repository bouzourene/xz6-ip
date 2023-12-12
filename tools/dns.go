package tools

import (
	"net"
)

func DnsReverse(ip string) string {
	results, err := net.LookupAddr(ip)
	if err != nil {
		return ""
	}

	var result string
	for _, address := range results {
		result = address
	}

	if len(result) > 0 && result[len(result)-1:] == "." {
		result = result[:len(result)-1]
	}

	return result
}
