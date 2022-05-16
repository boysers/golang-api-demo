docker build -t go-api-demo .

docker images | grep go-api-demo

docker run --rm -p 3000:3000 go-api-demo