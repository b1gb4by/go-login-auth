###############################
# Builder Container
###############################
FROM golang:1.16-alpine as builder
ENV CGO_ENABLED=0
WORKDIR /go/src/app/
RUN apk --no-cache add git \
    && go get -u github.com/cosmtrek/air \
    && go get github.com/google/wire/cmd/wire
COPY api/go.* ./
RUN go mod download
COPY api/ ./
RUN go build main.go wire_gen.go
CMD ["air", "wire_gen.go"]

###############################
# Release Container
###############################
FROM alpine:latest AS release
RUN apk --no-cache add ca-certificates
ARG APP_USER_NAME='gopher'
ARG APP_USER_GROUP=$APP_USER_NAME
RUN addgroup -g 10000 -S $APP_USER_NAME \
    && adduser -u 10000 -S $APP_USER_NAME -G $APP_USER_GROUP \
    && echo $APP_USER_NAME ALL=\(ALL\) NOPASSWD:ALL > /etc/sudoers \
    && echo $APP_USER_NAME:$APP_USER_GROUP | chpasswd
COPY --from=builder /go/src/app/main /app
USER $APP_USER_NAME
CMD ["/app"]
