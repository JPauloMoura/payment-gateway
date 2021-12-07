FROM golang:1.17

WORKDIR /go/src

# permite que o container continue em execução
CMD ["tail", "-f", "/dev/null"]