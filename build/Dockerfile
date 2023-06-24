FROM postgres

EXPOSE 5432

COPY ./init.sql /docker-entrypoint-initdb.d/ 

# FROM golang:1.20.5-bookworm AS build-stage

# WORKDIR /server

# COPY go.mod go.sum ./
# # COPY go.sum ./
# RUN go mod download

# COPY . .

# ENV GOOS linux
# ENV CGO_ENABLED 0

# RUN go build -o /go-fit-tracker

# # Deploy app binary into lean image
# FROM gcr.io/distroless/base-debian11 AS build-release-stage

# COPY --from=build-stage /go-fit-tracker /go-fit-tracker

# EXPOSE ${HTTP_PORT}

# USER nonroot:nonroot

# CMD ["/go-fit-tracker"]
