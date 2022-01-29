FROM golang:latest AS build
COPY . /wask-task
WORKDIR /wask-task
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/main.go


FROM alpine:latest
EXPOSE 8000
RUN mkdir wasm-task
COPY --from=build /wask-task/server /wasm-task/server
COPY --from=build /wask-task/.env /wasm-task/.env
COPY --from=build /wask-task/web /wasm-task/web
RUN ls wasm-task
RUN chmod +x wasm-task/server
ENTRYPOINT ["./wasm-task/server"]