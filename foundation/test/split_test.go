package split

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{
		"simple":      {"a:b:cd:eee", ":", []string{"a", "b", "cd", "eee"}},
		"wrong sep":   {"a:b:cd:bee", "b", []string{"a:", ":cd:", "ee"}},
		"more sep":    {"a:b:cd:eee", "cd", []string{"a:b:", ":eee"}},
		"uni sep":     {"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
		"leading sep": {"沙河有沙又有河", "沙", []string{"河有", "又有河"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}

}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:cd:eee", ":")
	}

}

func TestMain(m *testing.M) {
	fmt.Println("before test")
	m.Run()
	fmt.Println("after test")
}
