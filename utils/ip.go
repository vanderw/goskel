package utils

import (
	"net"
	"net/http"
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

// Transform '::' to propriate numbers of '0000'
func NomalizeIpv6(ip string) string {
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		return ip
	}
	return addr.StringExpanded()
}

func GetRemoteIp(r *http.Request) string {
	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		ips := strings.Split(forwarded, ",")
		for _, ip := range ips {
			realIP := net.ParseIP(strings.TrimSpace(ip))
			if !realIP.IsLoopback() && realIP.To4() != nil {
				return ip
			}
		}
	}
	// ip:port
	remoteAddr := r.RemoteAddr
	ip, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return ""
	}

	realIP2 := net.ParseIP(ip)
	if realIP2 == nil || realIP2.IsLoopback() {
		return ""
	}

	return ip
}
