/**
 * @author zhangyuehao
 * @date 2019-02-18 18:43
 */

package goserbench

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"math/rand"
	"testing"
)

func BenchmarkGoprotobufMarshal(b *testing.B) {
	b.StopTimer()
	data := generateGoProto()
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		proto.Marshal(data[rand.Intn(len(data))])
	}
}

func BenchmarkGoprotobufUnmarshal(b *testing.B) {
	validate := true
	b.StopTimer()
	data := generateGoProto()
	ser := make([][]byte, len(data))
	for i, d := range data {
		ser[i], _ = proto.Marshal(d)
	}
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &GoUser{}
		t := string(ser[n])
		fmt.Println(t)
		err := proto.Unmarshal([]byte(t), o)
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
