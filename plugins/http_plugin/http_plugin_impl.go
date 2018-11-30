package http_plugin

import (
	"github.com/eosspark/eos-go/common"
	"github.com/eosspark/eos-go/log"
	"github.com/eosspark/eos-go/plugins/appbase/asio"
	"github.com/eosspark/eos-go/plugins/http_plugin/rpc"
	"net"
	"net/http"
)

type HttpPluginImpl struct {
	UrlHandlers                   map[string]http.Handler
	AccessControlAllowOrigin      string
	AccessControlAllowHeaders     string
	AccessControlMaxAge           string
	AccessControlAllowCredentials bool //default false
	MaxBodySize                   common.SizeT
	httpsCeryChain                string
	httpsKey                      string

	listenStr           string
	ListenEndpoint      *http.ServeMux
	HttpsListenEndpoint *http.ServeMux

	Listener net.Listener
	Handlers *rpc.Server

	self *HttpPlugin
	log  log.Logger
}

func NewHttpPluginImpl(io *asio.IoContext) *HttpPluginImpl {
	impl := new(HttpPluginImpl)
	impl.UrlHandlers = make(map[string]http.Handler)
	impl.AccessControlAllowCredentials = false

	impl.log = log.New("HttpPlugin")
	impl.log.SetHandler(log.TerminalHandler)

	return impl
}
