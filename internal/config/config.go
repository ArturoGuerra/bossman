package config

import "github.com/arturoguerra/bossman/internal/structs"

func New(token, prefix string) *structs.Config {
    return &structs.Config{
        Token: token,
        Prefix: prefix,
    }
}
