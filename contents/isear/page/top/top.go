package top

import (
	"github.com/intelfike/mulpage/types"
)

type Page struct{}

func (p *Page) Define(parent *types.Package) *types.Package {
	child := parent.NewChild("top")

	child.SetMethod("index", func(info types.PageInfo, tpl *types.TplData) *types.Redirect {
		tpl.Assign("Title", "トップページ")
		return nil
	})
	return child
}
