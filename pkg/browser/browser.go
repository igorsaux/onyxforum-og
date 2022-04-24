package browser

import (
	"context"
	"sync"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func HtmlToImage(html string, targetEl string) []byte {
	allocatorCtx, cancel := chromedp.NewRemoteAllocator(context.Background(), "ws://chrome:9222")

	defer cancel()

	ctx, cancel := chromedp.NewContext(allocatorCtx)

	defer cancel()

	picture := []byte{}

	err := chromedp.Run(ctx, chromedp.Navigate("about:blank"), chromedp.ActionFunc(func(ctx context.Context) error {
		loadingCtx, cancel := context.WithCancel(ctx)

		defer cancel()

		var wg sync.WaitGroup
		wg.Add(1)

		chromedp.ListenTarget(loadingCtx, func(ev interface{}) {
			if _, ok := ev.(*page.EventLoadEventFired); ok {
				cancel()
				wg.Done()
			}
		})

		frameTree, err := page.GetFrameTree().Do(ctx)

		if err != nil {
			return err
		}

		page.SetDocumentContent(frameTree.Frame.ID, html).Do(ctx)

		wg.Wait()

		return nil
	}), chromedp.FullScreenshot(&picture, 100))

	if err != nil {
		panic(err)
	}

	return picture
}
