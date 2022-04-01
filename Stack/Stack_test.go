package structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Stack(t *testing.T) {
	asr := assert.New(t)

	stack := NewStack()
	asr.True(stack.IsEmpty(), "栈不为空")

	start, end := 0, 100
	for i := start; i < end; i++ {
		stack.Push(i)
		asr.Equal(i-start+1, stack.Len(), "Push 元素后坚持栈的长度")
	}

	for i := end - 1; i >= 0; i-- {
		asr.Equal(i, stack.Pop(), "从栈中 Pop元素出来")
	}

	asr.True(stack.IsEmpty(), "栈不为空")
}
