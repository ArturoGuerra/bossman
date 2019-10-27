package router

import (
    "github.com/arturoguerra/bossman/internal/structs"
    "github.com/bwmarrin/discordgo"
)


type (
    Route struct {
        Command string
        Handler *interface{}
    }

    Router struct {
       Routes []*Route
       Session *discordgo.Session
       Prefix string


    }
}


func New(s *discordgo.Session, c *structs.MessageConfig) *Router {
    return &Router{
        Session: s,
        Prefix, c.Prefix,
    }
}

func (rt *Router) Handler(_ *discordgo.Session, m *discordgo.MessageCreate) {
}
