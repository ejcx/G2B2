###Golang Gorilla Bit Butcher
Brute force low entropy golang sessions. This file also contains a list of secrets that were included in the first 30 pages of a search on github.com of repositories that initialized a session with a published session key.

###Arguments
####-n: Name
The name of the cookie assigned to you by the server

####-v: Value
The value of the cookie assigned to you by the server

####-f: File
A dictionary of newline delimited passwords you would like to try

####-s: File
A particular secret you would like to guess.

###Example
    ./g2b2 -n=session\-name -v=MTQyNDQ4OTEwMHxEdi1CQkFFQ180SUFBUkFCRUFBQU1QLUNBQUlHYzNSeWFXNW5EQVVBQTJadmJ3WnpkSEpwYm1jTUJRQURZbUZ5QTJsdWRBUUNBRlFEYVc1MEJBSUFWZz09fEX2BtX_bJQshoq38z5ByNLvtxl3xixPriJ4Xe9ywKfq -f=asdf -s=something\-very\-secret
