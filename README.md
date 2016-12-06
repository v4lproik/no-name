# no-name

## Environment Installation
1. Visit https://golang.org/doc/install
2. Install glide (MAC Osx : brew install glide)

## Compile
```
git clone https://github.com/Yinkozi/no-name
cd no-name
go build main.go
```

## Run
```
./main --help
```

## Testing boxes
```
docker pull infoslack/dvwa
docker run -d -p 80:80 infoslack/dvwa
docker run -d -p 80:80 -p 3306:3306 -e MYSQL_PASS=p@ssw0rd infoslack/dvwa
```  
```
docker pull citizenstig/nowasp
docker run -d -p 8081:80 citizenstig/nowasp
```
```
docker run -p 8082:8080 webgoat/webgoat-container`
```