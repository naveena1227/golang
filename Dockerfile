FROM golang:latest
LABEL maintainer="naveena1227<naveenathollamadugu@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 8001
CMD ["./main"]

