package proc

import (
	"github.com/intelfike/mulpage/types"
)

type Install struct{}

func (p *Install) Define(pack *types.Package) {
	pack.Init("インストール")

	pack.SetMethod("Index", "インストール", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})
}
