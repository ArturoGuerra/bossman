package main

import (
    "fmt"
    "regexp"
)

func main() {
    re, err := regexp.Compile(`!([\w\d]+).+`)

    if err != nil {
        fmt.Println(err)
    }

    res := re.FindStringSubmatch("!hello its me ive been wondering if after all this time youd like me to come over")

    for k, v := range res {
	  fmt.Printf("%d. %s\n", k, v)
    }
}
