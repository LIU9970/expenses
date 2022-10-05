FROM golang:1.19

COPY src/expenses /go/src/expenses/

WORKDIR /go/src/expenses/
# RUN ls
# RUN go mod download

# COPY . /code
RUN go get github.com/gin-gonic/gin \
    && go get github.com/jinzhu/gorm \
    && go get github.com/go-sql-driver/mysql

CMD go run main.go