/*This is us trying to break simple text/low-level repo parts of Github (and obviously report our findings to them)
   - help us & GitHub by submitting a PR!*/

package main

import (
    "encoding/json"
    "fmt"

    "github.com/google/go-github/github"
    "golang.org/x/oauth2"
)

var (
    personalAccessToken = "undefined"
)

type TokenSource struct {
    AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
    token := &oauth2.Token{
        AccessToken: t.AccessToken,
    }
    return token, nil
}

func main() {
    tokenSource := &TokenSource{
        AccessToken: personalAccessToken,
    }

    paOauthToken := oauth2.NewClient(oauth2.NoContext, tokenSource)
    client := github.NewClient(paOauthToken)

    user, _, err := client.Users.Get("undefined")
    if err != nil {
        fmt.Printf("client.Users.Get() faled with '%s'\n", err)
        return
    }

    d, err := json.MarshalIndent(user, "", "  ")
    if err != nil {
        fmt.Printf("json.MarshlIndent() failed with %s\n", err)
        return
    }

    fmt.Printf("User:\n%s\n", string(d))
}
