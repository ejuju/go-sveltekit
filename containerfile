FROM node:18.9.0 AS build-website
WORKDIR /src
COPY website .
RUN npm install
RUN npm run build

FROM golang:1.19.1 as build-go-server
WORKDIR /usr/local/go/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
COPY --from=build-website /src/build/ website/build/
ARG CGO_ENABLED=0
RUN go build -ldflags "-s -w" -o ./main .

FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=build-go-server /usr/local/go/app/main ./
ENTRYPOINT [ "/app/main" ]