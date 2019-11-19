package test

import (
    "github.com/arturoguerra/bossman/internal/router"
    "fmt"
)

func New(r *router.Router) {
    r.On("test", testHandler)
}

func testHandler(ctx *router.Context) {
    fmt.Println(ctx)
}
