package stat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SlotStatDoc_ToModel(t *testing.T) {
	doc := slotStatDoc{
		ID: 1,
		BannerStat: map[int64]socialDemGroupStatDoc{
			1: map[int64]statDoc{
				1: {
					Clicked: 1,
					Shown:   1,
				},
			},
			2: map[int64]statDoc{
				2: {
					Clicked: 2,
					Shown:   2,
				},
			},
		},
	}

	want := SlotStat{
		SlotID: 1,
		BannerStat: map[int64]SocialDemGroupStat{
			1: map[int64]Stat{
				1: {
					Clicked: 1,
					Shown:   1,
				},
			},
			2: map[int64]Stat{
				2: {
					Clicked: 2,
					Shown:   2,
				},
			},
		},
	}

	assert.Equal(t, want, doc.toModel())
}
