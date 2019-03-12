package isear

import (
	"github.com/intelfike/mulpage/types"
	"github.com/intelfike/web/contents/isear/page/install"
	"github.com/intelfike/web/contents/isear/page/support"
	"github.com/intelfike/web/contents/isear/page/thanks"
	"github.com/intelfike/web/contents/isear/page/top"
	"github.com/intelfike/web/contents/isear/page/usage"
)

type Contents struct{}

func (p *Contents) Define(parent *types.Package) *types.Package {
	child := parent.NewChild("isear")
	child.SetBefore(func(info types.PageInfo, tpl *types.TplData) *types.Redirect {
		if info.Group == "" {
			return types.NewRedirect("/_top/index", 302)
		}
		return nil
	})
	child.FallDown(
		&install.Page{},
		&support.Page{},
		&thanks.Page{},
		&top.Page{},
		&usage.Page{},
	)
	return child
}
