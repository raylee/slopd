// parser for SLOPD police log format (and other cities)
package main

/*
There was a question at the meeting about whether there's a public
police log. There is, here:

    http://pdreport.slocity.org/policelog/rpcdsum.txt

It is posted daily during the week (M-F) and then the weekend logs are
posted by noon on Monday.

Some tips on understanding the format of these reports and the shorthand
notation:
https://docs.google.com/document/d/1GRgzenaLfqZubrfd4spXU6cz-7rLVibqjqJqucspHGg/pub

You can get email updates of incidents that happen in your area by
adding your address to this form:
https://docs.google.com/forms/d/e/1FAIpQLSd_F0icZlwPu60pCcxHeuSSkseBG6nzEiDXIvHGcqXZaRgGuA/viewform
*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode"
)

type entry struct {
	raw            []string
	id             string
	day            time.Time
	received       time.Time
	dispatched     time.Time
	arrived        time.Time
	cleared        time.Time
	call_type      string
	grid_location  string
	as_observed    string
	address        string
	clearance_code string
	resp_officer   string
	units          []string
	description    string
	call_comments  string
}

func (e *entry) append_line(line string) {
	if !is_separator(line) {
		e.raw = append(e.raw, strings.TrimRightFunc(line, unicode.IsSpace))
	}
}

func is_separator(line string) bool {
	return strings.Contains(line, "==============") || is_end(line)
}

func is_end(line string) bool {
	return strings.Contains(line, "-------------")
}

func parse_report(in io.Reader) <-chan entry {
	ec := make(chan entry)

	go func() {
		s := bufio.NewScanner(in)
		e := entry{}
		in_entry_header := false

		header_pattern := regexp.MustCompile(`[/0-9]{5,8}               San Luis Obispo Police Department`)

	parse_line:
		for s.Scan() {
			line := s.Text()

			if header_pattern.MatchString(line) {
				for s.Scan() {
					line = s.Text()
					if is_separator(line) {
						continue parse_line
					}
				}
			}

			if is_separator(line) {
				switch in_entry_header {
				case true:
					if len(e.raw) > 0 {
						ec <- e
					}
					e = entry{}
					in_entry_header = false
				case false:
					in_entry_header = true
				}
			}
			e.append_line(line)

			if is_end(line) {
				for s.Scan() {
					line := s.Text()
					if strings.Contains(line, "==============") {
						break
					}
				}
			}
		}

		close(ec)

		// Check for errors during `Scan`. End of file is
		// expected and not reported by `Scan` as an error.

		// if err := scanner.Err(); err != nil {
		// 	fmt.Fprintln(os.Stderr, "error:", err)
		// 	os.Exit(1)
		// }
	}()

	return ec
}

var unparsed []entry

func time_from(day, tm string) time.Time {
	// Deal with the specific time format and missing times
	// "08/26/16 Received:02:43 Dispatched:      Arrived:02:43 Cleared:04:08"

	day = strings.TrimSpace(day)
	tm = strings.TrimSpace(tm)

	loc, _ := time.LoadLocation("America/Los_Angeles")

	t, err := time.ParseInLocation("01/02/06 15:04", day+" "+tm, loc)
	if err != nil {
		t, _ = time.ParseInLocation("01/02/06", day, loc)
	}
	return t
}

// test the regexp parser via https://play.golang.org/p/mJSLTs2NBB
const pattern = `([0-9]*) (........) Received:(.....) Dispatched:(.....) Arrived:(.....) Cleared:(.*)
(?s:\s*)Type:(.*)Location:(.*)
As Observed:((?s).*)
Addr:((?s).*)Clearance Code:((?s).*)
Responsible Officer:(.*)
Units:(.*)
 Des:((?s).*)
(CALL COMMENTS:|TO VEHS)((?s).*)`

func (e *entry) parse_raw() {
	pat := regexp.MustCompile(pattern)
	res := pat.FindAllStringSubmatch(strings.Join(e.raw, "\n"), -1)
	if res != nil {
		r := res[0]

		for i := range r {
			r[i] = strings.TrimSpace(r[i])
		}

		e.id = r[1]
		e.day = time_from(r[2], "")
		e.received = time_from(r[2], r[3])
		e.dispatched = time_from(r[2], r[4])
		e.arrived = time_from(r[2], r[5])
		e.cleared = time_from(r[2], r[6])

		e.call_type, e.grid_location, e.as_observed = r[7], r[8], r[9]
		e.address, e.clearance_code, e.resp_officer = r[10], r[11], r[12]

		// Units responding, a list of unit ids
		// split "4265  ,4266  ,S8" into [4265 4266 S8]
		split_on_func := func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		}
		e.units = strings.FieldsFunc(r[13], split_on_func)

		e.description, e.call_comments = r[14], r[15]

		e.raw = nil
	} else {
		unparsed = append(unparsed, *e)
	}
}

func (e entry) String() string {
	return fmt.Sprintf(`
id             %v
day            %v
received       %v
dispatched     %v
arrived        %v
cleared        %v
call_type      %v
grid_location  %v
as_observed    %v
address        %v
clearance_code %v
resp_officer   %v
units          %v
description    %v
call_comments  %v
`,
		e.id, e.day, e.received, e.dispatched, e.arrived, e.cleared, e.call_type, e.grid_location, e.as_observed, e.address,
		e.clearance_code, e.resp_officer, e.units, e.description, e.call_comments)

}
func main() {
	ec := parse_report(os.Stdin)
	// t := entry_template{}
	for e := range ec {
		e.parse_raw()
		fmt.Print(e)
	}
	fmt.Printf("\n\nUnparsed:\n---------\n")
	for _, e := range unparsed {
		fmt.Println(strings.Join(e.raw, "\n"))
		fmt.Println("--------")
	}
}
