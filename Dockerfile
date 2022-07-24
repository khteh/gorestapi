FROM khteh/ubuntu:latest
WORKDIR /app
ADD templates templates
#ADD go.mod ./
#ADD go.sum ./
#ADD *.go ./
#ADD fibonacci fibonacci
#ADD greetings greetings
RUN openssl req -new -newkey rsa:4096 -x509 -nodes -days 365 -keyout server.key -out server.crt -subj "/C=SG/ST=Singapore/L=Singapore /O=Kok How Pte. Ltd./OU=GoRestAPI/CN=localhost/emailAddress=funcoolgeek@gmail.com" -passin pass:GoRestAPI
ADD server.* ./
#RUN go get github.com/khteh/fibonacci
#RUN go get github.com/khteh/greetings
#RUN go mod download
#RUN go build
ADD restapi ./
EXPOSE 8080
CMD [ "./restapi" ]
