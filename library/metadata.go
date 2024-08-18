package library

type ChapterMetadata struct {
	Slug  string   `toml:"slug"`
	Name  string   `toml:"name"`
	Pages []string `toml:"pages"`
}

type SerieMetadata struct {
	Slug     string            `toml:"slug"`
	Title    string            `toml:"title"`
	CoverArt string            `toml:"coverArt"`
	Chapters []ChapterMetadata `toml:"chapters"`

	path    string
	new     bool
	changed bool
}

func (s *SerieMetadata) MarkChanged() {
	s.changed = true
}

func (s SerieMetadata) Path() string {
	return s.path
}
