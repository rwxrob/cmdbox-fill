# Simple Form Filler Utility

[![GoDoc](https://godoc.org/cmdbox-fill?status.svg)](https://godoc.org/cmdbox-fill)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)
[![Go Report
Card](https://goreportcard.com/badge/cmdbox-fill)](https://goreportcard.com/report/cmdbox-fill)

This utility can be used as a standalone command or composed into a
[CmdBox] composite command.

[CmdBox]: <https://github.com/rwxrob/cmdbox>

## Install

As standalone ...

```
go install github.com/rwxrob/cmdbox-fill/fill@latest
# go get is deprecated
```

As composable CmdBox `init` module ...

```go
import (
  _ "github.com/rwxrob/cmdbox-fill"
)
```

## Usage

```
fill help
fill <field> ... < form.txt
```
