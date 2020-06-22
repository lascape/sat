package sat

import (
	"testing"
)

func BenchmarkNewDict(b *testing.B) {
	dicter := DefaultDict()
	for i := 0; i < b.N; i++ {
		_ = dicter.ReadReverse("么")
	}
}