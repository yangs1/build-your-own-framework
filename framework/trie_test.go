package framework

import (
	"testing"
)

func TestNewTree(t *testing.T) {
	var arr2 = []struct{}{}
	var arr1 = []struct{}{}
	t.Log(append(arr1, arr2...))
}
