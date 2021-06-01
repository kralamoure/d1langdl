# retrolangdl

[![CI](https://github.com/kralamoure/retrolangdl/actions/workflows/ci.yml/badge.svg)](https://github.com/kralamoure/retrolangdl/actions/workflows/ci.yml)

`retrolangdl` is a program for concurrently downloading Dofus Retro lang files.

## Build

```sh
git clone https://github.com/kralamoure/retrolangdl
cd retrolangdl
go build
```

## Installation

```sh
go install github.com/kralamoure/retrolangdl@latest
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
