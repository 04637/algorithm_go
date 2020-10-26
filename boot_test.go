package algorithm_go

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	// case1: 正常匹配中间子串
	str := "hello hhhworldl"
	sub := "hworld"
	i := strStr(str, sub)
	if i != 8 {
		t.Errorf("Case1 failed: expected 7 but got %d\n", i)
	}

	// case2: 匹配尾部子串
	str = "hello world"
	sub = "world"
	i = strStr(str, sub)
	if i != 6 {
		t.Errorf("Case2 failed: expected 6 but got %d\n", i)
	}

	// case3: 子串长度过大
	str = "world"
	sub = "hello world"
	i = strStr(str, sub)
	if i != -1 {
		t.Errorf("Case3 failed: expected -1 but got %d\n", i)
	}

	// case4: 子串前缀匹配但未匹配结束
	str = "hello worl"
	sub = "world"
	if i != -1 {
		t.Errorf("Case4 failed: expected -1 but got %d\n", i)
	}
}

func TestSubSets(t *testing.T) {
	arr := []int{1, 2, 3}
	result := subsets(arr)
	fmt.Println(result)
}
