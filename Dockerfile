FROM golang:1.22 as builder

WORKDIR /usr/src/app
COPY . .
RUN go build -v -o memhold ./...
RUN chmod 755 memhold

#################
FROM scratch

COPY --from=builder /usr/src/app/memhold /memhold
CMD ["/memhold"]