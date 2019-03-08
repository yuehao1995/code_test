/**
 * @author zhangyuehao
 * @date 2019-02-18 18:50
 */

package goserbench

import (
	"fmt"
	"math/rand"
)

func randString(l int) string {
	buf := make([]byte, l)
	for i := 0; i < (l+1)/2; i++ {
		buf[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x", buf)[:l]
}
