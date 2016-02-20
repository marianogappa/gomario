# gomario
Basic Golang text-mode Mario-like game - My first Go program

![Screenshot](screenshot.png?raw=true)

## Quick run on Darwin (Mac OS X)
```
wget https://github.com/MarianoGappa/gomario/raw/master/bin/darwin/main && chmod +x main && ./main
```

## Quick run on Linux
```
wget https://github.com/MarianoGappa/gomario/raw/master/bin/linux/main && chmod +x main && ./main
```

## Quick run on Windows
Download this file and double click on it, I guess?
https://github.com/MarianoGappa/gomario/raw/master/bin/windows/main.exe

## Build & Run
- Install go and set `$GOPATH`
https://golang.org/doc/install


```
cd $GOPATH/src && git clone git@github.com:MarianoGappa/gomario.git
cd gomario
go build .
./gomario
```

## Disclaimer
Mostly works; there is some race condition on gravity checking vs moving that sometimes makes Mario fall off perfectly good floors :P
