# Fiber REST API

Simple REST API built with Go & Fiber

## About

This is a re-creation of [echo REST API](https://github.com/ReeceRose/echo-rest-api), which is a REST API for books that uses the Fiber web framework. This is to compare the frameworks and see which I enjoy more and will be using for upcoming projects.

## Clone repository

In the desired folder, run:

```bash
git clone git@github.com:ReeceRose/fiber-rest-api.git .
```

## Setup for development

Install all project dependencies

```bash
go mod download
```

## Run API (for development purposes)

```bash
go run .
```

## Build API

Depending on your system, this will either generate _fiber-rest-api_ or _fiber-rest-api.exe_

```bash
go build .
```

## Run production API

```bash
./fiber-rest-api # Note, this file name and exact syntax may differ depending on your system. This is simply and example.
```
