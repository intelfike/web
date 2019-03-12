package thanks

import (
	"github.com/intelfike/mulpage/types"
)

type Page struct{}

func (p *Page) Define(parent *types.Package) *types.Package {
	child := parent.NewChild("thanks")

	child.SetMethod("index", func(info types.PageInfo, tpl *types.TplData) *types.Redirect {
		tpl.Assign("Title", "Thanks")
		return nil
	})
	return child
}
