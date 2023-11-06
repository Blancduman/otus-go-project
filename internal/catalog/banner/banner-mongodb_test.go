package banner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Banner_ToModel(t *testing.T) {
	t.Parallel()

	doc := bannerDoc{
		ID:          1,
		Description: "Test 1",
	}

	want := Banner{
		ID:          1,
		Description: "Test 1",
	}

	assert.Equal(t, want, doc.toModel())
}

func Test_Banner_DocFromModel(t *testing.T) {
	t.Parallel()

	m := Banner{
		ID:          2,
		Description: "Test 2",
	}

	want := bannerDoc{
		ID:          2,
		Description: "Test 2",
	}

	assert.Equal(t, want, bannerDocFromModel(m))
}
