# Fiber REST API

Simple REST API built with Go & Fiber

## About

This is a re-creation of [Chi REST API](https://github.com/ReeceRose/chi-rest-api) and [Echo REST API](https://github.com/ReeceRose/echo-rest-api). This is to compare the frameworks and see which I enjoy more and will be using for upcoming projects.

## Clone repository

In the desired folder, run:

```bash
git clone git@github.com:ReeceRose/fiber-rest-api.git .
```

## Setup for development

Install all project dependencies

```bash
go mod vendor
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
