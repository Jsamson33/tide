package domain

import (
	"testing"
	"time"
)

func TestDrawDaily(t *testing.T) {
	now := time.Now()
	h1 := DrawDaily(now)
	h2 := DrawDaily(now)

	if h1 != h2 {
		t.Errorf("DrawDaily is not deterministic for the same date: %v != %v", h1, h2)
	}

	tomorrow := now.AddDate(0, 0, 1)
	h3 := DrawDaily(tomorrow)
	// While it's theoretically possible for two different dates to have the same hexagram,
	// it's unlikely for adjacent days in a simple test.
	// This is just a basic sanity check.
	if h1 == h3 {
		t.Log("Note: DrawDaily produced same result for today and tomorrow. This is possible but rare (1/64).")
	}
}

func TestDrawLucky(t *testing.T) {
	h1 := DrawLucky()
	// Small sleep to ensure different seed if UnixNano is the same
	time.Sleep(1 * time.Millisecond)
	h2 := DrawLucky()

	if h1 == h2 {
		// 1/64 chance of collision, so we won't fail the test, but it's good to know.
		t.Log("Note: Two lucky draws produced the same hexagram.")
	}
}
