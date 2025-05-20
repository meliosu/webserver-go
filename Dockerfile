FROM go:1.24
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /webserver
ENTRYPOINT [ "/webserver" ]
