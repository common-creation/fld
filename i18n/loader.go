package i18n

import (
	"github.com/mohemohe/temple"
	"os"
	"strings"
)

type (
	LANG       int
	Dictionary map[string]interface{}
)

const (
	Unknown LANG = iota + 1
	Japanese
)

var emptyReplacer = map[string]interface{}{}

func currentLang() LANG {
	lang := os.Getenv("LANG")

	switch {
	case strings.HasPrefix(lang, "ja"):
		return Japanese
	default:
		return Unknown
	}
}

func template(base Dictionary, key string) string {
	paths := strings.Split(key, ".")
	template := base
	for i, v := range paths {
		if template[v] != nil {
			if i < len(paths)-1 {
				template = template[v].(Dictionary)
			} else {
				result := template[v].(string)
				if template[v] == nil || result == "" {
					return key
				}
				return result
			}
		} else {
			return key
		}
	}
	panic("i18n key error")
}

func T(key string, replacer map[string]interface{}) string {
	if replacer == nil {
		replacer = emptyReplacer
	}

	lang := currentLang()

	var base map[string]interface{}
	switch lang {
	case Japanese:
		base = ja
	default:
		// TODO: en
		base = ja
	}

	if result, err := temple.Execute(template(base, key), replacer); err != nil || len(result) == 0 {
		return key
	} else {
		return result
	}
}
