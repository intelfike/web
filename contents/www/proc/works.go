package proc

import (
	"github.com/intelfike/mulpage/types"
)

type Works struct{}

func (p *Works) Define(pack *types.Package) {
	pack.Init("作品")

	pack.SetMethod("Index", "作品", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})
}
