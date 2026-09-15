package main

import (
	"fmt"
	"time"
)

type RingBuffer struct {
	data       []*Data
	size       int
	lastInsert int
	nextRead   int
	emitTime   time.Time
}

type Data struct {
	Stamp time.Time
	Value string
}

func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		data:       make([]*Data, size),
		size:       size,
		lastInsert: -1,
	}
}

func (r *RingBuffer) Insert(input Data) {
	r.lastInsert = (r.lastInsert + 1) % r.size
	r.data[r.lastInsert] = &input

	if r.nextRead == r.lastInsert {
		r.nextRead = (r.nextRead + 1) % r.size
	}
}

func (r *RingBuffer) Emit() []*Data {
	output := []*Data{}
	for {
		if r.data[r.nextRead] != nil {
			output = append(output, r.data[r.nextRead])
			r.data[r.nextRead] = nil
		}
		if r.nextRead == r.lastInsert || r.lastInsert == -1 {
			break
		}
		r.nextRead = (r.nextRead + 1) % r.size
	}
	return output
}

func main() {
	rb := NewRingBuffer(5)
	fmt.Println("Empty test: ", rb)
	currentRune := 'a' - 1

	for i := 0; i < 10; i++ {
		currentRune++
		rb.Insert(Data{
			Stamp: time.Now(),
			Value: string(currentRune),
		})
	}
	fmt.Println("Full test: ", rb) 
}

