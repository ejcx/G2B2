##Golang Gorilla Bit Butcher
G2B2 does two things.
1) Perform dictionary attacks on low entropy Gorilla Sessions. 
2) If you recover or already have the key to a Gorilla Session, rebuild the session with arbitrary values.

##Arguments
###Attacking Params
####-n: Name
If you are attacking the session, the name of the session cookie you are attacking.

####-v: Value
The base64 session provided to you by the server.

####-f: File
A list of passwords you would like to try to use to find the HMAC Key

####-s: Secret
A single key you would like to use to attempt to decrypt the session.

###Rebuilding Params
If you are rebuilding the session, the name of the session cookie you are rebuilding.

####-f: File
A JSON string containing the values that you would like to have in the session

####-s: Secret
The key used on the server

####-r: Rebuild
Indicate that you wish to rebuild the session

###Debug Session
####-d: Deserialize
Attempt to deserialize the session without the HMAC or Encryption keys. This will get you the values from inside the session if the session is integrity only, however it will fail if the sessions are encrypted and HMAC'ed.

####-v: Value
The base64 session provided to you by the server
##Dependencies
This program depends on two parts of the Gorilla session package.

    go get github.com/gorilla/sessions
    go get github.com/gorilla/securecookie

##Example
###Attacking
    ./g2b2 -n session-name -v MTQyODg2MTkwOHxEdi1CQkFFQ180SUFBUkFCRUFBQU1QLUNBQUlHYzNSeWFXNW5EQVVBQTJadmJ3WnpkSEpwYm1jTUJRQURZbUZ5QTJsdWRBUUNBRlFEYVc1MEJBSUFWZz09fGZWZuwf1oFeGx0R7d7EzCqJix5E-N4AGeXlMXNb0Uju
You can also specify to only attack with a specific secret value, which allows you to use this program with your own dictionary
```
./g2b2 -s something-very-secret -n session-name -v MTQyODg2MTkwOHxEdi1CQkFFQ180SUFBUkFCRUFBQU1QLUNBQUlHYzNSeWFXNW5EQVVBQTJadmJ3WnpkSEpwYm1jTUJRQURZbUZ5QTJsdWRBUUNBRlFEYVc1MEJBSUFWZz09fGZWZuwf1oFeGx0R7d7EzCqJix5E-N4AGeXlMXNb0Uju
```
###Rebuilding
    ./g2b2 -n session-name -f values.json -s something-very-secret -r
###Debugging
    ./g2b2 -d -v MTQyOTQwODY4MnxEdi1CQkFFQ180SUFBUkFCRUFBQU1QLUNBQUlHYzNSeWFXNW5EQVVBQTJadmJ3WnpkSEpwYm1jTUJRQURZbUZ5QTJsdWRBUUNBRlFEYVc1MEJBSUFWZz09fDUvf7LpwRm8JFqDqjN9x209B8jXSB7SMukKnyxi8KsC
