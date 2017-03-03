FROM golang

ADD main.go /
ADD templates/* /templates/

RUN go build -o /myfirstgoapp /main.go

EXPOSE 12345
ENTRYPOINT ["/main"]
