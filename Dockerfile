FROM golang:1.21.1
RUN go version
ENV GOPATH=/
COPY ./ ./

RUN go mod download
RUN go build -o myhouse ./run_project.go

CMD ["./myhouse"]