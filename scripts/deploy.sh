
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./dist/main ./cmd/lambda/lambda.go

zip -r -j ./dist/lambda.zip ./dist/*

aws s3 cp ./dist/lambda.zip s3://swagger-golang-api-example

aws cloudformation deploy --template-file ./deploy/template.yml --stack-name golang-swagger-api --capabilities CAPABILITY_NAMED_IAM