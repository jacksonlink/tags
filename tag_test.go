package tags

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_Tag(t *testing.T) {
	r := require.New(t)
	tag := New("input", Options{})
	r.Equal("input", tag.Name)
}

func Test_Tag_WithName(t *testing.T) {
	r := require.New(t)
	tag := New("br", Options{})
	r.Equal("br", tag.Name)
	r.Equal(`<br />`, tag.String())
}

func Test_Tag_NonVoid(t *testing.T) {
	r := require.New(t)
	tag := New("div", Options{})
	r.Equal("div", tag.Name)
	r.Equal(`<div></div>`, tag.String())
}

func Test_Tag_WithValue(t *testing.T) {
	r := require.New(t)
	tag := New("input", Options{
		"value": "Mark",
	})
	r.Equal(`<input value="Mark" />`, tag.String())
}

func Test_Tag_WithTimeValue(t *testing.T) {
	r := require.New(t)

	cases := map[string]string{
		"":           "0001-01-01",
		"01-02-2006": "01-01-0001",
		"01-02":      "01-01",
	}

	for format, expected := range cases {
		tag := New("input", Options{
			"value":  time.Time{},
			"format": format,
		})

		r.Equal(fmt.Sprintf(`<input value="%v" />`, expected), tag.String())
	}

}

func Test_Tag_WithBody(t *testing.T) {
	r := require.New(t)

	tag := New("div", Options{
		"body": "hi there!",
	})
	r.Equal(`<div>hi there!</div>`, tag.String())
	r.Nil(tag.Options["body"])
}

func Test_Tag_WithBody_And_BeforeTag(t *testing.T) {
	r := require.New(t)
	s := `<span>Test</span>`

	tag := New("div", Options{
		"body":       "hi there!",
		"before_tag": s,
	})
	r.Equal(`<span>Test</span><div>hi there!</div>`, tag.String())
	r.Nil(tag.Options["body"])
}

func Test_Tag_WithBody_And_AfterTag(t *testing.T) {
	r := require.New(t)
	s := `<span>Test</span>`

	tag := New("div", Options{
		"body":      "hi there!",
		"after_tag": s,
	})
	r.Equal(`<div>hi there!</div><span>Test</span>`, tag.String())
	r.Nil(tag.Options["body"])
}

func Test_Tag_String(t *testing.T) {
	r := require.New(t)

	tag := New("div", Options{
		"body": "hi there!",
	})
	r.Equal(`<div>hi there!</div>`, tag.String())
}

func Test_Tag_String_WithOpts(t *testing.T) {
	r := require.New(t)

	tag := New("div", Options{
		"body":  "hi there!",
		"class": "foo bar baz",
	})
	r.Equal(`<div class="foo bar baz">hi there!</div>`, tag.String())
}

func Test_Tag_String_SubTag(t *testing.T) {
	r := require.New(t)

	tag := New("div", Options{
		"body": New("p", Options{
			"body": "hi!",
		}),
	})
	r.Equal(`<div><p>hi!</p></div>`, tag.String())
}

func Test_Tag_String_With_BeforeTag_Opt(t *testing.T) {
	r := require.New(t)
	s := `<span>Test</span>`

	tag := New("div", Options{
		"before_tag": s,
	})

	r.Equal(`<span>Test</span><div></div>`, tag.String())
}

func Test_Tag_String_With_AfterTag_Opt(t *testing.T) {
	r := require.New(t)
	s := `<span>Test</span>`

	tag := New("div", Options{
		"after_tag": s,
	})

	r.Equal(`<div></div><span>Test</span>`, tag.String())
}

func Test_Tag_With_Another_Tag_As_BeforeTag(t *testing.T) {
	r := require.New(t)
	s := New("span", Options{"body": "Test"})

	tag := New("div", Options{"before_tag": s})

	r.Equal(`<span>Test</span><div></div>`, tag.String())
}

func Test_Tag_With_Another_Tag_As_AfterTag(t *testing.T) {
	r := require.New(t)
	s := New("span", Options{"body": "Test"})

	tag := New("div", Options{"after_tag": s})

	r.Equal(`<div></div><span>Test</span>`, tag.String())
}
