FROM loads/alpine:3.8
LABEL maintainer="zhangzheming<1160627439@qq.com>"
ENV WORKDIR  /app
#设置工作目录
WORKDIR $WORKDIR
#暴露端口
EXPOSE 8000

# 拷贝应用程序
ADD ./temp/linux_amd64/main $WORKDIR/main
RUN chmod +x $WORKDIR/main
# 运行命令
CMD ./main
