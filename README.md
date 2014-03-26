log
===

To use log module, please refer to the following steps:

1.Get arkors mysql image from this url:

https://github.com/arkors/images/blob/master/databases/mysql/Dockerfile

2.Build mysql image using this command:

docker build -t arkors:mysql .

3.Manully creat arkors database on mysql.(will provide a script soon)

4.Build log image using this command:

docker build -t arkors:log .

5.Link log and mysql image using this command:

docker run --link mysql:db -i -t --name log arkors:log bash

then type env to get the acctual ip of mysql

change the ip in conf to the new ip

6.Init log db using this command:

docker run arkors:log cd $GOPATH/src/log && go build -o main main.go && ./main orm syncdb

7.Start log using this command:

docker run arkors:log cd $GOPATH/src/log && ./main

8.Test log api using this command:

curl -X POST -d '{"log":"This is just a test"}' http://IP:9119/log
