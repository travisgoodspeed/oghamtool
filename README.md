Howdy y'all,

I was typesetting an article for [Rob
Graham](https://github.com/robertdavidgraham), in which he included
text in the Ogham script, which was used fifteen hundred years ago to
write Old Irish.

Writing the text in a hex editor got tiresome, so I tossed together
this little tool in Golang to transcribe between ASCII and UTF-8
Ogham on the unix command line.

This tool is dedicated to the memory of Michael Mehaffey, who tried
valiantly to teach me Scots Gaelic in Knoxville, sometime around 2005.
Fuck cancer.

--Travis

## Usage

Install the package with `go install github.com/travisgoodspeed/oghamtool@latest`
or clone the repo and build it with `make clean all`.  The `latin`
flag converts Ogham to ASCII, and the `-ogham` flag converts ASCII
into Ogham.  Your terminal will need UTF-8 support.

```
air% oghamtool -ogham ">Rob Graham is a troll<"
᚛ᚏᚑᚁ ᚌᚏᚐᚆᚐᚋ ᚔᚄ ᚐ ᚈᚏᚑᚂᚂ᚜
air% oghamtool -latin "᚛ᚏᚑᚁ ᚌᚏᚐᚆᚐᚋ ᚔᚄ ᚐ ᚈᚏᚑᚂᚂ᚜"
`ROB GRAHAM IS A TROLL'
air%
```
