package wechat_spider

import (
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

type spider struct {
	proxy *goproxy.ProxyHttpServer
}

var _spider = NewSpider()

func Regist(proc Processor) {
	_spider.Regist(proc)
}

func Run(port string) {
	_spider.Run(port)
}

func NewSpider() *spider {
	sp := &spider{}
	sp.proxy = goproxy.NewProxyHttpServer()
	sp.proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	return sp
}

func (s *spider) Regist(proc Processor) {
	s.proxy.OnResponse().DoFunc(ProxyHandle(proc))
}

func (s *spider) Run(port string) {
	log.Println("server will at port:" + port)
	log.Fatal(http.ListenAndServe(":"+port, s.proxy))
}
