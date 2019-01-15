package proc

import (
	"github.com/intelfike/mulpage/types"
)

type Usage struct{}

func (p *Usage) Define(pack *types.Package) {
	pack.Init("isearの使い方")

	pack.Before = func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	}

	pack.SetMethod("Index", "", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return &types.Redirect{"/_usage/Functions", 301}
	})

	pack.SetMethod("Functions", "機能", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})

	pack.SetMethod("Settings", "設定", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})

	pack.SetMethod("Name", "名称・用語", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})

	pack.SetMethod("SKey", "ショートカットキー", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})

	pack.SetMethod("SWord", "検索テキスト", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})

	pack.SetMethod("CommandMode", "コマンドモード", func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		return nil
	})
}
