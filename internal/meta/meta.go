package meta

type PageMeta struct {
	Title       string
	Description string
	OGType      string
	OGImage     string
	Canonical   string
	NoIndex     bool
}

func New(title, description string) PageMeta {
	return PageMeta{
		Title:       title,
		Description: description,
		OGType:      "website",
	}
}

func (m PageMeta) WithOGImage(url string) PageMeta {
	m.OGImage = url
	return m
}

func (m PageMeta) AsArticle() PageMeta {
	m.OGType = "article"
	return m
}
