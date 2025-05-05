# preload-golang
LD_PRELOAD poc for golang

└─$ go build -buildmode=c-shared -o backdoor.so main.go
$ LD_PRELOAD=./backdoor.so python3   
Hello, world!
Python 3.13.2 (main, Mar 13 2025, 14:29:07) [GCC 14.2.0] on linux
Type "help", "copyright", "credits" or "license" for more information.
>>> exit() 
