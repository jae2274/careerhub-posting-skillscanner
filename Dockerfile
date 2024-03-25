FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./careerhub/posting_service/skillscanner/ ./careerhub/posting_service/skillscanner/
RUN ls --recursive ./

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./careerhub/posting_service/skillscanner/

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/myapp /myapp

USER nonroot:nonroot

ENTRYPOINT ["/myapp"]
