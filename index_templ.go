// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func index(postdate string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html><head><style>\n\t\t\t\tbody {\n\t\t\t\t\tbackground-color: #101820;\n\t\t\t\t}\n\t\t\t\t.btm-right {\n\t\t\t\t\tdisplay: flex;\n\t\t\t\t\tcolor: #5ebc8f;\n\t\t\t\t\tbackground: #000;\n\t\t\t\t\tposition: fixed;\n\t\t\t\t\tbottom: 0;\n\t\t\t\t\tright: 0;\n\t\t\t\t\theight: 1.8rem;\n\t\t\t\t\tz-index: 222;\n\t\t\t\t\tmargin: 14px;\n\t\t\t\t\tjustify-content: center;\n\t\t\t\t\talign-items: center;\n\t\t\t\t\tborder-radius: 10px;\n\t\t\t\t\tpadding-right: 6px;\n\t\t\t\t}\n\t\t\t\tsl-input {\n\t\t\t\t\tposition: relative;\n\t\t\t\t\tbottom: 2px;\n\t\t\t\t\tpadding-left: 5px;\n\t\t\t\t\tpadding-right: 5px;\n\t\t\t\t}\n\t\t\t\tsl-icon-button:focus {\n\t\t\t\t\tcolor: aliceblue;\n\t\t\t\t}\n\t\t\t\t#main:focus {\n\t\t\t\t\toutline: none;\n\t\t\t\t}\n\t\t\t</style><link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.14.0/cdn/themes/dark.css\"><script type=\"module\" src=\"https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.14.0/cdn/components/input/input.js\"></script><script type=\"module\" src=\"https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.14.0/cdn/components/icon-button/icon-button.js\"></script><script type=\"module\">\n\t\t\t\timport { registerIconLibrary, unregisterIconLibrary } from 'https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.14.0/cdn/utilities/icon-library.js';\n\t\t\t\tconsole.log(\"registering bootstrap-icons 1.11.3\");\n\t\t\t\tunregisterIconLibrary(\"default\");\n\t\t\t\tregisterIconLibrary(\"default\", {\n\t\t\t\t\tresolver: name => `https://cdn.jsdelivr.net/npm/bootstrap-icons@1.11.3/icons/${name}.svg`\n\t\t\t\t});\n\t\t\t</script></head><body hx-ext=\"head-support shoelace\"><script src=\"/dist/js/htmx.min.js\"></script><script src=\"/dist/js/htmx.ext.head-support.js\"></script><script src=\"/dist/js/htmx.ext.response-targets.js\"></script><script src=\"/dist/js/htmx.ext.shoelace.js\"></script><script>\n\t\t\t\t// https://stackoverflow.com/a/68058914/8608146\n\t\t\t\t// https://github.com/bigskysoftware/htmx/issues/1202\n\t\t\t\thtmx.defineExtension('path-params', {\n\t\t\t\t\tonEvent: function(name, evt) {\n\t\t\t\t\t\tif (name === \"htmx:configRequest\") {\n\t\t\t\t\t\t\tevt.detail.path = evt.detail.path.replace(/:([A-Za-z0-9_]+)/g, function (_, param) {\n\t\t\t\t\t\t\t\tlet val = evt.detail.parameters[param];\n\t\t\t\t\t\t\t\tdelete evt.detail.parameters[param];\n\t\t\t\t\t\t\t\treturn val;\n\t\t\t\t\t\t\t})\n\t\t\t\t\t\t}\n\t\t\t\t\t}\n\t\t\t\t});\n\t\t\t\tdocument.addEventListener('DOMContentLoaded', function() {\n\t\t\t\t\tdocument.querySelector(\"#main\").focus();\n\t\t\t\t});\n\t\t\t\tdocument.body.addEventListener('htmx:afterSwap', function (e) {\n\t\t\t\t\tlet mainElt = document.querySelector(\"#main\");\n\t\t\t\t\tmainElt.focus();\n\t\t\t\t\tmainElt\n\t\t\t\t\t\t.addEventListener(\"keydown\", function (evt) {\n\t\t\t\t\t\tif (evt.altKey || evt.ctrlKey || evt.shiftKey) return;\n\t\t\t\t\t\tif (evt.keyCode === 72 || evt.keyCode === 37) {\n\t\t\t\t\t\t\tevt.preventDefault();\n\t\t\t\t\t\t\tdocument.querySelector(\".prev-date\").click();\n\t\t\t\t\t\t} else if (evt.keyCode === 76 || evt.keyCode === 39) {\n\t\t\t\t\t\t\tevt.preventDefault();\n\t\t\t\t\t\t\tdocument.querySelector(\".next-date\").click();\n\t\t\t\t\t\t}\n\t\t\t\t\t});\n\t\t\t\t});\n\t\t\t</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = body(postdate).Render(ctx, templ_7745c5c3_Buffer)
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
