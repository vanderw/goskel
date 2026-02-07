package utils

import "testing"

func TestNormalizeIp(t *testing.T) {
	ip := "0083::fe01"
	t.Log(NomalizeIpv6(ip))
}
