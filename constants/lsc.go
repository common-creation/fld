package constants

type (
	LSCCommand struct {
		Build string
		Deploy string
		UseContainer *string
	}
)

var (
	LSC_COMMANDS = map[string]LSCCommand{
		"admin-web": {
			Build:        "admin-web build",
			Deploy:       "admin-web deploy",
			UseContainer: nil,
		},
		"platform": {
			Build:        "platform sam build",
			Deploy:       "platform sam deploy",
			UseContainer: strPtr("--useContainer"),
		},
		"scenario": {
			Build:        "scenario sam build",
			Deploy:       "scenario sam deploy",
			UseContainer: strPtr("--useContainer"),
		},
		"bosai": {
			Build:        "bosai sam build",
			Deploy:       "bosai sam deploy",
			UseContainer: strPtr("--useContainer"),
		},
		"bi": {
			Build:        "bi sam build",
			Deploy:       "bi sam deploy",
			UseContainer: strPtr("--useContainer"),
		},
		"survey": {
			Build:        "survey sam build",
			Deploy:       "survey sam deploy",
			UseContainer: strPtr("--useContainer"),
		},
		"liff": {
			Build:        "survey liff build",
			Deploy:       "survey liff deploy",
			UseContainer: strPtr("--useContainer"),
		},
		"distribution": {
			Build:        "distribution sam build",
			Deploy:       "distribution sam deploy",
			UseContainer: strPtr("--useContainer"),
		},
	}
)

func strPtr(str string) *string {
	return &str
}
