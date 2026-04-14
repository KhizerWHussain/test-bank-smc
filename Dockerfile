FROM golang:1.22

WORKDIR /chaincode

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy
RUN go mod vendor

RUN go build -o chaincode .

EXPOSE 9999

CMD ["./chaincode"]