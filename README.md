# no-name [![CircleCI](https://circleci.com/gh/Yinkozi/no-name.svg?style=svg&circle-token=a18ffbc369b8ddcf8de823bc2a1eeb628509fcb7)](https://circleci.com/gh/Yinkozi/no-name)


## Environment Installation
1. Visit https://golang.org/doc/install
2. Install glide (MAC Osx : brew install glide)
3. Configure your ssh key so your able to pull the private repository of the dependencies manager (ssh-add <your_ssh_github_key>)

## Compile
```
git clone https://github.com/Yinkozi/no-name
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
curl --retry 10 --retry-delay 5 'http://localhost/setup.php' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: en-US,en;q=0.8,fr;q=0.6' -H 'Upgrade-Insecure-Requests: 1' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.95 Safari/537.36' -H 'Content-Type: application/x-www-form-urlencoded' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8' -H 'Cache-Control: max-age=0' -H 'Referer: http://localhost/setup.php' -H 'Connection: keep-alive' --data 'create_db=Create+%2F+Reset+Database' --compressed
curl --retry 10 --retry-delay 5 'http://localhost:8081/set-up-database.php'
```
3 - Launch tests  
```
go test (go list ./... | grep -v /vendor/)
```
