package stat

import "github.com/Blancduman/banners-rotation/internal/ucb1"

type SlotStat struct {
	SlotID     int64
	BannerStat BannerStat
}

type BannerStat map[int64]SocialDemGroupStat

type SocialDemGroupStat map[int64]Stat

type Stat struct {
	Clicked int64
	Shown   int64
}

func (s *BannerStat) ToUCB1() ucb1.BannerStat {
	bSt := ucb1.BannerStat{}

	for k1, v1 := range *s {
		sg := ucb1.SocialDemGroupStat{}

		for k2, v2 := range v1 {
			sg[k2] = ucb1.Stat{
				Clicked: v2.Clicked,
				Shown:   v2.Shown,
			}
		}

		bSt[k1] = sg
	}

	return bSt
}
