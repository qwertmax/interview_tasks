package qsort_test

import (
	"github.com/qwertmax/interview_tasks/qsort"
	"math/rand"
	"sort"
	"testing"
	"time"
)

const size int = 1000000

var list = make([]int, size)
var prng = rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

func BenchmarkQsort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		for i := range list {
			list[i] = prng.Int()
		}
		b.StartTimer()

		qsort.Qsort(qsort.IntSlice(list), prng)
	}
}

func BenchmarkNativeQsort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		for i := range list {
			list[i] = prng.Int()
		}
		b.StartTimer()

		qsort.QuickSort(list)
	}
}

func BenchmarkSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		for i := range list {
			list[i] = prng.Int()
		}
		b.StartTimer()

		sort.Sort(sort.IntSlice(list))
	}
}
