FROM golang

ADD myfirstgoapp.go /
ADD templates/* /templates/

RUN go build -o /myfirstgoapp /main.go

EXPOSE 12345
ENTRYPOINT ["/myfirstgoapp"]
