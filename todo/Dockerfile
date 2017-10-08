FROM golang

RUN mkdir -p /go/src/github.com/kujtimiihoxha/todo-gokit-demo

ADD . /go/src/github.com/kujtimiihoxha/todo-gokit-demo

RUN curl https://glide.sh/get | sh

RUN cd /go/src/github.com/kujtimiihoxha/todo-gokit-demo && glide install

ENTRYPOINT  go install github.com/kujtimiihoxha/todo-gokit-demo/todo/cmd && /go/bin/cmd
