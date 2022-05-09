package lib

import (
	"errors"
	"fmt"
	"sync"

	"github.com/webview/webview"
)

var lock = sync.RWMutex{}
var c = make(chan string, 1)

type Window struct {
	WebView webview.WebView
	pages   []*Pager
}

func StartWindow(title string, w, h int, resize bool, f func()) {
	if window != nil {
		panic(errors.New("window exist!"))
	}
	startRPCHandler()
	startHttpHandler()
	w1 := webview.New(false)
	defer w1.Destroy()
	w1.SetSize(800, 600, webview.HintNone)
	window = &Window{w1, make([]*Pager, 0)}

	go func() {
		for {
			js := <-c
			fmt.Println(js)
			window.WebView.Dispatch(func() {
				window.WebView.Eval(js)
			})
		}

	}()
	go f()
	window.WebView.Run()
	//window.WebView.Destroy()
	server.Shutdown(nil)
}

func (w *Window) open(s string) {
	c <- fmt.Sprintf(`window.open("http://localhost:8080/%s", "_self")`, s)
}
func (w *Window) CallFunc(s string) {
	c <- s
}
func (w *Window) Backup() {
	fmt.Println("back1", len(w.pages))
	lock.Lock()
	defer lock.Unlock()
	if len(w.pages) <= 1 {
		w.pages = w.pages[0:0]
		window.WebView.Terminate()
	} else {
		fmt.Println("back")
		pre := w.pages[len(w.pages)-1]

		pre.StopPage()
		pre.page.Close()

		p := w.pages[len(w.pages)-2]
		w.pages = w.pages[:len(w.pages)-2]

		p.page.getWindow().addPage(p)

		p.page.Start()

		fmt.Println(p.page.getPageName())
		p.page.getWindow().open(p.page.getPageName())

	}
}
func (w *Window) addPage(p *Pager) {
	w.pages = append(w.pages, p)
}
