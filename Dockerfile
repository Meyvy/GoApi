FROM public.ecr.aws/docker/library/golang:bullseye
COPY ./code /src
WORKDIR /src 
RUN go mod download
EXPOSE 8080
ENTRYPOINT [ "go","run","main.go" ]