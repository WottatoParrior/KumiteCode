#multistage build distroless container

FROM golang:1.16-buster AS build
WORKDIR /app


COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /KumiteCode

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /KumiteCode /KumiteCode

EXPOSE 10000


USER nonroot:nonroot

CMD ["/KumiteCode"]