package website_parse

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"os"
)

func Screenshot(url string) error {
	// create options for custom path to Google Chrome binary
	opts := chromedp.DefaultExecAllocatorOptions[:]

	// create context with custom options
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create context
	ctx, cancel = chromedp.NewContext(
		ctx,
	)
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx, chromedp.EmulateViewport(1920, 1080), chromedp.Navigate(url), chromedp.WaitReady(`body`), chromedp.CaptureScreenshot(&buf)); err != nil {
		return err
	}

	if err := os.WriteFile("./storage/websites/"+url[8:]+".png", buf, 0644); err != nil {
		return err
	}

	return nil
}

func GetHTML(url string) (string, error) {
	// create options for custom path to Google Chrome binary
	opts := chromedp.DefaultExecAllocatorOptions[:]

	// create context with custom options
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create context
	ctx, cancel = chromedp.NewContext(
		ctx,
		chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// Check if the page is loaded
	if err := chromedp.Run(ctx, chromedp.Navigate(url), chromedp.WaitReady(`body`)); err != nil {
		return "", err
	}

	// Extract HTML
	var htmlContent string
	if err := chromedp.Run(ctx, chromedp.InnerHTML(`html`, &htmlContent)); err != nil {
		return "", err
	}

	return htmlContent, nil
}

func GetLocations(url string) ([]string, error) {
	fmt.Println("Getting locations for", url)
	temp := []string{
		"Houston, TX",
		"Austin, TX",
	}
	return temp, nil
}

//
//func getSitemap(rootURL, rootPath string, fetchLimit int) ([]byte, error) {
//	parser := gowebcrawler.UrlParser{}
//
//	crawler := gowebcrawler.WebCrawler{
//		Parser:     &parser,
//		RootUrl:    rootURL,
//		FetchLimit: fetchLimit,
//	}
//
//	json, err := crawler.Crawl(rootPath)
//	if err != nil {
//		return nil, err
//	}
//
//	return json, nil
//}
