FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY dist/pingme.exe /app/pingme.exe

ENTRYPOINT ["./pingme.exe"]