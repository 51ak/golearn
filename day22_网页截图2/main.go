package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/benbjohnson/phantomjs"
)

//https://github.com/benbjohnson/phantomjs
func GetScreenShot(url string, filename string, width int, height int) bool {
	defer func() {

		err := recover()
		if err != nil {
			println(err)
		}
	}()

	// Start the process once.
	if err := phantomjs.DefaultProcess.Open(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer phantomjs.DefaultProcess.Close()

	page, err := phantomjs.CreateWebPage()
	if err != nil {
		log.Fatal(err)
	}
	//set request headers
	requestHeader := http.Header{
		"User-Agent": []string{"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"},
	}

	if err := page.SetCustomHeaders(requestHeader); err != nil {
		log.Fatal(err)
	}

	// Setup the viewport and render the results view.
	if err := page.SetViewportSize(width, height); err != nil {
		log.Fatal(err)
	}
	// Open a URL.
	if err := page.Open(url); err != nil {
		log.Fatal(err)
	}
	if err := page.Render(filename, "png", 50); err != nil {
		log.Fatal(err)
	}
	return true
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
	flag.IntVar(&width, "width", 1024, "0：的时候是自动设置，单位是px")
	flag.IntVar(&height, "height", 768, "0：的时候是自动设置，单位是px")
	flag.Parse()
	fmt.Println(GetScreenShot(url, filename, width, height))
}
