// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "strings"
import "time"

func postUrl(postdate string) string {
	return "/nightly.changelog.com/" + strings.ReplaceAll(postdate, "-", "/")
}

func today() string {
	loc, _ := time.LoadLocation("America/Chicago")
	date := time.Now().In(loc)

	// 10 PM or later
	if date.Hour() >= 22 {
		return date.Format("2006-01-02")
	}
	return date.Add(-24 * time.Hour).Format("2006-01-02")
}

func todayUri() string {
	return "/browse/" + today()
}

func prevUri(postdate string) string {
	if postdate == "2015-01-01" {
		return ""
	}
	date, err := time.Parse("2006-01-02", postdate)
	if err != nil {
		return ""
	}
	previousDay := date.AddDate(0, 0, -1)
	return "/browse/" + previousDay.Format("2006-01-02")
}

func nextUri(postdate string) string {
	if postdate == today() {
		return ""
	}
	date, err := time.Parse("2006-01-02", postdate)
	if err != nil {
		return ""
	}
	nextDay := date.AddDate(0, 0, 1)
	return "/browse/" + nextDay.Format("2006-01-02")
}

func body(postdate string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div hx-ext=\"response-targets\"><div class=\"btm-right\" hx-ext=\"path-params\"><sl-input type=\"date\" name=\"postdate\" placeholder=\"Date\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(postdate)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `body.templ`, Line: 56, Col: 19}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" min=\"2015-01-01\" max=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(today())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `body.templ`, Line: 58, Col: 16}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-get=\"/browse/:postdate\" hx-push-url=\"true\" hx-trigger=\"input changed delay:1s\" hx-target=\"body\"></sl-input><div><sl-icon-button library=\"default\" name=\"chevron-double-left\" hx-push-url=\"true\" hx-get=\"/browse/2015-01-01\" hx-target=\"body\"></sl-icon-button> <sl-icon-button library=\"default\" name=\"chevron-left\" hx-push-url=\"true\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(prevUri(postdate))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `body.templ`, Line: 77, Col: 30}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"body\" class=\"prev-date\"></sl-icon-button> <sl-icon-button library=\"default\" name=\"chevron-right\" hx-push-url=\"true\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(nextUri(postdate))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `body.templ`, Line: 85, Col: 30}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"body\" class=\"next-date\"></sl-icon-button> <sl-icon-button library=\"default\" name=\"chevron-double-right\" hx-push-url=\"true\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(todayUri())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `body.templ`, Line: 94, Col: 23}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"body\"></sl-icon-button></div></div><div id=\"main\" tabindex=\"0\" style=\"height:100vh\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(postUrl(postdate))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `body.templ`, Line: 104, Col: 28}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-swap=\"innerHTML\" hx-target=\"#main\" hx-target-404=\"#main\" hx-indicator=\"#spinner\" hx-trigger=\"load\"></div><style>\n\t.htmx-indicator {\n\t\topacity: 0;\n\t\ttransition: opacity 200ms ease-in;\n\t}\n\t.htmx-request .htmx-indicator {\n\t\topacity: 1;\n\t}\n\t.htmx-request.htmx-indicator {\n\t\topacity: 1;\n\t}\n\t#spinner {\n\t\tposition: fixed;\n\t\tbottom: 49%;\n\t\tleft: 49%;\n\t\tscale: 2;\n\t\tfill: #56ad83;\n\t}\n\t</style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = loadingSvg().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func loadingSvg() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var8 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var8 == nil {
			templ_7745c5c3_Var8 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg id=\"spinner\" class=\"htmx-indicator\" width=\"17\" height=\"16\" viewBox=\"0 0 135 140\" xmlns=\"http://www.w3.org/2000/svg\"><rect y=\"10\" width=\"15\" height=\"120\" rx=\"6\"><animate attributeName=\"height\" begin=\"0.5s\" dur=\"1s\" values=\"120;110;100;90;80;70;60;50;40;140;120\" calcMode=\"linear\" repeatCount=\"indefinite\"></animate> <animate attributeName=\"y\" begin=\"0.5s\" dur=\"1s\" values=\"10;15;20;25;30;35;40;45;50;0;10\" calcMode=\"linear\" repeatCount=\"indefinite\"></animate></rect> <rect x=\"30\" y=\"10\" width=\"15\" height=\"120\" rx=\"6\"><animate attributeName=\"height\" begin=\"0.25s\" dur=\"1s\" values=\"120;110;100;90;80;70;60;50;40;140;120\" calcMode=\"linear\" repeatCount=\"indefinite\"></animate> <animate attributeName=\"y\" begin=\"0.25s\" dur=\"1s\" values=\"10;15;20;25;30;35;40;45;50;0;10\" calcMode=\"linear\" repeatCount=\"indefinite\"></animate></rect> <rect x=\"60\" width=\"15\" height=\"140\" rx=\"6\"><animate attributeName=\"height\" begin=\"0s\" dur=\"1s\" values=\"120;110;100;90;80;70;60;50;40;140;120\" calcMode=\"linear\" repeatCount=\"indefinite\"></animate> <animate attributeName=\"y\" begin=\"0s\" dur=\"1s\" values=\"10;15;20;25;30;35;40;45;50;0;10\" calcMode=\"linear\" repeatCount=\"indefinite\"></animate></rect> <rect x=\"90\" y=\"10\" width=\"15\" height=\"120\" rx=\"6\"><animate attributeName=\"height\" begin=\"0.25s\" dur=\"1s\" values=\"120;110;100;90;80;70;60;50;40;140;120\" calcMode=\"linear\" repeatCount=\"indefinite\"></animate> <animate attributeName=\"y\" begin=\"0.25s\" dur=\"1s\" values=\"10;15;20;25;30;35;40;45;50;0;10\" calcMode=\"linear\" repeatCount=\"indefinite\"></animate></rect> <rect x=\"120\" y=\"10\" width=\"15\" height=\"120\" rx=\"6\"><animate attributeName=\"height\" begin=\"0.5s\" dur=\"1s\" values=\"120;110;100;90;80;70;60;50;40;140;120\" calcMode=\"linear\" repeatCount=\"indefinite\"></animate> <animate attributeName=\"y\" begin=\"0.5s\" dur=\"1s\" values=\"10;15;20;25;30;35;40;45;50;0;10\" calcMode=\"linear\" repeatCount=\"indefinite\"></animate></rect></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
