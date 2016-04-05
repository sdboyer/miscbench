package miscbench_test

import (
	"os"
	"path"
	"strconv"
	"testing"
)

func BenchmarkMkdirAll1(b *testing.B) {
	td := os.TempDir() + "mkdt"
	for i := 0; i < b.N; i++ {
		el := strconv.Itoa(i)
		os.MkdirAll(path.Join(td, el), 0777)
	}

	os.RemoveAll(td)
}

func BenchmarkMkdirAll2(b *testing.B) {
	td := os.TempDir() + "mkdt"
	for i := 0; i < b.N; i++ {
		el := strconv.Itoa(i)
		os.MkdirAll(path.Join(td, el, el), 0777)
	}

	os.RemoveAll(td)
}

func BenchmarkMkdirAll3(b *testing.B) {
	// ...for some reason this is taking 3x as long as i had it in vsolver.
	// w...t...f...
	td := os.TempDir() + "mkdt"
	for i := 0; i < b.N; i++ {
		el := strconv.Itoa(i)
		os.MkdirAll(path.Join(td, el, el, el), 0777)
	}

	os.RemoveAll(td)
}
