module github.com/vertex-language/arm64

go 1.23

// The container modules are developed in this tree beside this one, hence
// the replace directives; a published build drops them and keeps the
// requires.
require (
	github.com/vertex-language/elf v0.0.0
	github.com/vertex-language/macho v0.0.0
	github.com/vertex-language/pe v0.0.0
)

replace (
	github.com/vertex-language/elf => ../elf
	github.com/vertex-language/macho => ../macho
	github.com/vertex-language/pe => ../pe
)
