package www

import (
	"github.com/intelfike/mulpage/types"
	"github.com/intelfike/web/contents/www/proc"
)

type Content struct{}

func (c *Content) Define(con *types.Package) {
	con.Init("intelfikeのHP")

	topPack := con.NewChild("top")
	topPack.FallDown(&proc.Top{})

	worksPack := con.NewChild("works")
	worksPack.FallDown(&proc.Works{})

	con.Before = func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		tpl.Assign("Packages", con.ChildrenArray())
		tpl.Assign("Package", con.Children[info.Package])
		if v, ok := con.Children[info.Package]; ok {
			if v1, ok1 := v.Articles[info.Article]; ok1 {
				tpl.Assign("Title", v1.Name)
			}
		}
		tpl.Assign("Info", info)
		return nil
	}

	con.After = func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		tpl.Assign("Template", tpl.Template)
		return nil
	}
}
