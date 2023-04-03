package main

import (
	"context"
	"log"
	"os"

	"github.com/cgenuity/gowebcrawler"
	"github.com/chromedp/chromedp"
)

func getHTML(url string) (string, error) {
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
func takeScreenshot(url string, width, height int) ([]byte, error) {
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

	var buf []byte
	if err := chromedp.Run(ctx, chromedp.EmulateViewport(1920, 1080), chromedp.Navigate(url), chromedp.WaitReady(`body`), chromedp.CaptureScreenshot(&buf)); err != nil {
		return nil, err
	}

	return buf, nil
}
func getSitemap(rootURL, rootPath string, fetchLimit int) ([]byte, error) {
	parser := gowebcrawler.UrlParser{}

	crawler := gowebcrawler.WebCrawler{
		Parser:     &parser,
		RootUrl:    rootURL,
		FetchLimit: fetchLimit,
	}

	json, err := crawler.Crawl(rootPath)
	if err != nil {
		return nil, err
	}

	return json, nil
}
func main() {
	// Get HTML
	url := ""
	htmlContent, err := getHTML(url)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("page.html", []byte(htmlContent), 0o644); err != nil {
		log.Fatal(err)
	}

	// Take Screenshot
	screenshot, err := takeScreenshot(url, 1920, 1080)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(url, screenshot, 0o644); err != nil {
		log.Fatal(err)
	}

	// Get Sitemap
	sitemap, err := getSitemap(url, "/", 50)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("sitemap.json", sitemap, 0o644); err != nil {
		log.Fatal(err)
	}
}
