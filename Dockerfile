FROM golang

ADD main.go /
ADD templates/* /templates/
ADD /Users/guineveresaenger/samsung/go/src/github.com/go-sql-driver/mysql/* /go-sql-driver/mysql
ADD /Users/guineveresaenger/samsung/go/src/github.com/gorilla/mux/* /gorilla/mux

RUN go build -o /myfirstgoapp /main.go

EXPOSE 12345
ENTRYPOINT ["/myfirstgoapp"]
