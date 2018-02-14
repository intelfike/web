package proc

import (
	"github.com/intelfike/mulpage/types"
)

type Top struct{}

func (p *Top) Define(pack *types.Package) {
	pack.Init("top")

	pack.SetMethod("Index", "top", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})
}
