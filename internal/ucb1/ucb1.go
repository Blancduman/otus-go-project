package ucb1

import (
	"math"
)

type BannerStat map[int64]SocialDemGroupStat

type SocialDemGroupStat map[int64]Stat

type Stat struct {
	Clicked int64
	Shown   int64
}

func Next(bannersStat BannerStat, scGroup int64) int64 {
	var selectedBannerID int64
	var maxValue float64

	totalShowsAmount := countShowsAmount(bannersStat, scGroup)

	for bannerID, bannerStat := range bannersStat {
		if selectedBannerID == 0 {
			selectedBannerID = bannerID
		}

		if bannerStat[scGroup].Shown == 0 {
			return bannerID
		}

		clicks := float64(bannerStat[scGroup].Clicked)
		shown := float64(bannerStat[scGroup].Shown)
		target := clicks/shown + math.Sqrt((2.0*math.Log(totalShowsAmount))/shown)

		if target > maxValue {
			maxValue = target
			selectedBannerID = bannerID
		}
	}

	return selectedBannerID
}

func countShowsAmount(bannersStat BannerStat, scGroup int64) float64 {
	var total int64 = 0

	for _, bannerStat := range bannersStat {
		total += bannerStat[scGroup].Shown
	}

	return float64(total)
}
