FROM golang:1.12
COPY . /src
WORKDIR /src
RUN go build -mod=vendor -o minecraft-sysl .

FROM heroku/heroku:18
WORKDIR /app
COPY --from=0 /src/minecraft-sysl /app
CMD ["./minecraft-sysl"]