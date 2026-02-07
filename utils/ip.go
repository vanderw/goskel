package utils

import (
	"net"
	"net/netip"
	"strings"
)

func ValidIpv4(ip string) bool {
	parsed := net.ParseIP(ip)
	if parsed == nil {
		return false
	}
	if parsed.To4() == nil {
		return false
	}
	return true
}

func ValidIpv6(ip string) bool {
	parsed := net.ParseIP(ip)
	return parsed != nil && strings.Contains(ip, ":")
}

func ValidIp(ip string) bool {
	return ValidIpv4(ip) || ValidIpv6(ip)
}

func NomalizeIpv6(ip string) string {
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		return ip
	}
	return addr.StringExpanded()
}
