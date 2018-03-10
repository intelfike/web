package proc

import (
	"github.com/intelfike/mulpage/types"
)

type {{.Package}} struct{}

func (p *{{.Package}}) Define(pack *types.Package) {
	pack.Init("{{.Package}}")

	pack.SetMethod("Index", "{{.Package}}", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})
}
