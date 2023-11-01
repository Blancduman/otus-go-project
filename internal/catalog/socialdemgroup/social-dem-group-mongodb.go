package socialdemgroup

type socialDemGroupDoc struct {
	ID          int64  `bson:"_id"`
	Description string `bson:"description"`
}

func (d socialDemGroupDoc) toModel() SocialDemGroup {
	return SocialDemGroup(d)
}

func socialDemGroupDocFromModel(socialDemGroup SocialDemGroup) socialDemGroupDoc {
	return socialDemGroupDoc(socialDemGroup)
}
