package usage

import (
	"github.com/intelfike/mulpage/types"
)

type Page struct{}

func (p *Page) Define(parent *types.Package) *types.Package {
	child := parent.NewChild("usage")

	child.SetMethod("index", func(info types.PageInfo, tpl *types.TplData) *types.Redirect {
		tpl.Assign("Title", "一覧")
		return nil
	})
	child.SetMethod("functions", func(info types.PageInfo, tpl *types.TplData) *types.Redirect {
		tpl.Assign("Title", "機能")
		return nil
	})
	child.SetMethod("settings", func(info types.PageInfo, tpl *types.TplData) *types.Redirect {
		tpl.Assign("Title", "設定")
		return nil
	})
	child.SetMethod("name", func(info types.PageInfo, tpl *types.TplData) *types.Redirect {
		tpl.Assign("Title", "名称・用語")
		return nil
	})
	child.SetMethod("sKey", func(info types.PageInfo, tpl *types.TplData) *types.Redirect {
		tpl.Assign("Title", "ショートカットキー")
		return nil
	})
	child.SetMethod("sWord", func(info types.PageInfo, tpl *types.TplData) *types.Redirect {
		tpl.Assign("Title", "検索テキスト")
		return nil
	})
	child.SetMethod("commandMode", func(info types.PageInfo, tpl *types.TplData) *types.Redirect {
		tpl.Assign("Title", "コマンドモード")
		return nil
	})
	return child
}
