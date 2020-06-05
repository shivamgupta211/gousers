FROM golang
WORKDIR $GOPATH/src/gousers
COPY ./ ./
RUN go get github.com/go-pg/pg
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure
RUN go build
EXPOSE 2369
CMD ["go", "run", "gousers"]
