FROM golang:1.15.6-alpine as builder
LABEL author="bienkma@ghtk.co"
WORKDIR /go/src/auto-calling/auto-calling-gateway
COPY . .
RUN apk add alpine-sdk gcc
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o service -tags musl main.go

FROM alpine:latest
ENV TZ=Asia/Ho_Chi_Minh
RUN apk --no-cache add tzdata
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app
ADD ./ci/run.sh /app/run.sh
RUN apk --no-cache add gcc tzdata \
    && chmod 755 /app/run.sh \
    && cp /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime \
    && echo "Asia/Ho_Chi_Minh" >  /etc/timezone
COPY --from=builder /go/src/auto-calling/auto-calling-gateway .
RUN rm -rf .git && rm -rf ci
CMD [ "/app/run.sh" ]
