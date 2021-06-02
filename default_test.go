package sat

import (
	"testing"
)

func TestDefaultDict_Init(t *testing.T) {
	err := InitDefaultDict(SetPath("./word.txt"))
	if err != nil {
		t.Error(err)
		return
	}
	dicter := DefaultDict()
	t.Log(dicter.ReadReverse("一繁"))
	t.Log(dicter.Read("五"))
}

func BenchmarkDefaultDict_Read(b *testing.B) {
	dicter := DefaultDict()
	for i := 0; i < b.N; i++ {
		_ = dicter.Read("什麼")
	}
}

func BenchmarkDefaultDict_ReadReverse(b *testing.B) {
	dicter := DefaultDict()
	for i := 0; i < b.N; i++ {
		_ = dicter.ReadReverse("什么")
	}
}
