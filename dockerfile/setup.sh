#log container setup
docker build -t arkors:log .

#link log and mysql
docker run --link arkors:mysql -i -t arkors:log

#mysql container setup
docker run arkors:log cd $GOPATH/src/log && go build -o main main.go && ./main orm
docker run arkors:log cd $GOPATH/src/log && $GOPATH/bin/bee run
