package main

import (
	"context"
	"flag"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"log"
)

// go build -ldflags="-H windowsgui" -o ../gcyp.exe ./crawler/app1.sfda.gov.cn/GCYP
func main() {
	flag.Parse()

	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	err := chromedp.Run(ctx, visible(urlTarget))
	if err != nil {
		log.Fatal(err)
	}
}

func visible(host string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(host),
		chromedp.ActionFunc(func(context.Context) error {
			log.Printf("waiting for " + host)
			return nil
		}),
		chromedp.WaitReady(`body`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			v, exp, err := runtime.Evaluate(`document.body.innerHtml`).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			log.Println(v)
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
			log.Printf(">>>>>>>>>>>>>>>>>>>> injectCommitForECMA OK")
			return nil
		}),
	}
}
