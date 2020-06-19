package sat

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path"
	"runtime"
	"sync"
)

type dict struct {
	mt          sync.RWMutex
	data        map[rune]rune
	dataReverse map[rune]rune
	opts        Options
}

func (d *dict) set(on, under rune) {
	d.mt.Lock()
	defer d.mt.Unlock()
	d.data[on] = under
	d.dataReverse[under] = on
}

func (d *dict) getData(char rune) rune {
	d.mt.RLock()
	defer d.mt.RUnlock()
	if s, ok := d.data[char]; ok {
		return s
	}
	return char
}

func (d *dict) getDataR(char rune) rune {
	d.mt.RLock()
	defer d.mt.RUnlock()
	if s, ok := d.dataReverse[char]; ok {
		return s
	}
	return char
}

func (d *dict) defaultIO() (io.Reader, error) {
	_, file, _, _ := runtime.Caller(1)
	base := path.Dir(file) + "/dict.txt"
	return os.Open(base)
}

func (d *dict) Load(reader io.Reader) error {
	var err error
	if reader == nil {
		reader, err = d.defaultIO()
		if err != nil {
			return err
		}
	}
	buf := bufio.NewScanner(reader)
	var simplified []rune
	var traditional []rune
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
	if len(simplified) != len(traditional) {
		return errors.New("the length of simplified varies from the length of traditional")
	}
	for i := 0; i < len(simplified); i++ {
		d.set(traditional[i], simplified[i])
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

func (d *dict) Read(s string) string {
	return read(s, d.getData)
}

func (d *dict) ReadReverse(s string) string {
	return read(s, d.getDataR)
}

var defaultDict Dicter

func DefaultDict(ops ...Option) Dicter {
	if defaultDict != nil {
		return defaultDict
	}
	d := &dict{
		data:        make(map[rune]rune),
		dataReverse: make(map[rune]rune),
	}
	var reader io.Reader
	var err error
	for _, op := range ops {
		op(&d.opts)
	}
	if d.opts.Path != "" {
		reader, err = os.Open(d.opts.Path)
		if err != nil {
			panic(err)
		}
	}
	d.Load(reader)
	return d
}
