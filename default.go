package sat

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path"
	"runtime"
	"sync"
)

type defaultDict struct {
	mt          sync.RWMutex
	data        map[rune]rune
	dataReverse map[rune]rune
	opts        Options
}

func (d *defaultDict) set(on, under rune) {
	d.mt.Lock()
	defer d.mt.Unlock()
	d.data[on] = under
	d.dataReverse[under] = on
}

func (d *defaultDict) getData(char rune) rune {
	d.mt.RLock()
	defer d.mt.RUnlock()
	if s, ok := d.data[char]; ok {
		return s
	}
	return char
}

func (d *defaultDict) getDataR(char rune) rune {
	d.mt.RLock()
	defer d.mt.RUnlock()
	if s, ok := d.dataReverse[char]; ok {
		return s
	}
	return char
}

func (d *defaultDict) defaultFile() (*os.File, error) {
	_, file, _, _ := runtime.Caller(1)
	base := path.Dir(file) + "/dict.txt"
	return os.Open(base)
}

func (d *defaultDict) Init(opts ...Option) error {
	for _, o := range opts {
		o(&d.opts)
	}
	return nil
}

func read(s string, f func(char rune) rune) string {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		r[i] = f(r[i]) //d.getData(r[i])
	}
	return string(r)
}

func (d *defaultDict) Read(s string) string {
	return read(s, d.getData)
}

func (d *defaultDict) ReadReverse(s string) string {
	return read(s, d.getDataR)
}

var d *defaultDict

func DefaultDict() Dicter {
	if d == nil {
		err := InitDefaultDict()
		if err != nil {
			log.Fatal(err)
		}
	}
	return d
}
func InitDefaultDict(opts ...Option) error {
	d = &defaultDict{
		data:        make(map[rune]rune),
		dataReverse: make(map[rune]rune),
	}
	d.Init(opts...)
	var (
		err  error
		file *os.File
	)
	var simplified []rune
	var traditional []rune
	if d.opts.Path != "" {
		file, err = os.Open(d.opts.Path)
		if err != nil {
			return err
		}
		buf := bufio.NewScanner(file)
		var i int
		for buf.Scan() {
			text := buf.Text()
			switch i {
			case 0:
				simplified = []rune(text)
			case 1:
				traditional = []rune(text)
			}
			i++
		}
	} else {
		simplified = defaultSimplified
		traditional = defaultTraditional
	}

	if len(simplified) != len(traditional) {
		return errors.New("the length of simplified varies from the length of traditional")
	}
	for i := 0; i < len(simplified) && i < len(traditional); i++ {
		d.set(traditional[i], simplified[i])
	}
	return nil
}
