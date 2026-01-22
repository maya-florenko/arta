FROM golang:1.25-alpine AS build

WORKDIR /src

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o /app ./cmd/app/

FROM gcr.io/distroless/static-debian12

COPY --from=build /app /app

CMD ["/app"]
