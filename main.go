package main

import (
	"jiuzhua/lib"
	"time"
)

func main() {
	lib.StartWindow("golang html", 400, 400, true, func() {
		pager := lib.NewPager(&MyPage{}, &MyPageAction{})
		pager.StartPage(0)
	})

}

type MyPage struct {
	lib.PageImpl
}
type MyPage1 struct {
	lib.PageImpl
}
type MyPageAction struct {
	P *MyPage
	i int
}

func (t *MyPageAction) Add(i int, s *string) error {
	pager := lib.NewPager(&MyPage1{}, nil)
	pager.StartPage(0)
	*s = "success"
	return nil

}
func (p *MyPage) Stop() {

}
func (p *MyPage) Start() {
	p.SetContentView("html/page1.html")

}
func (p *MyPage1) Stop() {

}
func (p *MyPage1) Start() {
	p.SetContentView("html/page2.html")
	go func() {
		time.Sleep(5 * time.Second)
		p.Window.Backup()
	}()

}
