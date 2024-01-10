package textstyles

import (
	"testing"

	"github.com/walles/moar/twin"
	"gotest.tools/v3/assert"
)

func TestIsManPageHeading(t *testing.T) {
	assert.Assert(t, !isManPageHeading(""))
	assert.Assert(t, !isManPageHeading("A"), "Incomplete sequence")
	assert.Assert(t, !isManPageHeading("A\b"), "Incomplete sequence")

	assert.Assert(t, isManPageHeading("A\bA"))
	assert.Assert(t, isManPageHeading("A\bA B\bB"), "Whitespace can be not-bold")

	assert.Assert(t, !isManPageHeading("A\bC"), "Different first and last char")
	assert.Assert(t, !isManPageHeading("a\ba"), "Not ALL CAPS")

	assert.Assert(t, !isManPageHeading("A\bAX"), "Incomplete sequence")

	assert.Assert(t, !isManPageHeading(" \b "), "Headings do not start with space")
}

func TestManPageHeadingFromString(t *testing.T) {
	// Set a marker style we can recognize and test for
	ManPageHeading = twin.StyleDefault.WithForeground(twin.NewColor16(2))

	result := manPageHeadingFromString("A\bA B\bB")

	assert.Assert(t, result != nil)
	assert.Equal(t, len(result.Cells), 3)
	assert.Equal(t, result.Cells[0].Rune, twin.Cell{Rune: 'A', Style: ManPageHeading})
	assert.Equal(t, result.Cells[1].Rune, twin.Cell{Rune: ' ', Style: ManPageHeading})
	assert.Equal(t, result.Cells[2].Rune, twin.Cell{Rune: 'B', Style: ManPageHeading})
}
