ARG GO_VERSION=1.23.3
ARG MYSQL_URL

# Build container
FROM golang:${GO_VERSION}-alpine AS build

WORKDIR /src
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app

# Run container
FROM gcr.io/distroless/static-debian12:nonroot AS final
COPY --from=build /app /app
COPY db/migrations db/migrations
ENV MYSQL_URL = ${MYSQL_URL}
EXPOSE 1337

ENTRYPOINT [ "/app" ]


