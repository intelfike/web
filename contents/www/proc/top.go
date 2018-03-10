package proc

import (
	"github.com/intelfike/mulpage/types"
)

type Top struct{}

func (p *Top) Define(pack *types.Package) {
	pack.Init("トップページ")

	pack.SetMethod("Index", "トップページ", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})
}
