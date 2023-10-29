package banner

type bannerDoc struct {
	ID          int64  `bson:"_id"`
	Description string `bson:"description"`
}

func (d bannerDoc) toModel() Banner {
	return Banner(d)
}

func bannerDocFromModel(banner Banner) bannerDoc {
	return bannerDoc(banner)
}
