// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func index(postdate string, jsenabled bool) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html><head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if jsenabled {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<noscript><meta http-equiv=\"refresh\" content=\"5.0; url=/browsehtml\"></noscript>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<link rel=\"stylesheet\" href=\"/dist/css/index.css\"></head><body hx-ext=\"head-support\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if jsenabled {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<noscript><div style=\"color: wheat\"><p>This site is best viewed with Javascript.</p><p>Automatically redirecting to the html version in 5 sec.</p></div></noscript>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script src=\"/dist/js/htmx.min.js\"></script><script src=\"/dist/js/htmx.ext.head-support.js\"></script><script src=\"/dist/js/htmx.ext.response-targets.js\"></script><script src=\"/dist/js/quicklink.umd.js\"></script><script>\n\t\t\t\t// https://stackoverflow.com/a/68058914/8608146\n\t\t\t\t// https://github.com/bigskysoftware/htmx/issues/1202\n\t\t\t\thtmx.defineExtension('path-params', {\n\t\t\t\t\tonEvent: function(name, evt) {\n\t\t\t\t\t\tif (name === \"htmx:configRequest\") {\n\t\t\t\t\t\t\tevt.detail.path = evt.detail.path.replace(/:([A-Za-z0-9_]+)/g, function (_, param) {\n\t\t\t\t\t\t\t\tlet val = evt.detail.parameters[param];\n\t\t\t\t\t\t\t\tdelete evt.detail.parameters[param];\n\t\t\t\t\t\t\t\treturn val;\n\t\t\t\t\t\t\t})\n\t\t\t\t\t\t}\n\t\t\t\t\t}\n\t\t\t\t});\n\t\t\t\tdocument.addEventListener('DOMContentLoaded', function() {\n\t\t\t\t\tdocument.querySelector(\"#main\").focus();\n\t\t\t\t});\n\t\t\t\tdocument.body.addEventListener('htmx:afterSwap', function (e) {\n\t\t\t\t\tlet mainElt = document.querySelector(\"#main\");\n\t\t\t\t\tmainElt.focus();\n\t\t\t\t\tmainElt\n\t\t\t\t\t\t.addEventListener(\"keydown\", function (evt) {\n\t\t\t\t\t\tif (evt.altKey || evt.ctrlKey || evt.shiftKey) return;\n\t\t\t\t\t\tif (evt.keyCode === 72 || evt.keyCode === 37) {\n\t\t\t\t\t\t\tevt.preventDefault();\n\t\t\t\t\t\t\tdocument.querySelector(\".prev-date\").click();\n\t\t\t\t\t\t} else if (evt.keyCode === 76 || evt.keyCode === 39) {\n\t\t\t\t\t\t\tevt.preventDefault();\n\t\t\t\t\t\t\tdocument.querySelector(\".next-date\").click();\n\t\t\t\t\t\t}\n\t\t\t\t\t});\n\t\t\t\t\t//document.querySelector('#instantjs')?.remove();\n\t\t\t\t\t//let instant = document.createElement('script');\n\t\t\t\t\t//instant.id = \"instantjs\";\n\t\t\t\t\t//instant.src = \"/dist/js/instant.js\";\n\t\t\t\t\t//document.body.appendChild(instant);\n\n\t\t\t\t\t//quicklink.listen();\n\t\t\t\t\tlet td = new Date(document.location.href.split('/')[4]);\n\t\t\t\t\ttd.setDate(td.getDate() - 1);\n\t\t\t\t\tlet prevDay = `/browse/${td.toISOString().split('T')[0]}`;\n\t\t\t\t\ttd.setDate(td.getDate() + 2);\n\t\t\t\t\tlet nextDay = `/browse/${td.toISOString().split('T')[0]}`;\n\t\t\t\t\tconsole.log(\"prefetching\", prevDay, nextDay);\n\t\t\t\t\tquicklink.prefetch([nextDay, prevDay]);\n\t\t\t\t});\n\t\t\t</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = body(postdate, jsenabled).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
