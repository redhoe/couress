package lang

import (
	"embed"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// If you use go:embed
//
//go:embed *.toml
var languageFs embed.FS
var bundle *i18n.Bundle = nil

func GetBundle() *i18n.Bundle {
	if bundle != nil {
		return bundle
	}
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFileFS(languageFs, "locale.zh-Hans.toml")
	return bundle
}
