all: cdk-install swagger build run

swagger:
	@echo "Generating Swagger documentation..."
	# General Swagger documentation
	swag init \
		--dir ./,./docs/response \
		--exclude ./endpoints/admin \
		-g ./main.go

build:
	go build

run:
	./jumbotravel-api -c dev.conf

cdk-install:
	@echo "[Install Deps] running phase"
	go get -u github.com/swaggo/swag/cmd/swag@v1.8.0