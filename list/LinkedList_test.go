package list_test

import (
	"testing"

	"github.com/AppBlitz/strutureData/list"
)

func TestSinglyLinkedList_AppendStart(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var list list.SinglyLinkedList
			list.AppendStart(tt.data)
		})
	}
}
