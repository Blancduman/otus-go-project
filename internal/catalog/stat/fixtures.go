//go:build integration

package stat

func fixture1() slotStatDoc {
	return slotStatDoc{
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
}

func fixture2() slotStatDoc {
	return slotStatDoc{
		ID: 2,
		BannerStat: map[int64]socialDemGroupStatDoc{
			3: map[int64]statDoc{
				3: {
					Clicked: 3,
					Shown:   3,
				},
			},
			4: map[int64]statDoc{
				4: {
					Clicked: 4,
					Shown:   4,
				},
			},
		},
	}
}

func fixture3() slotStatDoc {
	return slotStatDoc{
		ID: 3,
		BannerStat: map[int64]socialDemGroupStatDoc{
			5: map[int64]statDoc{
				5: {
					Clicked: 5,
					Shown:   5,
				},
			},
			6: map[int64]statDoc{
				6: {
					Clicked: 6,
					Shown:   6,
				},
			},
		},
	}
}
