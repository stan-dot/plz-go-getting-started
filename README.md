# plz-go-getting-started

## lesson one  
the toolchain must be 
`
go_toolchain(
    name = "toolchain",
    version = "1.20",
)
`
not `go_toolchain`

## a convoluted way to add third party deps
https://please.build/codelabs/go_intro/index.html#8


and then the auto-geneerated list of indirect modules included (unnecessarily) a package for the original lib itself.
that caused `plz test` to break
https://please.build/codelabs/go_intro/index.html#8


