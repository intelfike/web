package proc

import (
	"github.com/intelfike/mulpage/types"
)

type Thanks struct{}

func (p *Thanks) Define(pack *types.Package) {
	pack.Init("Thanks")

	pack.SetMethod("Index", "Thanks", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})
}
