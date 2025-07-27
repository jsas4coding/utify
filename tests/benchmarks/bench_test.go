package benchmarks

import (
	"testing"

	"github.com/jsas4coding/utify"
	"github.com/jsas4coding/utify/pkg/formatter"
	"github.com/jsas4coding/utify/pkg/messages"
	"github.com/jsas4coding/utify/pkg/options"
)

func BenchmarkEcho(b *testing.B) {
	opts := options.Default()
	for i := 0; i < b.N; i++ {
		_, _ = formatter.Echo(messages.Success, "Benchmark test", opts)
	}
}

func BenchmarkSuccess(b *testing.B) {
	opts := utify.OptionsDefault()
	for i := 0; i < b.N; i++ {
		utify.Success("Benchmark test", opts)
	}
}

func BenchmarkSuccessWithBold(b *testing.B) {
	opts := utify.OptionsDefault().WithBold()
	for i := 0; i < b.N; i++ {
		utify.Success("Benchmark test", opts)
	}
}

func BenchmarkSuccessWithoutColor(b *testing.B) {
	opts := utify.OptionsDefault().WithoutColor()
	for i := 0; i < b.N; i++ {
		utify.Success("Benchmark test", opts)
	}
}

func BenchmarkSuccessf(b *testing.B) {
	opts := utify.OptionsDefault()
	for i := 0; i < b.N; i++ {
		utify.Successf("Benchmark test %d", opts, i)
	}
}

func BenchmarkGetSuccess(b *testing.B) {
	opts := utify.OptionsDefault()
	for i := 0; i < b.N; i++ {
		_, _ = utify.GetSuccess("Benchmark test", opts)
	}
}