package ioc

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func LoadI18nBundle() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	//加载不同的语言文件
	bundle.LoadMessageFile("locales/en-US.json")
	bundle.LoadMessageFile("locales/zh-CN.json")
	return bundle
}
