package router

import (
    "github.com/arturoguerra/bossman/internal/structs"
    "github.com/bwmarrin/discordgo"
    "strings"
    "errors"
    "regexp"
)


type (
    Context struct {
        *discordgo.Message
        Channel *discordgo.Channel
        Guild *discordgo.Guild
        Session *discordgo.Session
    }

    HandlerFunc func(interface{})

    Route struct {
        Name string
        Handler HandlerFunc
    }

    Router struct {
       Routes []*Route
       Session *discordgo.Session
       Prefix string
    }
)

var ErrRouteAlreadyRegistered = errors.New("Route has already been registered")

// Route Methods
func (rt *Route) Match(name string) bool {
    if strings.ToLower(name) == strings.ToLower(rt.Name) {
        return true
    }

    return false
}

// Router Methods
func New(s *discordgo.Session, c *structs.Config) *Router {
    return &Router{
        Session: s,
        State: s.State,
        Prefix: c.Prefix,
    }
}

func (r *Router) Handler(m *discordgo.MessageCreate) error {
    if r.State.User.ID == m.Author.ID {
        return
    }

    c, err := r.State.Channel(m.ChannelID)
    if err != nil {
        return
    }

    g, err := r.State.Guild(c.GuildID)
    if err != nil {
        return
    }

    restr := r.Prefix + `([\w\d]+).+`
    re, err := regexp.Compile(restr)

    if err != nil {
        return err
    }

    name := re.FindStringSubmatch(m.Content)[1]

    if rt := r.Find(name); rt != nil {
        ctx := &Context{
            m,
            c,
            g,
            r.Session,
        }

        rt.Handler(ctx)
    }

    return nil
}

func (r *Router) Find(name string) *Route {
    for _, rt := range r.Routes {
        if rt.Match(name) {
            return rt
        }
    }

    return nil
}

func (r *Router) On(name string, handler HandlerFunc) error {
    route := &Route{
        Name: name,
        Handler: handler,
    }

    if rt := r.Find(route.Name); rt != nil {
        return ErrRouteAlreadyRegistered
    }

    r.Routes = append(r.Routes, route)
    return nil
}
