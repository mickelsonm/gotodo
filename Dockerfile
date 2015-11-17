FROM scratch

#first we need to build it like so:
#CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gotodo .

ADD gotodo /
EXPOSE 9090
ENTRYPOINT ["/gotodo"]
