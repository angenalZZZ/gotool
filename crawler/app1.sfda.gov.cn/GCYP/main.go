package main

import (
	"context"
	"flag"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

// go build -ldflags="-H windowsgui" -o ../gcyp.exe ./crawler/app1.sfda.gov.cn/GCYP
func main() {
	flag.Parse()

	// create context
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, time.Minute)
	defer cancel()

	// run task
	err := chromedp.Run(ctx, newTask(urlTarget, map[string]interface{}{
		`Referer`:    urlTarget,
		`User-Agent`: `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36`,
		`Accept`:     `text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3`,
		`Cookie`:     `JSESSIONID=FA86521B2220A08EAAA51B8BFBB8704F.7; neCYtZEjo8Gm80S=Old2Gpq5J24XsgCZc922FAoK5FmE9.asBmFGvTJz0OyPyr6kKD58wMk0jR3Uz7OZ; neCYtZEjo8Gm80T=5THCJT79xNXTm2isynZm_eHG4yN7.IPQbYDXxkenYQDRX0GlodjhisRK.sVuN88lmTtOE9B8MzEtnZ7jX_9P0JJXQcD9.2bLDULPn9gXflaI22gCAcQEkLrjlJDXhjxQSAFtuSq85ZMrVVaUFDiuxxzVcBs2R7Q3VT3aV6EF1wa2oF36PCW6Hk_o0gMC5NQ9yrg0v6YII0TlC9ySF7GYoIbE5krA0.mH58i0NDfwke8uWPhswujf7qCBKD7RMXXt1ihwtBdOKq63F1HCj2usuGN3IJIbGyWBoQrsZprJHYiOo7XVEo3pEkA4en33EtASOpBve7uAxSMZtMtixUgAhi8rG`,
	}))
	if err != nil {
		log.Fatal(err)
	}
}

func newTask(pageUrl string, headers map[string]interface{}) chromedp.Tasks {
	htmlBody := ""
	return chromedp.Tasks{
		//chromedp.Emulate(device.IPad),
		network.Enable(),
		network.SetExtraHTTPHeaders(headers),
		chromedp.Navigate(pageUrl),
		chromedp.ActionFunc(func(context.Context) error {
			log.Printf("Navigate url: " + pageUrl)
			return nil
		}),
		chromedp.WaitReady(`body`, chromedp.ByQuery),
		chromedp.OuterHTML(`body`, &htmlBody),
		chromedp.ActionFunc(func(context.Context) error {
			log.Println("Html body:", htmlBody)
			return nil
		}),
		chromedp.WaitVisible(`#select1`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(injectCommitForECMA).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.ActionFunc(func(context.Context) error {
			log.Printf(">>>>>>>>>>>>>>>>>>>> inject javascript OK")
			return nil
		}),
	}
}
