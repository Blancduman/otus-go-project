package ucb1_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/Blancduman/banners-rotation/internal/ucb1"
	"github.com/stretchr/testify/assert"
)

const TRIES = 1000

func Test_UCB1(t *testing.T) {
	rand.NewSource(time.Now().Unix())

	bannerStat := ucb1.BannerStat{
		1: {
			1: ucb1.Stat{
				Clicked: 1,
				Shown:   1,
			},
		},
		2: {
			1: ucb1.Stat{
				Clicked: 1,
				Shown:   1,
			},
		},
		3: {
			1: ucb1.Stat{
				Clicked: 1,
				Shown:   1,
			},
		},
	}

	clickFns := map[int64]func(){
		1: func() {
			bannerStat[1][1] = ucb1.Stat{
				Clicked: bannerStat[1][1].Clicked + rand.Int63n(2) + 1, //nolint: gosec
				Shown:   bannerStat[1][1].Shown + 1,
			}
		},
		2: func() {
			bannerStat[2][1] = ucb1.Stat{
				Clicked: bannerStat[2][1].Clicked + rand.Int63n(2), //nolint: gosec
				Shown:   bannerStat[2][1].Shown + 1,
			}
		},
		3: func() {
			bannerStat[3][1] = ucb1.Stat{
				Clicked: bannerStat[3][1].Clicked + rand.Int63n(2), //nolint: gosec
				Shown:   bannerStat[3][1].Shown + 1,
			}
		},
	}

	gimme := make(map[int64]int)

	for i := 0; i < TRIES; i++ {
		bannerID := ucb1.Next(bannerStat, 1)
		gimme[bannerID]++
		clickFns[bannerID]()
	}

	assert.GreaterOrEqual(t, float64(gimme[1])/TRIES*100, float64(80))
	assert.LessOrEqual(t, float64(gimme[2])/TRIES*100, float64(20))
	assert.LessOrEqual(t, float64(gimme[3])/TRIES*100, float64(20))
}
