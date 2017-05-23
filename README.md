# no-name [![CircleCI](https://circleci.com/gh/Yinkozi/no-name.svg?style=svg&circle-token=a18ffbc369b8ddcf8de823bc2a1eeb628509fcb7)](https://circleci.com/gh/Yinkozi/no-name)


## Environment Installation
1. Visit https://golang.org/doc/install
2. Install glide (MAC Osx : brew install glide)
3. Configure your ssh key so your able to pull the private repository of the dependencies manager (ssh-add <your_ssh_github_key>)

## Compile
```
git clone https://github.com/v4lproik/no-name
cd no-name
glide install
go build main.go
```

## Run
```
./main --help
```

## Test
1 - Download & Set up vulnerable boxes  
```
docker-compose up -d
```
2 - Configure vulnerable boxes' databases  
```
sh configure-vulnerable-boxes.sh
```
3 - Launch tests  
```
go test (go list ./... | grep -v /vendor/)
```
