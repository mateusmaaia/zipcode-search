FROM golang:1.11

RUN go get github.com/revel/cmd/revel

RUN revel new zipcode-search

CMD ["sleep", "infinity"]
