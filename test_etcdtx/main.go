/**
 * @author zhangyuehao
 * @date 2019-03-08 14:36
 */

package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/eechains/code_test/test_etcdtx/etcd"
	"github.com/eechains/code_test/test_etcdtx/logger"
	"github.com/etcd-io/etcd/pkg/transport"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
	"os"
	"path/filepath"
	"strings"
)

var (
	log            = logger.InitLogger()
	client         *etcd.Client
	defaultTLSPath = "/opt/eechains/scaring/assets"
)

func GetCertPath() string {
	path := os.Getenv("GOPATH")
	if path == "" {
		return defaultTLSPath
	}
	s := strings.Split(path, ":")
	l := len(s)
	if l == 0 {
		return ""
	}
	path = filepath.Join(s[l-1:][0], "src", "github.com", "eechains", "scaring", "assets")
	return path
}

func newEtcdClient() *etcd.Client {
	config := &etcd.ClientConfig{
		Endpoints:   []string{"116.62.118.133:2379"},
		TLSFunc:     buildTLSConfig,
		DialTimeOut: etcd.DIAL_TIMEOUT,
	}
	cli, err := etcd.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("client created")

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

func NewMutex(prefixKey string) (*concurrency.Mutex, error) {
	return client.NewMutex(prefixKey)
}

func Del(key string) error {
	return client.DeleteWithTimeOut(key)
}

func init() {
	client = newEtcdClient()
}

func main() {
	kvtx := clientv3.NewKV(client.Client)
	ctx, cancel := context.WithTimeout(context.Background(), client.Config.RequestTimeOut)
	_, err := kvtx.Txn(ctx).
		Then(clientv3.OpPut("txTest", "JIANGJING"), clientv3.OpPut("txTest1", "XIAOLAN")).
		Commit()
	cancel()
	if err != nil {
		log.Fatal(err)
	}

	v, err := client.Get("txTest")
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("txText is %v", v)
	v, err = client.Get("txTest1")
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("txText1 is %v", v)
}
