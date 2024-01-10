#不在Golang容器中,Scratch镜像，简洁、小巧，基本是个空镜像
#IDE本地编译可执行文件CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-gin .
#再构建镜像docker build -t go-gin-scratch .

#FROM scratch
#LABEL authors="songj"
#WORKDIR $GOPATH/src/
#COPY . $GOPATH/src/
#EXPOSE 8000
#CMD ["./go-gin"]


#第二种放编译后可执行文件
#DOCKER_BUILDKIT=0 docker buildx  build --platform linux/amd64 -t go-gin:latest .

FROM alpine
WORKDIR /build
RUN  #go mod tidy && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-gin .
COPY go-gin .
ENV GOPROXY https://goproxy.cn,direct
ADD go.mod .
COPY conf ./conf
RUN  echo $(ls -1 /build)
CMD ["./go-gin"]

# 第二种使用builder分层构, 把builder中可执行文件COPY到运行镜像scratch中, conf文件夹可拷贝,也可在K8s中利用configMap定义
#go整合opencv
#FROM 10.203.36.4:8088/base/jdk/2.0.20220606.1-8u342-opencv3.4.16-gocv:v1 AS builder
#我的是mac m1 --platform linux/arm64/v8.  linux 指定linux/amd64, 其他自行百度找到自己系统platfor以及交叉编译功能
#DOCKER_BUILDKIT=0 docker buildx  build --platform linux/arm64/v8 -t go-gin:latest .

#FROM bitnami/golang:1.20.5 AS builder
#MAINTAINER songj
#WORKDIR /build
#ENV GOPROXY https://goproxy.cn,direct
#ADD go.mod .
#COPY . .
#RUN  echo $(ls -1 /build)
#ENV PKG_CONFIG_PATH /usr/local/lib64/pkgconfig
#ENV LD_LIBRARY_PATH /usr/local/lib64/
#RUN #go build .
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-gin .

##FROM alpine
##FROM 10.203.36.4:8088/base/jdk/2.0.20220606.1-8u342-opencv3.4.16-gocv:v1

#FROM scratch
#MAINTAINER songj
#WORKDIR /build
#COPY --from=builder /build/go-gin /build/go-gin
#COPY --from=builder /build/conf /build/conf
#ENV PKG_CONFIG_PATH /usr/local/lib64/pkgconfig
#ENV LD_LIBRARY_PATH /usr/local/lib64/
#CMD ["./go-gin"]
