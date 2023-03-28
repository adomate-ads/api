package website_parse

import (
	"github.com/go-rod/rod/lib/proto"
	"github.com/ysmood/gson"
	"os"
	"time"
)

func CaptureWebsite(url string) error {
	page := browser.MustPage(url)
	page.MustEvalOnNewDocument(`
		() => {
			return new Promise(resolve => {
				window.addEventListener('load', resolve);
			});
		}
	`)

	err := page.WaitLoad()
	if err != nil {
		return err
	}

	err = page.WaitIdle(2 * time.Second)
	if err != nil {
		return err
	}

	img, err := page.Screenshot(true, &proto.PageCaptureScreenshot{
		Format:  proto.PageCaptureScreenshotFormatWebp,
		Quality: gson.Int(90),
		Clip: &proto.PageViewport{
			X:      0,
			Y:      0,
			Width:  1920,
			Height: 1080,
			Scale:  1,
		},
		FromSurface: true,
	})
	if err != nil {
		return err
	}

	err = os.WriteFile("./storage/websites/"+url[8:]+".webp", img, 0644)
	if err != nil {
		return err
	}

	return nil
}
