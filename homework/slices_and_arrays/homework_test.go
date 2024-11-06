package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// // go test -v homework_test.go

type CircularQueue struct {
	values   []int
	firstIdx int
	lastIdx  int
	count    int
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values:   make([]int, size),
		firstIdx: 0,
		lastIdx:  0,
		count:    0,
	}
}

func (q *CircularQueue) increaseLastIdx() {
	q.lastIdx = (q.lastIdx + 1) % cap(q.values)
}

func (q *CircularQueue) increaseFirstIdx() {
	q.firstIdx = (q.firstIdx + 1) % cap(q.values)
}

func (q *CircularQueue) Push(value int) bool {
	if q.count == cap(q.values) {
		return false
	}

	q.values[q.lastIdx] = value
	q.count++
	q.increaseLastIdx()
	return true
}

func (q *CircularQueue) Pop() bool {
	if q.count == 0 {
		return false
	}

	q.count--
	q.increaseFirstIdx()

	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}

	return q.values[q.firstIdx]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}
	if q.lastIdx == 0 {
		return q.values[cap(q.values)-1]
	}
	return q.values[q.lastIdx-1]
}

func (q *CircularQueue) Empty() bool {
	return q.count == 0
}

func (q *CircularQueue) Full() bool {
	return q.count == cap(q.values)
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
