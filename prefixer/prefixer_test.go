package prefixer

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	invocationCount := 0

	data := []struct {
		prefixFunc        func() string
		in                string
		callEnsureNewline bool
		out               string
	}{
		{
			func() string { return "prefix: " },
			"foo\nbar",
			false,
			"prefix: foo\nprefix: bar",
		},
		{
			func() string { return "prefix: " },
			"",
			false,
			"",
		},
		{
			func() string { return "prefix: " },
			"foo",
			true,
			"prefix: foo\n",
		},
		{
			func() string { return "prefix: " },
			"foo\n",
			true,
			"prefix: foo\n",
		},
		{
			func() string {
				prefixes := []string{"fizz ", "buzz "}
				prefix := prefixes[invocationCount%len(prefixes)]
				invocationCount++
				return prefix
			},
			"foo\nbar",
			false,
			"fizz foo\nbuzz bar",
		},
	}

	for _, d := range data {
		var buf bytes.Buffer
		p := New(&buf, d.prefixFunc)
		fmt.Fprintf(p, d.in)
		if d.callEnsureNewline {
			p.EnsureNewline()
		}
		out := buf.String()
		if out != d.out {
			t.Errorf("'%s' != '%s'", strings.Replace(out, "\n", "\\n", -1), strings.Replace(d.out, "\n", "\\n", -1))
		}
	}
}
