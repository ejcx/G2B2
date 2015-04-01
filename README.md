###Golang Gorilla Bit Butcher
G2B2 does two things.
1) Perform dictionary attacks on low entropy Gorilla Sessions. 
2) If you recover or already have the key to a Gorilla Session, rebuild the session with arbitrary values.

###Arguments
####-n: Name
If you are rebuilding the session, the name of the session cookie you are rebuilding.
If you are attacking the session, the name of the session cookie you are attacking.

####-v: Value
The -v argument should only be used when attacking the session. This is the real session provided to you by the server.

####-f: File
If you are attacking the session, a list of passwords you would like to try to use to decrypt the session.
If you are rebuilding the session, a JSON string containing the values that you would like to have in the session

####-s: Secret
If you are attacking the session, a key you would like to use to attempt to decrypt the session
If you are rebuilding the session, the key used on the server

####-r: Rebuild
Set this flag to indicate that you wish to rebuild the session, otherwise it is assumed you wish to attack.

###Dependencies
This program depends on two parts of the Gorilla session package.
    go get github.com/gorilla/sessions
    go get github.com/gorilla/securecookie
###Example
####Attacking
    ./g2b2 -n session-name -v MTQyNDQ4OTEwMHxEdi1CQkFFQ180SUFBUkFCRUFBQU1QLUNBQUlHYzNSeWFXNW5EQVVBQTJadmJ3WnpkSEpwYm1jTUJRQURZbUZ5QTJsdWRBUUNBRlFEYVc1MEJBSUFWZz09fEX2BtX_bJQshoq38z5ByNLvtxl3xixPriJ4Xe9ywKfq -s something-very-secret
####Rebuilding
    ./g2b2 -n session-name -f values.json -s something-very-secret -r
