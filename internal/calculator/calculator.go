package calculator

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"math/big"
)

type Calculator struct {
	app.Compo
	number1   string
	operation string
	number2   string
	isResult  bool
}

var rowStyles = map[string]string{
	"display": "flex",
}

var outputStyles = map[string]string{
	"width":         "248px",
	"height":        "50px",
	"border-style":  "solid",
	"border-radius": "5px",
	"display":       "flex",
	"align-items":   "center",
	"font-size":     "xx-large",
	"margin":        "5px",
}

func (c *Calculator) Render() app.UI {
	textToShow := c.number1
	if c.operation != "" {
		if c.number2 == "" {
			textToShow = c.operation
		} else {
			textToShow = c.number2
		}
	}
	return app.Div().Body(
		app.Div().Body(
			app.Div().Text(textToShow),
		).Styles(outputStyles),
		app.Div().Body(
			&CalcButton{character: "7", onClickHandler: c.concatDigit},
			&CalcButton{character: "8", onClickHandler: c.concatDigit},
			&CalcButton{character: "9", onClickHandler: c.concatDigit},
			&CalcButton{character: "/", onClickHandler: c.useOperator},
		).Styles(rowStyles),
		app.Div().Body(
			&CalcButton{character: "4", onClickHandler: c.concatDigit},
			&CalcButton{character: "5", onClickHandler: c.concatDigit},
			&CalcButton{character: "6", onClickHandler: c.concatDigit},
			&CalcButton{character: "*", onClickHandler: c.useOperator},
		).Styles(rowStyles),
		app.Div().Body(
			&CalcButton{character: "1", onClickHandler: c.concatDigit},
			&CalcButton{character: "2", onClickHandler: c.concatDigit},
			&CalcButton{character: "3", onClickHandler: c.concatDigit},
			&CalcButton{character: "-", onClickHandler: c.useOperator},
		).Styles(rowStyles),
		app.Div().Body(
			&CalcButton{character: "0", onClickHandler: c.concatDigit},
			&CalcButton{character: ".", onClickHandler: c.concatDigit},
			&CalcButton{character: "+", onClickHandler: c.useOperator},
			&CalcButton{character: "=", onClickHandler: c.calculate},
		).Styles(rowStyles),
	).Styles(map[string]string{
		"display":        "flex",
		"align-items":    "center",
		"flex-direction": "column",
	})
}

func (c *Calculator) calculate(ctx app.Context, _ app.Event) {
	number1, _, _ := big.ParseFloat(c.number1, 10, 64, big.ToNearestEven)
	number2, _, _ := big.ParseFloat(c.number2, 10, 64, big.ToNearestEven)
	var total big.Float
	total.SetPrec(64)
	switch c.operation {
	case "+":
		total.Add(number1, number2)
	case "-":
		total.Sub(number1, number2)
	case "*":
		total.Mul(number1, number2)
	case "/":
		total.Quo(number1, number2)
	}
	c.number1 = total.String()
	c.operation = ""
	c.number2 = ""
	if ctx.JSSrc().Get("innerText").String() == "=" {
		c.isResult = true
	}
}

func (c *Calculator) concatDigit(ctx app.Context, _ app.Event) {
	fmt.Println(c.isResult)
	if c.isResult {
		c.number1 = ctx.JSSrc().Get("innerText").String()
		c.isResult = false
	} else {
		if c.operation == "" {
			c.number1 += ctx.JSSrc().Get("innerText").String()
		} else {
			c.number2 += ctx.JSSrc().Get("innerText").String()
		}
	}
}

func (c *Calculator) useOperator(ctx app.Context, e app.Event) {
	if c.number2 != "" {
		c.calculate(ctx, e)
	}
	c.operation = ctx.JSSrc().Get("innerText").String()
	c.isResult = false
}
