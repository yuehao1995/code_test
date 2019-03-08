/**
 * Created by martin on 18/02/2019
 */

package etcd

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/pkg/transport"

	"github.com/eechains/scaring/pkg/core/notify"
)

var (
	tlsPath     = GetCertPath()
	testClient  = newTestClient()
	emailNotify = NewEmailNofify()
)

func newTestClient() *Client {
	config := &ClientConfig{
		Endpoints:   []string{"116.62.118.133:2379"},
		TLSFunc:     buildTLSConfig,
		DialTimeOut: DIAL_TIMEOUT,
	}
	cli, err := NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	return cli
}

func TestApi(t *testing.T) {
	err := testClient.Put("/test/key", "/test/value")
	err = testClient.Put("/test/key2", "/test/value")

	if err != nil {
		t.Fatal(err)
	}

	v, err := testClient.Get("/test", clientv3.WithPrefix())
	if err != nil {
		t.Log(err)
	}

	t.Log(v)
}

func TestClient_Watch(t *testing.T) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go testClient.Watch(ctx, "/test/key", emailNotify)

	// 三次put操作
	go func() {
		count := 0
		for {
			select {
			case <-time.After(3 * time.Second):
				t.Log()
				if count > 2 {
					wg.Done()
				}
				log.Printf("执行PUT操作-写入%s\n", time.Now().String())
				testClient.Put("/test/key", strconv.Itoa(count))
				count++

			}
		}

	}()
	time.AfterFunc(15*time.Second, func() {
		// 测试cancel 方法
		cancel()
	})
	wg.Wait()
}

// 代理模式
type testEmail struct {
	*notify.Email
}

func NewEmailNofify() WatchDo {

	s := "dev@ee-chain.com"
	r := []string{"yuehao@ee-chain.com", "jiangjin@ee-chain.com"}
	config := &notify.SMTPConfig{
		AuthKey: "zV01vHa373cox0Kxdb0MKtZJj1I8EU",
		Host:    "smtp.exmail.qq.com:587",
	}

	e := notify.NewEmail(config)
	email := e.Sender(s, r)

	testE := &testEmail{
		Email: email,
	}

	return testE
}

func (t *testEmail) DoFunc(wp clientv3.WatchResponse) error {

	fmt.Println()
	e, err := t.SetBody("测试邮件", []byte("接受到事件"))
	if err != nil {
		log.Println(err)
		return err
	}
	e.Send()
	return nil
}

func GetCertPath() string {
	path := os.Getenv("GOPATH")
	s := strings.Split(path, ":")
	l := len(s)
	if l == 0 {
		return ""
	}
	path = filepath.Join(s[l-1:][0], "src", "github.com", "eechains", "scaring", "asset")
	return path
}

func GetEtcdClient() *Client {
	config := &ClientConfig{
		Endpoints:   []string{"116.62.118.133:2379"},
		TLSFunc:     buildTLSConfig,
		DialTimeOut: DIAL_TIMEOUT,
	}
	cli, err := NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	return cli
}

func buildTLSConfig() (*tls.Config, error) {
	certPath := GetCertPath()
	tlsInfo := transport.TLSInfo{
		CertFile:      certPath + "/tls/client.pem",
		KeyFile:       certPath + "/tls/client-key.pem",
		TrustedCAFile: certPath + "/tls/ca.pem",
	}
	return tlsInfo.ClientConfig()
}
