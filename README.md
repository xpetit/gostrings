# gostrings

`gostrings` displays the string literals in Golang source files with their number of occurrences.

## Installation

```
go install github.com/xpetit/gostrings@latest
```

## Usage

```shell
gostrings -help
```

```shell
gostrings "$(go env GOROOT)/src/os" | sort -n | tail -n 20
```

> ```
>   5 symlink
>   6 android
>   6 chdir
>   6 open
>   6 plan9
>   6 readdir
>   6 stat
>   6 |0
>   7 truncate
>   8 \
>   8 pipe
>   8 rename
>   9 :
>   9 chmod
>   9 write
>   9 |1
>  13 windows
>  14 .
>  16 /
> 517 strings
> ```
