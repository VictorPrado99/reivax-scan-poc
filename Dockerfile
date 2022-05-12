# PROTOTYPE

FROM golang:1.18-alpine AS build

WORKDIR .

RUN go install

ENTRYPOINT ["reivax-scan-poc"]
CMD ["--help"]