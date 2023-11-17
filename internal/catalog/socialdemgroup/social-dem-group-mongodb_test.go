package socialdemgroup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SocialDemGroup_ToModel(t *testing.T) {
	t.Parallel()

	doc := socialDemGroupDoc{
		ID:          1,
		Description: "Test 1",
	}

	want := SocialDemGroup{
		ID:          1,
		Description: "Test 1",
	}

	assert.Equal(t, want, doc.toModel())
}

func Test_SlotDocFromModel(t *testing.T) {
	t.Parallel()

	m := SocialDemGroup{
		ID:          2,
		Description: "Test 2",
	}

	want := socialDemGroupDoc{
		ID:          2,
		Description: "Test 2",
	}

	assert.Equal(t, want, socialDemGroupDocFromModel(m))
}
