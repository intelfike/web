package proc

import (
	"github.com/intelfike/mulpage/types"
)

type Support struct{}

func (p *Support) Define(pack *types.Package) {
	pack.Init("サポート情報")

	pack.SetMethod("Index", "サポート情報", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})
}
