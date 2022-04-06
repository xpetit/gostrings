# gostrings

`gostrings` displays the string literals in non-test Golang source files.
It doesn't show the imports and struct tags.

## Installation

```
go install github.com/xpetit/gostrings@latest
```

## Usage

```shell
gostrings "$(go env GOROOT)/src/strings"
```

> ```
> strings.Reader.ReadAt: negative offset
> strings.Reader.UnreadByte: at beginning of string
> strings.Reader.UnreadRune: at beginning of string
> strings.Reader.UnreadRune: previous operation was not ReadRune
> strings.Reader.Seek: invalid whence
> strings.Reader.Seek: negative position
> strings.Reader.WriteTo: invalid WriteString count
> strings.NewReplacer: odd argument count
> strings: negative Repeat count
> strings: Repeat count causes overflow
> strings: illegal use of non-zero Builder copied by value
> strings.Builder.Grow: negative count
> ```
