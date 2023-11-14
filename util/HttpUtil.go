package util

import (
	"net"
	"net/http"
	"sync"
	"time"
	//"testing"
)

var (
	client *http.Client
	once   sync.Once
)

func GetHttpClient() *http.Client {
	once.Do(func() {
		if client == nil {
			newHttpClient := http.Client{
				// 总体的超时设置为10秒，需要注意，如果超时并不代表该处理失败，
				// 只代表该处理在10秒内未完成，处理结果未知
				Timeout: 15 * time.Second,
				Transport: &http.Transport{
					// 指定dial的超时设置
					DialContext: (&net.Dialer{
						Timeout:   10 * time.Second,
						KeepAlive: 30 * time.Second,
					}).DialContext,
					MaxIdleConns:          50,
					IdleConnTimeout:       60 * time.Second,
					TLSHandshakeTimeout:   5 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
					// 限制响应头的大小，避免依赖的服务过多使用响应头
					MaxResponseHeaderBytes: 5 * 1024,
				},
			}
			client = &newHttpClient
		}
	})
	return client
}

/*func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go GetHttpClient()
	}
}*/
