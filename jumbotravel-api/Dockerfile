FROM golang:1.17 AS build-app

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

# Environment default DEV
ARG ENV=DEV
# Install dependencies
RUN make cdk-install
# Build swagger-ui
RUN make swagger
# Build the app
RUN make build ENV=$ENV

FROM ubuntu:20.04 AS run-app

WORKDIR /var/run/task

COPY --from=build-app /usr/src/app/jumbotravel-api /var/run/task/jumbotravel-api
COPY assets /var/run/task/assets

# Run binary
CMD ["/var/run/task/jumbotravel-api", "-c", "/var/run/task/config.conf"]