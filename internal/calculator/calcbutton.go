package calculator

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type CalcButton struct {
	app.Compo
	character      string
	onClickHandler app.EventHandler
}

var buttonStyles = map[string]string{
	"width":           "50px",
	"height":          "50px",
	"border-style":    "solid",
	"border-radius":   "6px",
	"display":         "flex",
	"justify-content": "center",
	"align-items":     "center",
	"margin":          "5px",
	"font-size":       "larger",
}

func (d *CalcButton) Render() app.UI {
	return app.Div().Text(d.character).OnClick(d.onClickHandler).Styles(buttonStyles)
}
