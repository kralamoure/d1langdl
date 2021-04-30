# d1langdl

`d1langdl` is a program for downloading Dofus 1 lang files.

## Requirements

- [Git](https://git-scm.com/)
- [Go](https://golang.org/)

## Build

```sh
git clone https://github.com/kralamoure/d1langdl
cd d1langdl
go build
```

## Installation

```sh
go get -u -v github.com/kralamoure/d1langdl
```

## Usage

```sh
d1langdl --help
```

### Output

```text
d1langdl is a program for downloading Dofus 1 lang files.

Find more information at: https://github.com/kralamoure/d1langdl

Options:
  -h, --help                Print usage information
  -d, --debug               Enable debug mode
  -u, --url string          Data URL (default "https://dofusretro.cdn.ankama.com/")
  -l, --languages strings   Language codes, separated by comma (default [de,en,es,fr,it,nl,pt])

Usage: d1langdl [options]
```
