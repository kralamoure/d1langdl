# retrolangdl

[![CI](https://github.com/kralamoure/retrolangdl/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/kralamoure/retrolangdl/actions/workflows/ci.yml)

`retrolangdl` is a program for concurrently downloading Dofus Retro lang files.

## Requirements

- [Git](https://git-scm.com/)
- [Go](https://golang.org/)

## Build

```sh
git clone https://github.com/kralamoure/retrolangdl
cd retrolangdl
go build
```

## Installation

```sh
go get -u -v github.com/kralamoure/retrolangdl
```

## Usage

```sh
retrolangdl --help
```

### Output

```text
retrolangdl is a program for concurrently downloading Dofus Retro lang files.

Find more information at: https://github.com/kralamoure/retrolangdl

Options:
  -h, --help                Print usage information
  -d, --debug               Enable debug mode
  -u, --url string          Data URL (default "https://dofusretro.cdn.ankama.com/")
  -l, --languages strings   Language codes, separated by comma (default [de,en,es,fr,it,nl,pt])

Usage: retrolangdl [options]
```
