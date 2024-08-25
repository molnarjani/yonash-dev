package types

import (
	"context"
)

type contextKey string

var CDNUrlContextKey contextKey = "cdn_url"
var IndexPageContentContextKey contextKey = "page_content"

func GetCDNUrl(ctx context.Context) string {
	if cdnUrl, ok := ctx.Value(CDNUrlContextKey).(string); ok {
		return cdnUrl
	}
	return ""
}
