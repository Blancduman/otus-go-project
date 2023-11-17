package slot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Slot_ToModel(t *testing.T) {
	t.Parallel()

	doc := slotDoc{
		ID:          1,
		Description: "Test 1",
	}

	want := Slot{
		ID:          1,
		Description: "Test 1",
	}

	assert.Equal(t, want, doc.toModel())
}

func Test_SlotDocFromModel(t *testing.T) {
	t.Parallel()

	m := Slot{
		ID:          2,
		Description: "Test 2",
	}

	want := slotDoc{
		ID:          2,
		Description: "Test 2",
	}

	assert.Equal(t, want, slotDocFromModel(m))
}
