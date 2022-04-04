package structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	ast := assert.New(t)

	queue := NewQueue()
	ast.True(queue.IsEmpty(), "队列不为空")

	start, end := 0, 1000
	for i := start; i < end; i++ {

	}

}
