// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"github.com/molnarjani/yonash-dev/internal/types"
	"github.com/molnarjani/yonash-dev/internal/web/templates/components"
)

// MetaTags defines meta tags.
func MetaTags(keywords, description string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<meta name=\"keywords\" content=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(keywords)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/web/templates/pages/index.templ`, Line: 12, Col: 41}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><meta name=\"description\" content=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(description)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/web/templates/pages/index.templ`, Line: 13, Col: 47}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

// BodyContent defines HTML content.
func BodyContent() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"app\"><div id=\"navbar\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.NavBar().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"sm:content content-center0\"><div id=\"page-content\" class=\"content0\"><section class=\"bg-white dark:bg-gray-900\"><div class=\"grid max-w-screen-xl px-4 py-8 mx-auto xl:gap-0 md:py-8 lg:gap-8 xl:gap-0 lg:py-16 lg:grid-cols-12 md:gap-8 md:py-16 md:grid-cols-12\"><div class=\"opacity-0 animate-fade-in px-8 py-8 md:mt-0 md:col-span-5 lg:mt-0 lg:col-span-5 lg:flex\"><img class=\"rounded-sm shadow-xl\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(string(templ.URL(fmt.Sprintf("%s/static/images/me.jpg", types.GetCDNUrl(ctx)))))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/web/templates/pages/index.templ`, Line: 28, Col: 125}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" alt=\"personal-photo\"></div><div class=\"mr-auto pl-8 place-self-center md:col-span-7 lg:col-span-7\"><h1 style=\"animation-delay: 1s;\" class=\"opacity-0 animate-fade-in max-w-2xl mb-4 text-2xl font-extrabold text-left tracking-tight leading-none md:text-4xl xl:text-3xl dark:text-white\">Hi. I'm Jani.</h1><h2 style=\"animation-delay: 2s;\" class=\"opacity-0 animate-fade-in max-w-2xl mb-4 text-2xl font-extrabold text-left tracking-tight leading-none md:text-4xl xl:text-3xl dark:text-white\">I build performant backends.</h2><p style=\"animation-delay: 3s;\" class=\"opacity-0 animate-fade-in max-w-2xl mb-6 font-light text-gray-500 text-left lg:mb-8 md:text-lg lg:text-xl dark:text-gray-400\">I'm a backend developer based in Budapest, who likes to do fun projects involving Python, Go or other backends.<br>I believe in strong ownership built upon scalability and obervability, I'm often involved in infrastructure and DevOps work as well.</p><a style=\"animation-delay: 3.5s;\" href=\"#\" hx-get=\"/projects\" hx-target=\"#page-content\" class=\"w-60 text-center opacity-0 animate-fade-in shadow-xl float-left justify-content mr-3 mb-3 px-5 py-3 text-base font-medium text-center text-gray-900 border border-gray-300 rounded-lg hover:bg-gray-100 dark:text-white dark:border-gray-700 dark:hover:bg-gray-700 dark:focus:ring-gray-800\">See my projects</a> <a style=\"animation-delay: 3.5s;\" href=\"#\" hx-get=\"/projects\" hx-target=\"#page-content\" class=\"w-60 text-center opacity-0 animate-fade-in shadow-xl float-left justify-content px-5 py-3 text-base font-medium text-center text-gray-900 border border-gray-300 rounded-lg hover:bg-gray-100 dark:text-white dark:border-gray-700 dark:hover:bg-gray-700 dark:focus:ring-gray-800\">Check my work stuff</a></div></div></section></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
