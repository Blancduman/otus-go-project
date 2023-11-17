package migration

import "embed"

//go:embed *
var FS embed.FS

const (
	CatalogPath = "catalog"
)
