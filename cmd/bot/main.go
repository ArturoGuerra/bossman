package main

import (
    "github.com/arturoguerra/bossman/internal/router"
    "github.com/arturoguerra/bossman/internal/handlers"
    "github.com/arturoguerra/bossman/internal/config"
    "github.com/bwmarrin/discordgo"
    "fmt"
    "os"
)

var token string

func init () {
    token = os.Getenv("TOKEN")

    if token == "" {
        os.Exit(3)
    }
}


func main () {
    dbot, err := discordgo.New("Bot " + token)

    if err != nil {
        fmt.Println(err)
        return
    }

    r, err := router.New(dbot)

    if err != nil {
        fmt.Println(err)
        return
    }

    dbot.addHandler(r.Handler)
    dbot.addHandler(handlers.OnReady)

    err = dbot.Open()
    if err != nil {
        fmt.Println(err)
        return
    }

    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <- sc

    dbot.Close()
}
