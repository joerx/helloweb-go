IMAGE ?= quay.io/joerx/helloweb
TAG ?= latest

default: build

build:
	docker build -t $(IMAGE):$(TAG) .

run: build
	docker run -it --rm -p 8000:8000 $(IMAGE):$(TAG)

push:
	docker push $(IMAGE):$(TAG)
