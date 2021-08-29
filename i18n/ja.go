package i18n

var ja = Dictionary{
	"command": Dictionary{
		"error": "'{{.command}}' コマンドがありません。インストールしてから再度試してください",
	},
	"deploy": Dictionary{
		"help": Dictionary{
			"short": "LSCのデプロイを行います",
		},
		"common": Dictionary{
			"ok":          "OK",
			"error":       "エラー",
			"builderror":  "ビルドエラー",
			"deployerror": "デプロイエラー",
			"execerror":   "実行エラー",
			"codeerror":   "終了コードエラー",
		},
		"summary": Dictionary{
			"build":  "ビルド結果",
			"deploy": "デプロイ結果",
		},
	},
}
