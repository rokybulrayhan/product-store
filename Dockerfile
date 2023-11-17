FROM  golang:alpine
WORKDIR  /

COPY  .  ./

RUN go build -o  /product_store

EXPOSE 8082

EXPOSE 8082

ENTRYPOINT [ "/product_store"]