tag = 0.1.0
docker-registry = cmcgalliard/secondmate

build:
	docker build -t $(docker-registry):$(tag) -f ./ci/Dockerfile  ./
push:
	docker push $(docker-registry):$(tag)
deploy:
	helm upgrade --install secondmate.io ./ci/helm
ci: build push
all: build push deploy