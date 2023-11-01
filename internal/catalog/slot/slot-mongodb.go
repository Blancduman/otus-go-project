package slot

type slotDoc struct {
	ID          int64   `bson:"_id"`
	Description string  `bson:"description"`
	BannerIDs   []int64 `bson:"bannerIDs"`
}

func (d slotDoc) toModel() Slot {
	return Slot(d)
}

func slotDocFromModel(slot Slot) slotDoc {
	return slotDoc(slot)
}
