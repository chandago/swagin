package swagin

import "testing"

const (
	limitedIntLimit   = 10
	limitedIntThreads = 20
	limitedIntIter    = 10000
)

func TestLimitedInt(t *testing.T) {
	limitedInt := NewLimitedInt(limitedIntLimit)
	for i:=0; i<limitedIntThreads; i++ {
		go func() {
			for j :=0; j<limitedIntIter; j++ {
				_ = limitedInt.Incr()
				counter := limitedInt.Get()
				println(counter)
				if counter > limitedIntLimit {
					t.Fatalf("counted out of limits: %d", counter)
				}
			}
		}()
	}
}
