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
		"admin-web": Dictionary{
			"usage": "Web管理画面 をデプロイ対象に加えます",
		},
		"all": Dictionary{
			"usage": "全てのモジュールをデプロイ対象に加えます",
		},
		"bi": Dictionary{
			"usage": "統計モジュール をデプロイ対象に加えます",
		},
		"bosai": Dictionary{
			"usage": "防災モジュール をデプロイ対象に加えます",
		},
		"distribution": Dictionary{
			"usage": "配信モジュール をデプロイ対象に加えます",
		},
		"liff": Dictionary{
			"usage": "LIFFモジュール をデプロイ対象に加えます",
		},
		"platform": Dictionary{
			"usage": "汎用API をデプロイ対象に加えます",
		},
		"scenario": Dictionary{
			"usage": "シナリオモジュール をデプロイ対象に加えます",
		},
		"sequential": Dictionary{
			"usage": "モジュールのビルドを逐次で行います",
		},
		"survey": Dictionary{
			"usage": "帳票モジュール をデプロイ対象に加えます",
		},
		"useContainer": Dictionary{
			"usage": "ビルドにDockerを使用します",
		},
	},
}
