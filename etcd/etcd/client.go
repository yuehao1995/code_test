/**
 * Created by martin on 02/02/2019
 */

package etcd

import (
	"context"
	"crypto/tls"
	"log"
	"sync"
	"time"

	"go.etcd.io/etcd/clientv3"
)

const (
	// 连接超时时间
	DIAL_TIMEOUT    = 5 * time.Second
	REQUEST_TIMEOUT = 3 * time.Second
)

var (
	client *Client
	once   = sync.Once{}
)

type ClientConfig struct {
	// 节点ip
	Endpoints []string
	// 证书配置生成函数
	TLSFunc TLSConfigFunc
	// 连接超时时间
	DialTimeOut time.Duration
	// 请求超时时间
	RequestTimeOut time.Duration
}

// client clientv3.Clinet对客户端进行了简单的封装
type Client struct {
	// etcd 连接实例
	client *clientv3.Client
	// 连接接配置
	Config *ClientConfig
	// watch 结构体的实例
	watch clientv3.Watcher
}

// 生成TLS 证书配置
type TLSConfigFunc func() (*tls.Config, error)

// NewClient 创建一个连接使用once.Do 保证只会创建一个实例
func NewClient(config *ClientConfig) (*Client, error) {
	var err error
	if config == nil {
		return nil, configError
	}

	once.Do(func() {
		client, err = newClient(config)
	})

	return client, err
}

func newClient(config *ClientConfig) (*Client, error) {
	if len(config.Endpoints) == 0 {
		return nil, endpointsEmptyError
	}

	if config.RequestTimeOut == 0 {
		config.RequestTimeOut = REQUEST_TIMEOUT
	}

	if config.DialTimeOut == 0 {
		config.DialTimeOut = DIAL_TIMEOUT
	}

	clientConfig := clientv3.Config{
		Endpoints:   config.Endpoints,
		DialTimeout: config.DialTimeOut,
	}

	if config.TLSFunc != nil {
		// 生成TLS配置
		tlsConfig, err := config.TLSFunc()
		if err != nil {
			return nil, err
		}

		clientConfig.TLS = tlsConfig
	}

	cli, err := clientv3.New(clientConfig)

	if err != nil {
		return nil, err
	}

	client := &Client{
		Config: config,
		client: cli,
	}

	client.client = cli
	return client, nil
}

// 关闭所有实例的连接
func (e *Client) CloseClient() {
	e.client.Close()
	e.watch.Close()
}

func (e *Client) get(context context.Context, key string, opts ...clientv3.OpOption) (values []string, err error) {
	resp, err := e.client.Get(context, key, opts...)
	if err != nil {
		return values, err
	}

	s := resp.Kvs
	length := len(s)

	for i := 0; i < length; i++ {
		values = append(values, string(s[i].Value))
	}

	return values, nil
}

func (e *Client) put(context context.Context, key, value string, opts ...clientv3.OpOption) error {
	if key == "" || value == "" {
		return argError
	}
	_, err := e.client.Put(context, key, value, opts...)
	return err
}

func (e *Client) Put(key, value string, opts ...clientv3.OpOption) error {
	return e.put(context.TODO(), key, value, opts...)
}

func (e *Client) Get(key string, opts ...clientv3.OpOption) ([]string, error) {
	return e.get(context.TODO(), key, opts...)
}

func (e *Client) PutWithTimeOut(key, value string, opts ...clientv3.OpOption) error {
	ctx, cancel := context.WithTimeout(context.Background(), e.Config.RequestTimeOut)
	defer cancel()
	err := e.put(ctx, key, value, opts...)
	return err
}

func (e *Client) GetWithTimeOut(key string, opts ...clientv3.OpOption) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), e.Config.RequestTimeOut)
	defer cancel()
	values, err := e.get(ctx, key, opts...)
	return values, err
}

// Watch 监听某个key的动作
func (e *Client) Watch(ctx context.Context, key string, do WatchDo, opts ...clientv3.OpOption) {
	if e.watch == nil {
		w := clientv3.NewWatcher(e.client)
		e.watch = w
	}
	watchChan := e.watch.Watch(ctx, key, opts...)
	for {
		select {
		case watchResponse := <-watchChan:
			//收到时间后将response 传递给实现给实现了WatchDo接口的struct去处理
			//由具体的DoFunc 去处理通知逻辑
			err := do.DoFunc(watchResponse)
			if err != nil {
				log.Printf("do error:%s", err)
			}
		case <-ctx.Done():
			// 退出Watch 监听goroutine
			return
		}
	}

}
