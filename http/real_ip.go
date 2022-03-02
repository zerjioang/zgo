package http

import (
	"bytes"
	"github.com/valyala/fasthttp"
)

const (
	HeaderXForwardedFor = "X-Forwarded-For"
	HeaderXRealIP       = "X-Real-IP"
)

// FindRealIP returns real IP of the HTTP client
func FindRealIP(h *fasthttp.RequestHeader, fallback func() []byte) []byte {
	if h == nil {
		return []byte{}
	}
	// Fallback to legacy behavior
	if ip := h.Peek(HeaderXForwardedFor); ip != nil && len(ip) > 0 {
		i := bytes.IndexAny(ip, ",")
		if i > 0 {
			return bytes.TrimSpace(ip[:i])
		}
		return ip
	}
	if ip := h.Peek(HeaderXRealIP); ip != nil && len(ip) > 0 {
		return ip
	}
	if ip := h.Peek("CF-Connecting-IP"); ip != nil && len(ip) > 0 {
		return ip
	}
	return fallback()
}
