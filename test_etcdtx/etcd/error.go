/**
 * Created by martin on 19/02/2019
 */

package etcd

import "errors"

var (
	endpointsEmptyError = errors.New("the node list IP is empty")
	argError            = errors.New("illegal parameter")
	tlsError            = errors.New("invalid TLS certificate configuration")
	configError         = errors.New("invalid configuration")
)
