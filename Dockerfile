FROM golang:1.18-alpine

RUN apk add git
#create app folder
RUN mkdir /app

#set work dir
WORKDIR /app

#copy all to app folder
ADD . /app

#create exe file
RUN go build -o main .

#run exe file
CMD ["/app/main"]