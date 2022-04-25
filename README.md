# GoMakeHelp

A simple and minimal command that provides documentation and help messages
for your makefiles. This piggybacks of [TJ's Mmake functionality](https://github.com/tj/mmake)
but extracts the help functionality as is as it's own CLI.

## Usage


```sh
gomakehelp # This will automatically find Makefile (or if not found makefile)
           # and print out the help message based on its comments.

gomakehelp --makefile ./some/path/to/makefile # You can use the `--makefile ` option to custom define the makefile you wanna use
```
