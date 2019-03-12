package contents

import (
	"github.com/intelfike/mulpage/types"
	"github.com/intelfike/web/contents/isear"
)

type App struct{}

func (c *App) Define(app *types.Package) *types.Package {
	app.FallDown(
		&isear.Contents{},
		// &www.Contents{},
	)
	return app
}
