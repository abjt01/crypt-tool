# crypt-tool
a simple CLI tool for encrypting and decrypting text and files using Go.

## what it does?

- Text encryption using a shift cipher
- File encryption using the AES-256-GCM algorithm
- Password-protected file security

## try it!
Text
```
go run . encrypt text "heylo world" 5 
go run . decrypt text "CZTGJ RJMGY" 5
```

File
```
go run . encrypt file test.jpeg
go run . decrypt file test.jpeg
```
Build
```
go build -o crypt
./crypt encrypt text "HELLO" 3
```

## NOTE !!!
-- file encryption overwrites the original file  
-- keep backups while testing it
