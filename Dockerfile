FROM golang

RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/mux

ADD main.go /
ADD templates/* /templates/


RUN go build -o /myfirstgoapp /main.go

EXPOSE 12345
ENTRYPOINT ["/myfirstgoapp"]
