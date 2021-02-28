#FROM alpine:latest as base
#RUN apk --no-cache --update upgrade && apk --no-cache add ca-certificates
#RUN mkdir /buildtmp
#
#FROM scratch
#COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
#COPY --from=base /buildtmp /tmp
#
#WORKDIR /app
#
#COPY dist /app/
#COPY static /app/static
#
#ENTRYPOINT ["./agent"]

#FROM golang
FROM golang:alpine
COPY . /$GOPATH/src/github.com/dockermanage/
WORKDIR /$GOPATH/src/github.com/dockermanage/
#设置环境变量，开启go module和设置下载代理
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
#会在当前目录生成一个go.mod文件用于包管理
#RUN go mod init
#增加缺失的包，移除没用的包
RUN go mod tidy
#RUN go build app.go
EXPOSE 8080:8080
CMD ["go","run","main.go"]