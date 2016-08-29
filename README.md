# SLO PD report parser

This is a work in progress snapshot of a simple parser for the San Luis Obispo police log. It's all of six hours of work, so be wary. It has been tested under go 1.6 on Linux and OS X.
If you don't have the Go compiler installed, read the official [Go install guide](http://golang.org/doc/install).

`parser.go` has all the logic currently. To test it and see the current state of the project, try `cat testdata/* | go run parser.go | less` . These will be turned into real test cases at some point.

`codes.go` contains a random list of 'nature of the call' and disposition codes I found, as a potential for expanding any acronyms found.

`pull-police-log` is a bash script to regularly record the output of the online reports. A couple weeks' worth of output is in the `testdata/` directory. pull-police-log intentionally creates duplicates to ensure not missing any events.

`template.go` is no longer used. It was a quick first pass effort to automatically find the output template they use for displaying incidents. It was easier to write a regexp by hand instead, but the hour of effort on that was useful regardless.

ToDo:
- Each of the multiline regexp captures can include newlines, remove them
- split the description field into a set of tokens, expand the shorthand
- insert into a db for reports, ignore dupes
