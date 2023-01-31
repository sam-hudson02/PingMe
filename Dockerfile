FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY dist/pingme.exe pingme.exe

ENTRYPOINT ["./pingme.exe"]