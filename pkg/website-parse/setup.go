package website_parse

import "github.com/go-rod/rod"

var browser *rod.Browser

func Setup() error {
	browser = rod.New()
	err := browser.Connect()
	if err != nil {
		return err
	}
	return nil
}
