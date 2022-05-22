package main

import (
	"bytes"
	"fmt"
	"github.com/ainilili/enter/util"
	"time"
)

type rhythm struct {
	ctx      *context
	frames   []*frame
	ready    int
	interval int
}

type frame struct {
	ctx    *context
	point  int
	length int
}

type context struct {
	frame   int
	index   int
	over    bool
	builder *bytes.Buffer
	results map[int]int
}

func (r *rhythm) listen() {
	go func() {
		for !r.ctx.over {
			util.Await()
			r.ctx.results[r.ctx.frame] = r.ctx.index
		}
	}()
}

func (f *frame) sprint(ready string) []byte {
	b := f.ctx.builder
	b.Reset()
	b.WriteString(ready)
	for i := 0; i < f.length; i++ {
		if f.point == i {
			b.WriteByte('-')
			continue
		}
		b.WriteByte(' ')
	}
	return b.Bytes()
}

func newRhythm(ready int, interval int, length int, points []int) *rhythm {
	ctx := &context{builder: &bytes.Buffer{}, results: map[int]int{}}
	fs := make([]*frame, len(points))
	for i, point := range points {
		fs[i] = &frame{ctx: ctx, point: point, length: length}
	}
	return &rhythm{ctx: ctx, ready: ready, frames: fs, interval: interval}
}

func main() {
	r := newRhythm(3, 100, 20, []int{3, 5, 7, 6, 15, 7, 18, 16, 2})
	r.listen()
	ready := util.RepeatString(" ", r.ready)
	for k, f := range r.frames {
		bs := f.sprint(ready)
		for i := 0; i < len(bs); i++ {
			if _, ok := r.ctx.results[k]; ok {
				break
			}
			ob := bs[i]
			bs[i] = '|'
			r.ctx.frame = k
			r.ctx.index = i - r.ready
			fmt.Printf("\r=%s", bs)
			bs[i] = ob
			time.Sleep(time.Duration(r.interval) * time.Millisecond)

		}
		if v, ok := r.ctx.results[k]; ok {
			if v == f.point {
				fmt.Print("pass")
			} else {
				fmt.Print("miss")
			}
		} else {
			fmt.Print("\nmiss")
		}
		fmt.Print("\n")
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}
