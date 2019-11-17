package voice

import (
    "github.com/arturoguerra/bossman/internal/router"
    "fmt"
)

func New(r *router.Router) {
    r.On("bossman", soundHandler("bossman"))
    r.On("boy", soundHandler("boy"))
    r.On("musklex", soundHandler("musklex"))
    r.On("xenith", soundHandler("xenith"))
    r.On("kenith", soundHandler("kenith"))
    r.On("random", soundHandler("random"))
}

func soundHandler(person string) func(ctx *router.Context) {
    return func(ctx *router.Context) {
        for _, vs := range ctx.Guild {
            vs.UserID == ctx.Author.ID {
                err := playSound(ctx.Session, ctx.Guild.GuildID, vs.ChannelID, bossman)
                if err != nil {
                    fmt.Println("Error playing sound:", err)
                }

                return
            }
        }
    }
}

func playSound(s, *router.Context.Session, gID, cID, who string) (err error) {
}
