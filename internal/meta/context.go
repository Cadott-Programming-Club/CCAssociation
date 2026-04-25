package meta

import (
	"context"
	"strings"

	"ccassociation/internal/config"
	"ccassociation/internal/ctxkeys"
)

func SiteFromCtx(ctx context.Context) config.SiteConfig {
	if cfg, ok := ctx.Value(ctxkeys.SiteConfig).(config.SiteConfig); ok {
		return cfg
	}
	return config.SiteConfig{Name: "Cadott Community Association"}
}

func SiteNameFromCtx(ctx context.Context) string {
	return SiteFromCtx(ctx).Name
}

// AbsoluteURL returns href as-is when it already has a scheme; otherwise it
// joins it with the configured SiteConfig.URL. Crawlers and social previews
// require absolute URLs for og:image, so relative paths must be expanded
// before they hit the rendered HTML.
func AbsoluteURL(ctx context.Context, href string) string {
	if href == "" {
		return ""
	}
	if strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") {
		return href
	}
	base := strings.TrimRight(SiteFromCtx(ctx).URL, "/")
	if !strings.HasPrefix(href, "/") {
		href = "/" + href
	}
	return base + href
}
