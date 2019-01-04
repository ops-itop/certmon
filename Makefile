TAG=$(shell git rev-parse --short HEAD)

image:
	CGO_ENABLED=0 go build -o bin/certmon cmd/certmon/main.go
	docker build -t $(PREFIX)certmon .
ifneq ("$(PREFIX)","")
	docker push $(PREFIX)certmon:latest
	docker tag $(PREFIX)certmon $(PREFIX)certmon:$(TAG)
	docker push $(PREFIX)certmon:$(TAG)
endif
