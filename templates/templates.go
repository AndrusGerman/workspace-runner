package templates

import (
	_ "embed"
)

//go:embed home.html
var HomeTemplate []byte

//go:embed add-workspace.html
var AddWorkspaceTemplate []byte
