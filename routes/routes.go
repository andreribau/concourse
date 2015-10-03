package routes

import "github.com/tedsuo/rata"

const (
	Index           = "Index"
	Pipeline        = "Pipeline"
	TriggerBuild    = "TriggerBuild"
	GetBuild        = "GetBuild"
	GetBuilds       = "GetBuilds"
	GetJoblessBuild = "GetJoblessBuild"
	Public          = "Public"
	GetResource     = "GetResource"
	GetJob          = "GetJob"
	LogIn           = "LogIn"
)

var Routes = rata.Routes{
	// public
	{Path: "/", Method: "GET", Name: Index},
	{Path: "/pipelines/:pipeline_name", Method: "GET", Name: Pipeline},
	{Path: "/pipelines/:pipeline_name/jobs/:job", Method: "GET", Name: GetJob},
	{Path: "/pipelines/:pipeline_name/resources/:resource", Method: "GET", Name: GetResource},
	{Path: "/public/:filename", Method: "GET", Name: Public},
	{Path: "/public/fonts/:filename", Method: "GET", Name: Public},
	{Path: "/public/favicons/:filename", Method: "GET", Name: Public},

	// public jobs only
	{Path: "/pipelines/:pipeline_name/jobs/:job/builds/:build", Method: "GET", Name: GetBuild},

	// private
	{Path: "/login", Method: "GET", Name: LogIn},
	{Path: "/pipelines/:pipeline_name/jobs/:job/builds", Method: "POST", Name: TriggerBuild},
	{Path: "/builds", Method: "GET", Name: GetBuilds},
	{Path: "/builds/:build_id", Method: "GET", Name: GetJoblessBuild},
}
