/**
 * @author zhangyuehao
 * @date 2019-02-18 18:49
 */

package goserbench

import (
	"math/rand"
	"time"
)

func generateGoProto() []*GoUser {
	a := make([]*GoUser, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &GoUser{
			Id:       randString(32),
			Name:     randString(16),
			Password: randString(18),
			Age:      rand.Int31n(5),
			BirthDay: time.Now().UnixNano(),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
		})
	}
	return a
}
