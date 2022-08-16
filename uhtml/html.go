package uhtml

type Source struct {
	Name    string
	Format  string
	Explain string
}

type Result struct {
	Titles      []string
	Dates       []string
	Authors     []string
	Identifiers []string
	Imprints    []string
}
