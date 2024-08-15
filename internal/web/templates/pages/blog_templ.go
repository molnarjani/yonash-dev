// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func BlogPage() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"bg-white\"><div class=\"container mx-auto py-8\"><div class=\"col-span-4 sm:col-span-9\"><div class=\"py-8 lg:py-16 px-4 mx-auto max-w-screen-md\"><p class=\"mb-4 lg:mb-8 font-light text-center text-gray-700 dark:text-gray-400 sm:text-xl\">No posts here yet, check back later...</p></div><section class=\"bg-white dark:bg-gray-900 flex justify-center items-center\"><div role=\"status\" class=\"max-w-md p-4 space-y-4 border border-gray-200 divide-y divide-gray-200 rounded shadow animate-pulse dark:divide-gray-700 md:p-6 dark:border-gray-700\"><div class=\"flex items-center justify-between\"><div><div class=\"h-2.5 bg-gray-300 rounded-full dark:bg-gray-600 w-24 mb-2.5\"></div><div class=\"w-32 h-2 bg-gray-200 rounded-full dark:bg-gray-700\"></div></div><div class=\"h-2.5 bg-gray-300 rounded-full dark:bg-gray-700 w-12\"></div></div><div class=\"flex items-center justify-between pt-4\"><div><div class=\"h-2.5 bg-gray-300 rounded-full dark:bg-gray-600 w-24 mb-2.5\"></div><div class=\"w-32 h-2 bg-gray-200 rounded-full dark:bg-gray-700\"></div></div><div class=\"h-2.5 bg-gray-300 rounded-full dark:bg-gray-700 w-12\"></div></div><div class=\"flex items-center justify-between pt-4\"><div><div class=\"h-2.5 bg-gray-300 rounded-full dark:bg-gray-600 w-24 mb-2.5\"></div><div class=\"w-32 h-2 bg-gray-200 rounded-full dark:bg-gray-700\"></div></div><div class=\"h-2.5 bg-gray-300 rounded-full dark:bg-gray-700 w-12\"></div></div><div class=\"flex items-center justify-between pt-4\"><div><div class=\"h-2.5 bg-gray-300 rounded-full dark:bg-gray-600 w-24 mb-2.5\"></div><div class=\"w-32 h-2 bg-gray-200 rounded-full dark:bg-gray-700\"></div></div><div class=\"h-2.5 bg-gray-300 rounded-full dark:bg-gray-700 w-12\"></div></div><div class=\"flex items-center justify-between pt-4\"><div><div class=\"h-2.5 bg-gray-300 rounded-full dark:bg-gray-600 w-24 mb-2.5\"></div><div class=\"w-32 h-2 bg-gray-200 rounded-full dark:bg-gray-700\"></div></div><div class=\"h-2.5 bg-gray-300 rounded-full dark:bg-gray-700 w-12\"></div></div><span class=\"sr-only\">Loading...</span></div></section></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
