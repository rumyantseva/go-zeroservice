[![Release](https://img.shields.io/badge/release-1.0.1-brightgreen.svg?style=default?style=flat)](https://github.com/rumyantseva/go-zeroservice/releases/latest)
[![Build Status](https://travis-ci.org/rumyantseva/go-zeroservice.svg?branch=master)](https://travis-ci.org/rumyantseva/go-zeroservice)
[![Go Report Card](https://goreportcard.com/badge/github.com/rumyantseva/go-zeroservice)](https://goreportcard.com/report/github.com/rumyantseva/go-zeroservice)

# go-zeroservice

A "zero" web service written in Go. This service respresents an example for this article: [Готовим сборку Go-приложения в продакшн](https://habrahabr.ru/post/337158/).

## Quick start to try how it works

- Use `make vendor` to update dependencies.
- Use `make check` to check if there are no problems in the source code.
- Use `make build` to prepare a build.

## What the files mean

Yes, there are a lot of configuration files here 🤓

- `Makefile` contains popular instructions to check and build the source code.
- `Gopkg.toml` was made automatically by [dep](https://github.com/golang/dep), so `Gopkg.lock` did. It contains configuration of external dependecies. If you need to know more about dependencies and dep, please, watch [this video](https://www.youtube.com/watch?v=eZwR8qr2BfI).
- `vendor` is a directory to store external dependencies, there is only [httrouter](https://github.com/julienschmidt/httprouter) here because it is my only dependency in this project. If I want to have production-readiness, I prefer to store this directory in git.
- `.travis.yml` describes CI configuration for [Travis CI](http://travis-ci.org/).
- And finally `app.go` and `version.go` contain the source code.
