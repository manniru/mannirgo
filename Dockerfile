FROM golang:alpine
ADD . /Users/mannir/dev/mannirgo
RUN go /Users/mannir/dev/mannirgo
CMD ["/go/bin/mannirgo"]
EXPOSE 1313