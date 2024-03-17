# syntax=docker/dockerfile:1

FROM ghcr.io/merliot/device

WORKDIR /app
COPY . .

RUN go generate ./...
RUN go build -tags prime -o /ps30m ./cmd

EXPOSE 8000

ENV PORT_PRIME=8000
CMD ["/ps30m"]
