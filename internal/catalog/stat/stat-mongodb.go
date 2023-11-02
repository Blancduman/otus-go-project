package stat

type slotStatDoc struct {
	ID         int64         `bson:"_id"` // slotID
	BannerStat bannerStatDoc `bson:"bannerStat"`
}

type bannerStatDoc map[int64]socialDemGroupStatDoc

type socialDemGroupStatDoc map[int64]statDoc

type statDoc struct {
	Clicked int64 `bson:"clicked"`
	Shown   int64 `bson:"shown"`
}

func (d slotStatDoc) toModel() SlotStat {
	return SlotStat{
		SlotID:     d.ID,
		BannerStat: d.BannerStat.toModel(),
	}
}

func (d bannerStatDoc) toModel() BannerStat {
	m := make(BannerStat, len(d))

	for k, v := range d {
		m[k] = v.toModel()
	}

	return m
}

func (d socialDemGroupStatDoc) toModel() SocialDemGroupStat {
	m := make(SocialDemGroupStat, len(d))

	for k, v := range d {
		m[k] = v.toModel()
	}

	return m
}

func (s statDoc) toModel() Stat {
	return Stat(s)
}
