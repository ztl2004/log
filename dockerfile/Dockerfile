FROM ubuntu
RUN apt-get update
RUN apt-get install -y wget ca-certificates build-essential
RUN apt-get install -y mercurial
RUN apt-get install -y git 
RUN apt-get install -y mysql-server
RUN wget --no-verbose https://go.googlecode.com/files/go1.2.src.tar.gz
RUN tar -v -C /usr/local -xzf go1.2.src.tar.gz
RUN cd /usr/local/go/src && ./make.bash --no-clean 2>&1
ENV PATH $PATH:/usr/local/go/bin
RUN mkdir /opt/arkors
ENV GOPATH /opt/arkors
RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee
RUN go get github.com/astaxie/beego/orm
RUN cd $GOPATH/src && git clone https://github.com/arkors/log.git
#RUN cd $GOPATH/src && git clone https://github.com/arkors/log.git
