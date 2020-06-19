package sat

import (
	"testing"
)

func BenchmarkNewDict(b *testing.B) {
	DefaultDict(SetPath("/users/go/src/xxx.txt"))
	dicter := DefaultDict()
	for i := 0; i < b.N; i++ {
		_ = dicter.ReadReverse("么")
	}
}
