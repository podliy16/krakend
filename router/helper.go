package router

import (
	"github.com/podliy16/krakend/config"
	"github.com/podliy16/krakend/proxy"
	"net/http"
)

func IsValidSequentialEndpoint(endpoint *config.EndpointConfig) bool {
	if endpoint.ExtraConfig[proxy.Namespace] == nil {
		return false
	}

	proxyCfg := endpoint.ExtraConfig[proxy.Namespace].(map[string]interface{})
	if proxyCfg["sequential"] == false {
		return false
	}

	for i, backend := range endpoint.Backend {
		if backend.Method != http.MethodGet && (i+1) != len(endpoint.Backend) {
			return false
		}
	}

	return true
}
