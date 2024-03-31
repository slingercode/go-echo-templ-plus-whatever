# 👀

A basic go http app... I'm going to add tailwind, maybe HTMX, cache headers
and whatever comes to my mind

## Why?

Learning. Now I'm using only `net/http` in the server but it's just a way to
probe how a real world example can be implemented. I liked so far.

This [blog post](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/c) is really a good way to start i believe

## Instructions

- Make sure that `$GOPATH` is defined so that `templ` can be used
- Install `pnpm` so that tailwind may be accessible and run `pnpm i`
- Run `go mod tidy` to install local dependencies
