# cmd lserver

`lserver` is a light tool to build local file server easily.

## Development Status

[![Build Status](https://travis-ci.org/Qs-F/lserver.svg?branch=master)](https://travis-ci.org/Qs-F/lserver)

## Installation

`go get -u github.com/Qs-F/lserver`

## Usage

- `-p`: Set server's exposing port (default is 8080)
- `-d`: Set which directory to expose (default is current directory)
- `-pub`: Switch internal or external server (default is internal)
- `-cors`: Swicth to use CORS or not (default is **use**)

## Example

`lserver -p 6060`

Expose current directory at `0.0.0.0:6060`.

---

`lserver -p 6060 -pub -cors`

Expose current directory at `localhost:6060`, and forbid to access over cross origin.
