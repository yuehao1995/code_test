/**
 * Created by martin on 20/02/2019
 */

package etcd

import (
	"go.etcd.io/etcd/clientv3"
)

// WatchDo 接口定义watch通知接口规范
type WatchDo interface {
	// DoFunc 通知函数接受clientv3.WatchResponse参数函数具体逻辑请自己实现
	DoFunc(clientv3.WatchResponse) error
}
