package sat

import (
	"testing"
)

func BenchmarkNewDict(b *testing.B) {
	//err := InitDefaultDict(SetPath("/users/go/src/xxx.txt"))
	//if err != nil {
	//	b.Fatal(err)
	//}
	dicter := DefaultDict()
	for i := 0; i < b.N; i++ {
		_ = dicter.ReadReverse("么")
	}
}
