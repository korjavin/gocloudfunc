go mod download
go mod vendor
gcloud functions deploy mypkg --entry-point LambdaX --runtime go111 --trigger-http --set-env-vars TEST=1
