/**
 * @author zhangyuehao
 * @date 2019-02-19 13:03
 */

package goserbench

import (
	"math/rand"
	"testing"
	"time"
)

func generateGenCodeProto() []*GencodeUser {
	a := make([]*GencodeUser, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &GencodeUser{
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

func BenchmarkMarshalGencodeUser(b *testing.B) {
	b.StopTimer()
	data := generateGenCodeProto()
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		data[rand.Intn(len(data))].Marshal(nil)
	}
}

func BenchmarkUnmarshalGencodeUser(b *testing.B) {
	validate := true
	b.StopTimer()
	data := generateGenCodeProto()
	ser := make([][]byte, len(data))
	for i, d := range data {
		ser[i], _ = d.Marshal(nil)
	}
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &GencodeUser{}
		_, err := o.Unmarshal(ser[n])
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
		// Validate unmarshalled data.
		if validate != false {
			d := data[n]
			correct := o.Id == d.Id && o.Name == d.Name && o.Password == d.Password && o.Age == d.Age && o.Spouse == d.Spouse && o.Money == d.Money && o.BirthDay == d.BirthDay //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}
