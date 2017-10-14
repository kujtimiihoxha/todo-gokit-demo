FROM golang

RUN mkdir -p /go/src/github.com/kujtimiihoxha/todo-gokit-demo

ADD . /go/src/github.com/kujtimiihoxha/todo-gokit-demo

RUN curl https://glide.sh/get | sh
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

RUN cd /go/src/github.com/kujtimiihoxha/todo-gokit-demo && glide install

ENTRYPOINT  watcher -run github.com/kujtimiihoxha/todo-gokit-demo/todo/cmd -watch github.com/kujtimiihoxha/todo-gokit-demo/todo
