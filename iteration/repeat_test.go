package iteration

import "testing"

func TestRepeat(t *testing.T) {
    t.Run("Running N times", func(t *testing.T) {
        got := Repeat("a", 6)
        want := "aaaaaa"

        if got != want {
            t.Errorf("Got %q and expected %q", got, want)
        }
    })

    t.Run("Repeat twice if not times parameter is provided", func(t *testing.T) {
        got := Repeat("b")
        want := "bb"

        if got != want {
            t.Errorf("Got %q and expected %q", got, want)
        }
    })
}

// Adding some benchmarks: https://pkg.go.dev/testing#hdr-Benchmarks
// Benchmarks run when passing the flag -bench=.
func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a", 6)
    }
}
