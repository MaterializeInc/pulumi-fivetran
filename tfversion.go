package tfversion

import (
	_ "embed"

	"golang.org/x/mod/modfile"
)

var TFVersion string

//go:embed go.mod
var gomod []byte

func init() {
	f, err := modfile.Parse("go.mod", gomod, nil)
	if err != nil {
		panic(err)
	}
	for _, r := range f.Require {
		if r.Mod.Path == "github.com/benesch/terraform-provider-fivetran" {
			TFVersion = r.Mod.Version
		}
	}
}
