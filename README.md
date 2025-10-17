# problem

[![License](https://img.shields.io/github/license/FollowTheProcess/problem)](https://github.com/FollowTheProcess/problem)
[![Go Reference](https://pkg.go.dev/badge/go.followtheprocess.codes/problem.svg)](https://pkg.go.dev/go.followtheprocess.codes/problem)
[![Go Report Card](https://goreportcard.com/badge/github.com/FollowTheProcess/problem)](https://goreportcard.com/report/github.com/FollowTheProcess/problem)
[![GitHub](https://img.shields.io/github/v/release/FollowTheProcess/problem?logo=github&sort=semver)](https://github.com/FollowTheProcess/problem)
[![CI](https://github.com/FollowTheProcess/problem/workflows/CI/badge.svg)](https://github.com/FollowTheProcess/problem/actions?query=workflow%3ACI)
[![codecov](https://codecov.io/gh/FollowTheProcess/problem/branch/main/graph/badge.svg)](https://codecov.io/gh/FollowTheProcess/problem)

[RFC-7807] Problem JSON structure in Go

## Project Description

`problem` is a single source of truth for an [RFC-7807] `application/problem+json` structure, use it whenever you need to return an error from a REST API and your
APIs will all be consistent 👌🏻

## Installation

> [!WARNING]
> This module requires the Go 1.25 jsonv2 experiment: `GOEXPERIMENT=jsonv2` in order to inline the `Extra` map when serializing
> See <https://go.dev/blog/jsonv2-exp#experimenting-with-jsonv2>

```shell
go get go.followtheprocess.codes/problem@latest
```

## Quickstart

Whenever you need a problem:

```go
prob := problem.Problem{
    Type:     "https://example.com/probs/out-of-credit",
    Title:    "Not enough credit",
    Detail:   "Your current balance is 30, but that costs 50",
    Instance: "/account/12345/msgs/abc",
    Status:   http.StatusBadRequest,
}
```

Or you can use the `New` function with a bunch of options if you like:

```go
prob := problem.New(
    problem.Type("https://example.com/probs/out-of-credit"),
    problem.Title("Not enough credit")
    problem.Detail("Your current balance is 30, but that costs 50")
    problem.Instance("/account/12345/msgs/abc"),
    problem.Status(http.StatusBadRequest)
)
```

And these will both serialize to the following JSON:

```json
{
  "type": "https://example.com/probs/out-of-credit",
  "title": "Not enough credit",
  "detail": "Your current balance is 30, but that costs 50",
  "instance": "/account/12345/msgs/abc",
  "status": 400
}
```

### Any Extras?

[RFC-7807] allows for arbitrary additional fields in this response to convey any additional context your error might have. Using `problem` you do this with the `Extra` map:

```go
prob := problem.Problem{
    Type:     "https://example.com/probs/out-of-credit",
    Title:    "Not enough credit",
    Detail:   "Your current balance is 30, but that costs 50",
    Instance: "/account/12345/msgs/abc",
    Status:   http.StatusBadRequest,
    Extra: map[string]any{
      "balance":  30,
      "accounts": []string{"/accounts/12345", "/accounts/67890"},
    },
}
```

Or with the `New` pattern:

```go
prob := problem.New(
    problem.Type("https://example.com/probs/out-of-credit"),
    problem.Title("Not enough credit")
    problem.Detail("Your current balance is 30, but that costs 50")
    problem.Instance("/account/12345/msgs/abc"),
    problem.Status(http.StatusBadRequest),
    problem.Extra("balance", 30),
    problem.Extra("accounts", []string{"/accounts/12345", "/accounts/67890"})
)
```

These will both serialize to the following JSON:

```json
{
  "type": "https://example.com/probs/out-of-credit",
  "title": "Not enough credit",
  "detail": "Your current balance is 30, but that costs 50",
  "instance": "/account/12345/msgs/abc",
  "status": 400,
  "balance": 30,
  "accounts": [
    "/accounts/12345",
    "/accounts/67890"
  ]
}
```

> [!TIP]
> There is also an `ExtraMap` to allow adding a whole map of extra variables in one go rather than one at a time as shown above

### Credits

This package was created with [copier] and the [FollowTheProcess/go-template] project template.

[copier]: https://copier.readthedocs.io/en/stable/
[FollowTheProcess/go-template]: https://github.com/FollowTheProcess/go-template
[RFC-7807]: https://datatracker.ietf.org/doc/html/rfc7807
