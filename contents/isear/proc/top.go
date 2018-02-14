package proc

import (
	"github.com/intelfike/mulpage/types"
)

type Top struct{}

func (p *Top) Define(pack *types.Package) {
	pack.Init("トップページ")

	pack.Before = func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		tpl.Assign("Title", "トップページ")
		return nil
	}

	pack.SetMethod("Index", "トップページ", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})
}
