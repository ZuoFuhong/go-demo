package http_proxy

import (
	"testing"
)

func Test_RunProxyServer(t *testing.T) {
	runProxyServer()
}

func Test_ProxyRequest(t *testing.T) {
	proxyRequest()
}

func Test_PureProxyRequest(t *testing.T) {
	pureProxyRequest()
}
