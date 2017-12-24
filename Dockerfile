FROM golang

# Copy the current directory contents into the container
COPY . /go/src/github.com/lpxxn/godockerswarm/

WORKDIR /go/src/github.com/lpxxn/godockerswarm/

RUN go build

EXPOSE 8000

CMD ["./godockerswarm"]