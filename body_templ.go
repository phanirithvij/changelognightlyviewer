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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n\t\t.btm-right {\n\t\t\tdisplay: flex;\n\t\t\tcolor: #5ebc8f;\n\t\t\tbackground: #000;\n\t\t\tposition: fixed;\n\t\t\tbottom: 0;\n\t\t\tright: 0;\n\t\t\theight: 1.8rem;\n\t\t\tz-index: 222;\n\t\t\tmargin: 14px;\n\t\t\tjustify-content: center;\n\t\t\talign-items: center;\n\t\t\tborder-radius: 10px;\n\t\t\tpadding-right: 6px;\n\t\t}\n\t\tsl-input {\n\t\t\tposition: relative;\n\t\t\tbottom: 2px;\n\t\t\tpadding-left: 5px;\n\t\t\tpadding-right: 5px;\n\t\t}\n\t</style><div class=\"btm-right\" hx-ext=\"path-params\"><sl-input type=\"date\" name=\"postdate\" placeholder=\"Date\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(postdate)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `body.templ`, Line: 78, Col: 19}
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
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `body.templ`, Line: 80, Col: 16}
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
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `body.templ`, Line: 99, Col: 30}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"body\"></sl-icon-button> <sl-icon-button library=\"default\" name=\"chevron-right\" hx-push-url=\"true\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(nextUri(postdate))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `body.templ`, Line: 106, Col: 30}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"body\"></sl-icon-button> <sl-icon-button library=\"default\" name=\"chevron-double-right\" hx-push-url=\"true\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(todayUri())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `body.templ`, Line: 114, Col: 23}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"body\"></sl-icon-button></div></div><div hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(postUrl(postdate))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `body.templ`, Line: 120, Col: 32}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-swap=\"outerHTML\" hx-trigger=\"load\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
