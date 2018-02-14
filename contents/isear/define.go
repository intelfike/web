package isear

import (
	"github.com/intelfike/mulpage/types"
	"github.com/intelfike/web/contents/isear/proc"
)

type Content struct{}

func (c *Content) Define(con *types.Package) {
	con.Init("isear")

	topPack := con.NewChild("top")
	topPack.FallDown(&proc.Top{})

	usagePack := con.NewChild("usage")
	usagePack.FallDown(&proc.Usage{})

	supportPack := con.NewChild("support")
	supportPack.FallDown(&proc.Support{})

	installPack := con.NewChild("install")
	installPack.FallDown(&proc.Install{})

	thanksPack := con.NewChild("thanks")
	thanksPack.FallDown(&proc.Thanks{})

	con.Before = func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		// ページに必要な情報を渡す
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
