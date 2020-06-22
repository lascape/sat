package sat

type Dicter interface {
	Init(opts ...Option) error
	Read(string) string
	ReadReverse(string) string
}

//func Load(dicter Dicter, io io.Reader) error {
//	return dicter.Load(io)
//}
//
//func Read(dicter Dicter, str string) string {
//	return dicter.Read(str)
//}
//
//func ReadReverse(dicter Dicter, str string) string {
//	return dicter.ReadReverse(str)
//}
