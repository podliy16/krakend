package proxy

import (
	"context"
	"testing"

	"github.com/podliy16/krakend/config"
)

func BenchmarkNewRequestBuilderMiddleware(b *testing.B) {
	backend := config.Backend{
		URLPattern: "/supu",
		Method:     "GET",
	}
	proxy := NewRequestBuilderMiddleware(&backend)(dummyProxy(&Response{}))

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		proxy(context.Background(), &Request{})
	}
}
