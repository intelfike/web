package contents

import (
	"path/filepath"

	"github.com/intelfike/mulpage/global"
	"github.com/intelfike/mulpage/types"
	"github.com/intelfike/web/contents/isear"
	"github.com/intelfike/web/contents/www"
)

func init() {
	// コンテンツリストを定義
	global.App.Init("web")
	global.App.Kind = "app"
	global.App.ChildKind = "contents"

	global.App.Before = func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		// テンプレートデフォルト情報を定義
		tpl.RootPath = filepath.Join("contents", info.Contents, "template")
		tpl.TemplatePath = filepath.Join("page", info.Package)
		tpl.LayoutPath = "layout"
		tpl.Template = info.Article
		tpl.Layout = "default"
		return nil
	}
	global.App.After = func(tpl *types.TplData, info types.PageInfo) *types.Redirect {
		tpl.AddTpl(tpl.GetLayoutPath(), tpl.GetTemplatePath())
		return nil
	}

	wwwCon := global.App.NewChild("www")
	wwwCon.ChildKind = "package"
	wwwCon.FallDown(&www.Content{})

	isearCon := global.App.NewChild("isear")
	isearCon.ChildKind = "package"
	isearCon.FallDown(&isear.Content{})
}
