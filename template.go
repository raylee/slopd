package main

import (
	"fmt"
	"regexp"
	"strings"
)

// functions for automatically discovering the template format. Sorta works.
var varying = []rune("?")[0]

type entry_template struct {
	initialized bool
	base        []string
	pattern     []string
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func (t *entry_template) update_template(e *entry) {
	// first time updating our template? Take the first entry as our baseline
	if !t.initialized {
		fmt.Println("Initializing template")
		t.base = make([]string, len(e.raw))
		t.pattern = make([]string, len(e.raw))
		copy(t.base, e.raw)
		copy(t.pattern, e.raw)

		// replace the text in pattern with a const marker
		re := regexp.MustCompile(".")
		for i := range t.pattern {
			t.pattern[i] = re.ReplaceAllString(t.pattern[i], ".")
		}

		// fmt.Println("Base:\n", t.base, "\nPattern:\n", t.pattern, "\nRaw\n:", e.raw)

		t.initialized = true
	}

	// does the new record have more lines than our previous patterns? If so, extend them
	for len(t.base) < len(e.raw) {
		fmt.Println("Extending by one line")
		t.base = append(t.base, "")
		t.pattern = append(t.pattern, "")
	}

	// does the new record have lines longer than the previous? If so, mark those positions
	// as variable, and extend the base for later comparison
	for i := range e.raw {
		diff := len(e.raw[i]) - len(t.pattern[i])
		if diff > 0 {
			// fmt.Println("Extending line", i, "by", diff, "characters")
			extra := strings.Repeat(string(varying), diff)
			t.pattern[i] = t.pattern[i] + extra
			t.base[i] = t.base[i] + extra
		}
	}

	// character by character comparison
	for i := range e.raw {
		for j := range e.raw[i] {
			if t.base[i][j] != e.raw[i][j] {
				t.pattern[i] = replaceAtIndex(t.pattern[i], varying, j)
			}
		}
	}
}

func (t *entry_template) result() string {
	res := make([]string, len(t.base))
	copy(res, t.base)
	for i := range t.pattern {
		for j := range t.pattern[i] {
			if rune(t.pattern[i][j]) == varying {
				res[i] = replaceAtIndex(res[i], varying, j)
			}
		}
	}
	return strings.Join(res, "\n")
}
