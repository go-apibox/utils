package utils

import (
	"regexp"
)

func IsPrivateIp(ip string) bool {
	// REF: http://stackoverflow.com/questions/2814002/private-ip-address-identifier-in-regular-expression
	pattern := `(^127\.)|(^192\.168\.)|(^10\.)|(^172\.1[6-9]\.)|(^172\.2[0-9]\.)|(^172\.3[0-1]\.)`
	matched, _ := regexp.MatchString(pattern, ip)
	return matched
}
