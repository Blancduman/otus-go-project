package stat

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
