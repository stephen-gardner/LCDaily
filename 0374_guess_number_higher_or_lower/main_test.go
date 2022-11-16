package main

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func getGuessFunc(pick int) func(int) int {
	return func(num int) int {
		if num == pick {
			return 0
		} else if num < pick {
			return 1
		} else {
			return -1
		}
	}
}

func test(t *testing.T, n, expected int) {
	guess = getGuessFunc(expected)
	if res := guessNumber(n); res != expected {
		t.Log("n =", n)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestGuessNumber(t *testing.T) {
	test(t, 1, 1)
	test(t, math.MaxInt32, math.MaxInt32)
	rand.Seed(time.Now().UnixMicro())
	for i := 0; i < 10_000; i++ {
		n := rand.Intn(math.MaxInt32-1) + 1
		pick := rand.Intn(n-1) + 1
		test(t, n, pick)
	}
}
