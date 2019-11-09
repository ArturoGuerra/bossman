package main

import (
    "github.com/arturoguerra/bossman/internal/router"
    "github.com/arturoguerra/bossman/internal/handlers"
    "github.com/arturoguerra/bossman/internal/config"
    "github.com/bwmarrin/discordgo"
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

var token string

func init () {
    token = os.Getenv("TOKEN")

    if token == "" {
        os.Exit(3)
    }
}


func main () {
    dgo, err := discordgo.New("Bot " + token)

    if err != nil {
        fmt.Println(err)
        return
    }

    cfg := config.New(token, "!")

    r := router.New(
        dgo,
        cfg,
    )

    dgo.AddHandler(func (_ *discordgo.Session, m *discordgo.MessageCreate) {
        r.Handler(m)
    })

    dgo.AddHandler(handlers.OnReady)

    err = dgo.Open()
    if err != nil {
        fmt.Println(err)
        return
    }

    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <- sc

    dgo.Close()
}
