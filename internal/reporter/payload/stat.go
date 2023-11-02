package payload

import "time"

type ClickStat struct {
	SlotID           int64     `json:"slotId"`
	BannerID         int64     `json:"bannerId"`
	SocialDemGroupID int64     `json:"socialDemGroupId"`
	Timestamp        time.Time `json:"timestamp"`
}

type ShownStat struct {
	SlotID           int64     `json:"slotId"`
	BannerID         int64     `json:"bannerId"`
	SocialDemGroupID int64     `json:"socialDemGroupId"`
	Timestamp        time.Time `json:"timestamp"`
}

func (s ClickStat) Type() Type {
	return TypeClick
}

func (s ShownStat) Type() Type {
	return TypeShown
}
