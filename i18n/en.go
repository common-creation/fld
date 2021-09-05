package i18n

var en = Dictionary{
	"command": Dictionary{
		"error": "Command not found: {{.command}}. Please install and try again.",
	},
	"deploy": Dictionary{
		"help": Dictionary{
			"short": "Deploy the Line Smart City.",
		},
		"common": Dictionary{
			"ok":          "OK",
			"error":       "error",
			"builderror":  "build error",
			"deployerror": "deploy error",
			"execerror":   "sub-precess execution error",
			"codeerror":   "exit code error",
		},
		"summary": Dictionary{
			"build":  "Build result",
			"deploy": "Deploy result",
		},
		"admin-web": Dictionary{
			"usage": "Deploy 'admin-web' module.",
		},
		"all": Dictionary{
			"usage": "Deploy all module.",
		},
		"bi": Dictionary{
			"usage": "Deploy 'bi' module.",
		},
		"bosai": Dictionary{
			"usage": "Deploy 'bosai' module.",
		},
		"distribution": Dictionary{
			"usage": "Deploy 'distribution' module.",
		},
		"liff": Dictionary{
			"usage": "Deploy 'liff' module.",
		},
		"platform": Dictionary{
			"usage": "Deploy 'platform' module.",
		},
		"scenario": Dictionary{
			"usage": "Deploy 'scenario' module.",
		},
		"sequential": Dictionary{
			"usage": "Disable concurrent build.",
		},
		"survey": Dictionary{
			"usage": "Deploy 'survey' module.",
		},
		"useContainer": Dictionary{
			"usage": "Use docker to build.",
		},
	},
}
