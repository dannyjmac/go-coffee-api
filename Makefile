check_install:
	which_swagger || go install github.com/go-swagger/go-swagger/cmd/swagger  

swagger:
	swagger generate spec -o ./swagger.yaml --scan-models
