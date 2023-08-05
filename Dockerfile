# FROM golang:1.20.7
FROM scratch

ENV GOPROXY https://mirrors.aliyun.com/goproxy/,direct

WORKDIR $GOPATH/src/github.com/aifuxi/jianyu-blog-api

COPY . $GOPATH/src/github.com/aifuxi/jianyu-blog-api

# RUN go build .

EXPOSE 9001

ENTRYPOINT [ "./jianyu-blog-api" ]

# 0. 在 Windows 上编译出 Linux 的可执行文件 （注意不要在powershell命令行下执行，需要在git bash的命令行环境执行，因为下面设置环境变量的方式是Linux的）
# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o jianyu-blog-api .

# 1. 构建镜像
# docker build -t jianyu-blog-sctratch .

# 2. 准备 MySQL
# docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -v docker-mysql:/var/lib/mysql -d mysql:8

# 3. 
# docker run --link mysql:mysql -p 9001:9001 jianyu-blog-sctratch
