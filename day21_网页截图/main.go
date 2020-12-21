package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

//GetScreenShot 得到网页快照
func GetScreenShot(url string, filename string, width int, height int) bool {
	ctx, cancel := chromedp.NewContext(context.Background()) //ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	defer cancel()

	if err := chromedp.Run(ctx, ScreenshotTasks(url, filename, float64(width), float64(height))); err != nil {
		log.Fatal(err)
	}
	return true
}

//ScreenshotTasks 网页快照任务
func ScreenshotTasks(url string, filename string, wdith float64, height float64) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		// chromedp.ActionFunc(func(ctx context.Context) (err error) {
		// 	*imageBuf, err = page.CaptureScreenshot().WithQuality(90).Do(ctx)
		// 	return err
		// }),

		chromedp.ActionFunc(func(ctxt context.Context) error {
			_, _, contentRect, err := page.GetLayoutMetrics().Do(ctxt)
			if err != nil {
				return err
			}
			if wdith == 0 {
				wdith = contentRect.Width
			}
			if height == 0 {
				height = contentRect.Height
			}
			v := page.Viewport{
				X:      contentRect.X,
				Y:      contentRect.Y,
				Width:  wdith,
				Height: height,
				Scale:  1,
			}
			log.Printf("Capture %#v", v)
			buf, err := page.CaptureScreenshot().WithClip(&v).Do(ctxt)
			if err != nil {
				return err
			}
			log.Printf("Write %v", filename)
			return ioutil.WriteFile(filename, buf, 0644)
		}),
	}
}

func main() {
	/*
		#build on mac
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o  dboopScreenLinux  main.go
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -o  dboopScreenWin.exe  main.go
		CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o  dboopScreenMac  main.go


		cp dboopScreen*  /Users/kouko/Desktop/Qdesk/share/dbooptool
		cd /Users/kouko/Desktop/Qdesk/share/dbooptool
		git add .
		git commit -m "dboopscreen"
		git push

		cd  /Users/kouko/go/src/github.com/51ak/golearn/day21_网页截图
	*/
	var url string
	var filename string
	var width int
	var height int
	flag.StringVar(&url, "url", "https://www.baidu.com/", "url")
	flag.StringVar(&filename, "filename", "baidu.png", "文件存放路径和文件名，不指定路径放在当前目录")
	flag.IntVar(&width, "width", 0, "0：的时候是自动设置，单位是px")
	flag.IntVar(&height, "height", 0, "0：的时候是自动设置，单位是px")
	flag.Parse()
	GetScreenShot(url, filename, width, height)
}
