package lib

import (
	"fmt"
	"reflect"
	"sync"
)

type Page interface {
	SetContentView(htmlPath string)

	getWindow() *Window
	getPageName() string
	setPageName(s string)
	getAction() interface{}
	setPageImpl(action interface{})

	Start()
	Stop()
	Close()
	CallFunc(s string)
}
type superPage struct {
	Window   *Window
	Data     map[string]interface{}
	Action   interface{}
	pageName string
}

type PageImpl struct {
	superPage
}

func (p *PageImpl) getWindow() *Window {
	return p.Window
}
func (p *PageImpl) getPageName() string {
	return p.pageName
}
func (p *PageImpl) getAction() interface{} {
	//
	return p.Action
}

func (p *PageImpl) setPageImpl(action interface{}) {

	p.superPage = superPage{
		Window: window,
		Action: action,
	}

}

func (p *PageImpl) SetContentView(htmlPath string) {

	pageMap["/"+p.pageName] = htmlPath

}
func (p *PageImpl) CallFunc(s string) {
	p.Window.CallFunc(s)
}
func (p *PageImpl) Start() {
	//
}
func (p *PageImpl) Stop() {

}
func (p *PageImpl) Close() {

}
func (p *PageImpl) setPageName(s string) {
	if pageMap["/"+s] != "" {
		//panic(errors.New("pageName exist"))
	}
	p.pageName = s
}

type Pager struct {
	lock sync.Mutex
	page Page
}

func (p *Pager) StartPage(c int) {
	lock.Lock()
	defer lock.Unlock()
	p.page.getWindow().addPage(p)
	if p.page.getAction()!=nil{
		rpcServer.Register(p.page.getAction())
	}
	if len(p.page.getWindow().pages)>2{
		pre:=p.page.getWindow().pages[len(p.page.getWindow().pages)-2]
		pre.StopPage()
		if c!=0{
			pre.page.Close()
			p.page.getWindow().pages = append(p.page.getWindow().pages[0:len(p.page.getWindow().pages)-2],p)
		}
	}
	p.page.Start()

	fmt.Println(p.page.getPageName())
	p.page.getWindow().open(p.page.getPageName())
}
func (p *Pager) StopPage() {
	p.page.Stop()
}
func NewPager(p Page, action interface{}) Pager {
	name := reflect.TypeOf(p).String()
	p.setPageImpl(action)
	p.setPageName(name)
	return Pager{lock:sync.Mutex{},page: p}
}
