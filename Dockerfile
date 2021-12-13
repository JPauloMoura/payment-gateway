FROM golang:1.17

WORKDIR /go/src
# adiciona o kafka
RUN apt-get update && apt-get install build-essential librdkafka-dev -y
# permite que o container continue em execução
CMD ["tail", "-f", "/dev/null"]