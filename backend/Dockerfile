FROM golang:alpine

WORKDIR /go/src/github.com/AllegroTechDays/poz_Two_Tired/backend
COPY . .

RUN go-wrapper download 
RUN go-wrapper install 

EXPOSE 3000

CMD ["go-wrapper", "run"]