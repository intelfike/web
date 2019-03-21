package isear

import (
	"regexp"
	"strings"

	"github.com/intelfike/mulpage/types"
	"github.com/intelfike/web/contents/isear/page/install"
	"github.com/intelfike/web/contents/isear/page/support"
	"github.com/intelfike/web/contents/isear/page/thanks"
	"github.com/intelfike/web/contents/isear/page/top"
	"github.com/intelfike/web/contents/isear/page/usage"
)

type Contents struct{}

var upperCase = regexp.MustCompile("[A-Z]")

func (p *Contents) Define(parent *types.Package) *types.Package {
	child := parent.NewChild("isear")
	child.SetBefore(func(info types.PageInfo, tpl *types.TplData) *types.Redirect {
		if info.Group == "" {
			return types.NewRedirect("/_top/index", 302)
		}
		if upperCase.MatchString(info.Page[:1]) {
			return types.NewRedirect("/_"+info.Group+"/"+strings.ToLower(info.Page), 301)
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
