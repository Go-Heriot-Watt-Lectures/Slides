Introduction to the Go Programming Language
11 Mar 2014

Tom Townsend
MEng Software Engineering, Heriot Watt University
tt87@hw.ac.uk
@TomTowny

* Introduction

During the past couple of years the Go programming language has gained popularity due to its comprehensive standard library and excellent execution speed. Companies around the world are looking to Go more and more for infrastructure services. Designed by programmers of Unix and Plan9 fame.

* My introduction

- Started using Go through my 4th Year dissertation
.link http://github.com/RadioactiveMouse/rgo Rgo
- Using it now to do websocket server for another course
.link https://github.com/Pollchat/pollchat Pollchat
- Contributed to a popular open source project that wraps net/http
.link https://github.com/martini-conrib/render Render

* Go strengths

- A statically typed compiled language with some of the flexibility of a dynamically typed interpreted language.
- Garbage collected
- Designed with Concurrency at the heart of the language

* Go weaknesses

- Garbage collection isnt very mature although this is coming in 1.3.
- A new way of handling inheritance requires a bit of thinking about

* What about the code?

.play hello.go

* net/http

- very fast http library within the standard library
- regularly comes top of various benchmarks due to amount of throughput it can handle
- Utilised by many companies including Google and Cloudflare.
.link http://talks.golang.org/2013/oscon-dl.slide#36 Google
.link http://blog.cloudflare.com/what-weve-been-doing-with-go Cloudflare

* Show me the code

.play http.go

* Why is this relevant to us?

- Go is starting to be used for more and more infrastructure projects such as Docker and the new linux distro CoreOS.
- Docker use Go for most of the system including a new library to directly handle containerisation in Linux without the need for a container library like LXC.
- CoreOS have written 2 very exciting programs that handle automatic launching of new instances around a variety cloud providers.

* More http

.play http-client-server.go
