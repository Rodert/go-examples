package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/tebeka/selenium"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/htmlindex"
)

func detectContentCharset(body io.Reader) string {
	r := bufio.NewReader(body)
	if data, err := r.Peek(1024); err == nil {
		if _, name, _ := charset.DetermineEncoding(data, ""); len(name) != 0 {
			return name
		}
	}

	return "utf-8"
}

/* 自动识别编码格式 */
func DecodeHTMLBody(body io.Reader, charset string) (io.Reader, error) {
	if charset == "" {
		charset = detectContentCharset(body)
	}

	e, err := htmlindex.Get(charset)
	if err != nil {
		return nil, err
	}

	if name, _ := htmlindex.Name(e); name != "utf-8" {
		body = e.NewDecoder().Reader(body)
	}

	return body, nil
}

/* baidu 案例 */
func BaiduHotSearch() {
	res, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	utf8Body, _ := DecodeHTMLBody(res.Body, "")
	doc, err := goquery.NewDocumentFromReader(utf8Body)
	// doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".s-hotsearch-content .hotsearch-item").Each(func(i int, s *goquery.Selection) {
		content := s.Find(".title-content-title").Text()
		fmt.Printf("%d: %s\n", i, content)
	})
}

/*
	curl 'https://coinmarketcap.com/currencies/tether/' \
	 --proxy http://127.0.0.1:4780 \
	 -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36' \
	 --compressed
*/
func CoinmarketcapInfo() {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()

	uri, _ := url.Parse("http://127.0.0.1:4780")
	client := &http.Client{
		Timeout: 20 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(uri),
		},
	}
	url := "https://coinmarketcap.com/currencies/tether/"
	url = "https://coinmarketcap.com/currencies/dogecoin/"
	reqest, _ := http.NewRequest("GET", url, nil)
	reqest.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")

	res, err := client.Do(reqest)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	utf8Body, _ := DecodeHTMLBody(res.Body, "")
	doc, err := goquery.NewDocumentFromReader(utf8Body)
	// doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".dropdownItem").Each(func(i int, s *goquery.Selection) {
		imgDom := s.Find("img")
		val, exists := imgDom.Attr("src")
		if exists {
			fmt.Println(val)
		}

	})

	fmt.Println("123")
}

func SeleniumTest() {

	const (
		seleniumPath     = `./chromedriver`
		chromeDriverPath = `./chromedriver`
		// port         = 9515
		port = 8080
	)

	// Start a WebDriver server instance
	opts := []selenium.ServiceOption{
		selenium.Output(os.Stderr), // Output debug information to STDERR.
	}
	selenium.SetDebug(true)
	service, err := selenium.NewChromeDriverService(chromeDriverPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get("http://play.golang.org/?simple=1"); err != nil {
		panic(err)
	}

	// Get a reference to the text box containing code.
	elem, err := wd.FindElement(selenium.ByCSSSelector, "#code")
	if err != nil {
		panic(err)
	}
	// Remove the boilerplate code already in the text box.
	if err := elem.Clear(); err != nil {
		panic(err)
	}

	// Enter some new code in text box.
	err = elem.SendKeys(`
		package main
		import "fmt"
		func main() {
			fmt.Println("Hello WebDriver!")
		}
	`)
	if err != nil {
		panic(err)
	}

	// Click the run button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, "#run")
	if err != nil {
		panic(err)
	}
	if err := btn.Click(); err != nil {
		panic(err)
	}

	// Wait for the program to finish running and get the output.
	outputDiv, err := wd.FindElement(selenium.ByCSSSelector, "#output")
	if err != nil {
		panic(err)
	}

	var output string
	for {
		output, err = outputDiv.Text()
		if err != nil {
			panic(err)
		}
		if output != "Waiting for remote server..." {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("%s", strings.Replace(output, "\n\n", "\n", -1))

	// Example Output:
	// Hello WebDriver!
	//
	// Program exited.

	return
}

func main() {
	// BaiduHotSearch()
	CoinmarketcapInfo()
	SeleniumTest()
}
