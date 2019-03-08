/**
 * @author zhangyuehao
 * @date 2019-02-18 18:02
 */

package goserbench

import (
	"math/rand"
	"time"
)

func generateGogoProto() []*GogoUser {
	a := make([]*GogoUser, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &GogoUser{
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
