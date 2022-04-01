check_defined = \
	$(strip $(foreach 1,$1, \
		$(call __check_defined,$1,$(strip $(value 2)))))
__check_defined = \
	$(if $(value $1),, \
	$(error Undefined $1$(if $2, ($2))))

all: cdk-install swagger build run

swagger:
	@echo "Generating Swagger documentation..."
	# General Swagger documentation
	swag init \
		--dir ./,./docs/response \
		--exclude ./endpoints/admin \
		-g ./main.go

build:
	$(call check_defined, ENV, ENV of API)
	go build -ldflags "-X main.environment=$(ENV)"

run:
	$(call check_defined, CONF, ENV of API)
	./jumbotravel-api -c $(CONF)

test: build
	$(call check_defined, CONF, ENV of API)
	PWS_WORKER=DEV ./jumbotravel-api -c $(CONF)

cdk-install:
	@echo "[Install Deps] running phase"
	go get -u github.com/swaggo/swag/cmd/swag@v1.8.0