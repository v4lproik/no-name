#!/bin/sh

#init credentials for webgoat
curl --retry 10 --retry-delay 5 'http://localhost:8080/WebGoat/register.mvc' -H 'Origin: http://localhost:8080' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: en-GB,da-DK;q=0.8,da;q=0.6,fr-CA;q=0.4,fr;q=0.2,en-US;q=0.2,en;q=0.2' -H 'Upgrade-Insecure-Requests: 1' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36' -H 'Content-Type: application/x-www-form-urlencoded' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8' -H 'Cache-Control: max-age=0' -H 'Referer: http://localhost:8080/WebGoat/register.mvc' -H 'Connection: keep-alive' --data 'username=admintest&password=admintest&matchingPassword=admintest&agree=agree' --compressed

#init database for citizenstig
curl --retry 10 --retry-delay 5 'http://localhost:8081/set-up-database.php'

#init dwva
export USER_TOKEN=$(curl 'http://localhost/setup.php' -c dwvacookie -H 'Accept-Encoding: gzip, deflate, sdch' -H 'Accept-Language: en-GB,da-DK;q=0.8,da;q=0.6,fr-CA;q=0.4,fr;q=0.2,en-US;q=0.2,en;q=0.2' -H 'Upgrade-Insecure-Requests: 1' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8' -H 'Referer: http://localhost/setup.php' -H 'Connection: keep-alive' -H 'Cache-Control: max-age=0' --compressed 2>/dev/null | grep user_token | cut -d "'" -f 6)
curl --retry 10 'http://localhost/setup.php' -L -b dwvacookie -H 'Origin: http://localhost' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: en-GB,da-DK;q=0.8,da;q=0.6,fr-CA;q=0.4,fr;q=0.2,en-US;q=0.2,en;q=0.2' -H 'Upgrade-Insecure-Requests: 1' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36' -H 'Content-Type: application/x-www-form-urlencoded' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8' -H 'Cache-Control: max-age=0' -H 'Referer: http://localhost/setup.php' -H 'Connection: keep-alive' --data "create_db=Create+%2F+Reset+Database&user_token=$USER_TOKEN" --compressed

#init wordpress
curl 'http://localhost:8088/wp-admin/install.php?step=1' -H 'Origin: http://localhost:8088' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: en-GB,da-DK;q=0.8,da;q=0.6,fr-CA;q=0.4,fr;q=0.2,en-US;q=0.2,en;q=0.2' -H 'Upgrade-Insecure-Requests: 1' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36' -H 'Content-Type: application/x-www-form-urlencoded' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8' -H 'Cache-Control: max-age=0' -H 'Referer: http://localhost:8088/wp-admin/install.php' -H 'Connection: keep-alive' --data 'language=' --compressed
curl 'http://localhost:8088/wp-admin/install.php?step=2' -H 'Origin: http://localhost:8088' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: en-GB,da-DK;q=0.8,da;q=0.6,fr-CA;q=0.4,fr;q=0.2,en-US;q=0.2,en;q=0.2' -H 'Upgrade-Insecure-Requests: 1' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36' -H 'Content-Type: application/x-www-form-urlencoded' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8' -H 'Cache-Control: max-age=0' -H 'Referer: http://localhost:8088/wp-admin/install.php?step=1' -H 'Connection: keep-alive' --data 'weblog_title=test&user_name=test&admin_password=test&admin_password2=test&admin_email=test%40test.com&Submit=Install+WordPress&language=' --compressed