FROM golang:1.20

WORKDIR /chaincode

COPY . .

RUN go mod tidy
RUN go build -o chaincode .

EXPOSE 9999

CMD ["./chaincode"]