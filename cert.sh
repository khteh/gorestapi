#!/bin/bash
openssl req -new -newkey rsa:4096 -x509 -nodes -days 365 -keyout server.key -out server.crt -subj "/C=SG/ST=Singapore/L=Singapore /O=Kok How Pte. Ltd./OU=GoRestAPI/CN=localhost/emailAddress=funcoolgeek@gmail.com" -passin pass:GoRestAPI
