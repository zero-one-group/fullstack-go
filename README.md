# fullstack-go

## Quick Start

Install dependencies:

```sh
go mod tidy && go mod vendor
```

Build the app:

```sh
go build -o ./build/fullstack-go main.go
```

### Hot Reload

Install [air](https://github.com/cosmtrek/air), cli tool for live reload for Go apps.

```sh
go install github.com/cosmtrek/air@latest
```

Run development:

```sh
air -c air.toml
```
