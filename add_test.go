package main

import "testing"

func TestAdd(t *testing.T) {
	_ = Add("GO PProf demo testcase")
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add("GO PProf demo testcase for B")
	}
}
