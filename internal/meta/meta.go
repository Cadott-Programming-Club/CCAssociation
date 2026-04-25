package meta

import "strings"

type PageMeta struct {
	Title         string
	Description   string
	OGType        string
	OGImage       string
	OGImageAlt    string
	OGImageWidth  int
	OGImageHeight int
	OGImageType   string
	Canonical     string
	NoIndex       bool
	Locale        string
}

func New(title, description string) PageMeta {
	return PageMeta{
		Title:         title,
		Description:   description,
		OGType:        "website",
		OGImage:       "/static/images/og-default.png",
		OGImageAlt:    "Cadott Community Association — neighbors since 1970",
		OGImageWidth:  1200,
		OGImageHeight: 630,
		OGImageType:   "image/png",
		Locale:        "en_US",
	}
}

func (m PageMeta) WithOGImage(url string) PageMeta {
	m.OGImage = url
	m.OGImageType = mimeFromExt(url)
	return m
}

func (m PageMeta) WithOGImageInfo(url, alt string, width, height int) PageMeta {
	m.OGImage = url
	m.OGImageAlt = alt
	m.OGImageWidth = width
	m.OGImageHeight = height
	m.OGImageType = mimeFromExt(url)
	return m
}

func (m PageMeta) WithOGImageAlt(alt string) PageMeta {
	m.OGImageAlt = alt
	return m
}

func (m PageMeta) WithCanonical(url string) PageMeta {
	m.Canonical = url
	return m
}

func (m PageMeta) AsArticle() PageMeta {
	m.OGType = "article"
	return m
}

func mimeFromExt(url string) string {
	lower := strings.ToLower(url)
	switch {
	case strings.HasSuffix(lower, ".jpg"), strings.HasSuffix(lower, ".jpeg"):
		return "image/jpeg"
	case strings.HasSuffix(lower, ".png"):
		return "image/png"
	case strings.HasSuffix(lower, ".webp"):
		return "image/webp"
	case strings.HasSuffix(lower, ".gif"):
		return "image/gif"
	default:
		return "image/png"
	}
}
