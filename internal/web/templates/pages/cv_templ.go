// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/molnarjani/yonash-dev/internal/web/templates/components"
)

func CVPage() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"bg-white\"><div class=\"container mx-auto py-8\"><div class=\"grid grid-cols-4 sm:grid-cols-12 gap-6 px-4\"><div class=\"col-span-4 sm:col-span-3\"><div class=\"bg-white shadow-xl rounded-lg p-6\"><div class=\"flex flex-col items-center\"><img src=\"/static/images/cv-pic.jpg\" class=\"w-32 h-32 bg-gray-300 rounded-full mb-4 shrink-0\"></img><h1 class=\"text-xl font-bold\">Janos Molnar</h1><p class=\"text-gray-700\">senior backend engineer</p><!-- social icons -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.SocialIcons().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- contacts button --><div class=\"mt-6 flex flex-wrap gap-4 justify-center\"><a href=\"#\" hx-get=\"/api/pages/contact\" hx-target=\"#page-content\" class=\"shadow-xl float-left inline-flex items-start justify-content px-5 py-3 text-base font-medium text-center text-gray-900 border border-gray-300 rounded-lg hover:bg-gray-100 dark:text-white dark:border-gray-700 dark:hover:bg-gray-700 dark:focus:ring-gray-800\">contact me</a></div></div><hr class=\"my-6 border-t border-gray-300\"><div class=\"flex flex-col text-left mb-8\"><span class=\"text-gray-700 uppercase font-bold tracking-wider mb-2\">programming skills</span><ul><li class=\"mb-0\">Python(5+ years)</li><li class=\"mb-0\">Django, FastAPI, etc</li><li class=\"mb-0\">Go</li><li class=\"mb-0\">Frontent(JS, TS)</li><li class=\"mb-0\">SQL</li></ul></div><div class=\"flex flex-col text-left\"><span class=\"text-gray-700 uppercase font-bold tracking-wider mb-2\">ops skills</span><ul><li class=\"mb-0\">AWS / GCP</li><li class=\"mb-0\">Kubernetes / Helm</li><li class=\"mb-0\">Prometheus</li><li class=\"mb-0\">Terraform</li><li class=\"mb-0\">Bash / Linux</li></ul></div></div></div><div class=\"col-span-4 sm:col-span-9\"><div class=\"bg-white shadow-xl rounded-lg p-6\"><h2 class=\"text-xl font-bold mt-6 mb-4\">curriculum vitae</h2><div class=\"mb-6\"><div class=\"flex justify-between flex-wrap gap-2 w-full\"><span class=\"text-blue-700 font-bold\">senior backend engineer</span><p><span><img src=\"/static/images/formlabs_logo.svg\" style=\"padding-bottom: 0.15rem;\" class=\"w-24 h-10 inline-block\" alt=\"Formlabs Logo\"></span> <span class=\"pl-4 text-blue-700\">2023 - &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span></p></div><p class=\"mt-2 text-left\">I'm currently working @Formlabs in Fleet Control Team, working on a Dashboard for Formlabs 3D printers.<br>For the most part I'm developing the synchonisation of data from printers to backend, so we can build cool feature's for the dashboard.<ul style=\"margin-left: 20px;\" class=\"mt-2 text-left list-disc\"><li>I'm owning and developing the Django backend of the Dashboard</li><li>Also developing a Go client application that is synchronizing data from printers to the backend</li><li>Transitioned the backend from raw kubernetes manifests to Helm charts</li><li>Implemented monitoring in Prometheus for the backend microsevice architecture</li><li>Resolved many scalability issues of the backend</li></ul></p></div><div class=\"mb-6\"><div class=\"flex justify-between flex-wrap gap-2 w-full\"><span class=\"text-blue-700 font-bold\">senior backend engineer</span><p><span><img src=\"/static/images/tier_logo.svg\" style=\"padding-bottom: 0.15rem;\" class=\"w-14 h-10 inline-block\" alt=\"TIER Logo\"></span> <span class=\"pl-4 text-blue-700\">2021 - 2023</span></p></div><p class=\"mt-2 text-left\">In the backend Map team, I've redisigned the poll based architecture of TIERs backend API to a full-duplex websocket based communication. This allowed showing vehicle data in real-time.</p></div><div class=\"mb-6\"><div class=\"flex justify-between flex-wrap gap-2 w-full\"><span class=\"text-blue-700 font-bold\">infrastucture engineer</span><p><span><img src=\"/static/images/prezi_logo.svg\" style=\"padding-bottom: 0.15rem;\" class=\"w-10 h-10 inline-block\" alt=\"Prezi Logo\"></span> <span class=\"pl-4 text-blue-700\">2020 - 2021</span></p></div><p class=\"mt-2 text-left\">I switched to the infrastucture team and worked on making sure Prezi's microsevice architecture is reliable.<ul style=\"margin-left: 20px;\" class=\"mt-2 text-left list-disc\"><li>Worked on improving observability through developing metrics and alerting for service SLOs</li><li>Developed Helm charts and Terraform modules</li><li>Worked on reducing AWS costs of Prezi</li></ul></p></div><div class=\"mb-6\"><div class=\"flex justify-between flex-wrap gap-2 w-full\"><span class=\"text-blue-700 font-bold\">full stack engineer</span><p><span><img src=\"/static/images/prezi_logo.svg\" style=\"padding-bottom: 0.15rem;\" class=\"w-10 h-10 inline-block\" alt=\"Prezi Logo\"></span> <span class=\"pl-4 text-blue-700\">2018 - 2020</span></p></div><p class=\"mt-2 text-left\">I was in the Payment and Licensing team, developing integrations with payment providers such as Zuora, PayPal and Braintree. As a full-stack developer I've been developing the Python / Django based backend with a Typescript / React frontend.<ul style=\"margin-left: 20px;\" class=\"mt-2 text-left list-disc\"><li>Led migration of payments and subscriptions for hundreds of thousands of users between major service providers, ensuring a seamless transition.</li><li>Integrated with 3rd party APIs like Zuora, AllPago, Braintree, Kount, Paypal, and Salesforce.</li><li>Developed internal RESTful APIs for efficient data retrieval and manipulation.</li></ul></p></div><div class=\"mb-6\"><div class=\"flex justify-between flex-wrap gap-2 w-full\"><span class=\"text-blue-700 font-bold\">linux sysadmin / automation engineer</span><p><span><img src=\"/static/images/ibm_logo.svg\" style=\"padding-bottom: 0.15rem;\" class=\"w-10 h-10 inline-block\" alt=\"IBM Logo\"></span> <span class=\"pl-4 text-blue-700 w-7 h-7\">2015 - 2018</span></p></div><p class=\"mt-2 text-left\">I've been maintaining IBM hosted servers of big airlines and other important customers.<br>This involved a lost of shell scripting for monitoring and configuration automations in Bash and Perl mostly.</p></div></div></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
