Ark Log Module
===

##How To Deploy

To use log module, please refer to the following steps:

* Get arkors mysql image from this url:

```
https://github.com/arkors/images/blob/master/databases/mysql/Dockerfile
```

* Build mysql image using this command:

```
docker build -t arkors:mysql .
```

* Manully creat arkors database on mysql.(will provide a script soon)

* Build log image using this command:

```
docker build -t arkors:log .
```

* Link log and mysql image using this command:

```
docker run --link mysql:db -i -t --name log arkors:log bash
```

then type `env` to get the acctual **ip** of mysql, and change the ip in conf to the new ip

* Init log db using this command:

```
docker run arkors:log cd $GOPATH/src/log && go build -o main main.go && ./main orm syncdb
```

* Start log using this command:

```
docker run arkors:log cd $GOPATH/src/log && ./main
```

* Test log api using this command:

```
curl -X POST -d '{"log":"This is just a test"}' http://IP:9119/log
```

> We will deploy it to [CoreOS](http://coreos.com) and use [etcd](https://github.com/coreos/etcd)  for service discover.
