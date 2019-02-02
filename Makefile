runnerdockerpath = .

build/runner/docker/docker:
	docker build $(runnerdockerpath) -t taask/runner-docker:$(shell cat $(runnerdockerpath)/.build/tag)

install/runner/docker: tag/runner/docker/dev build/runner/docker/docker
	helm template $(runnerdockerpath)/ops/chart \
	--set Tag=$(shell cat $(runnerdockerpath)/.build/tag) --set HomeDir=$(HOME) --set Count=$(count) \
	| linkerd inject --proxy-bind-timeout 30s - \
	| kubectl apply -f - -n taask

## this is essentially the same as install, without a build beforehand
scale/runner/docker:
	helm template $(runnerdockerpath)/ops/chart \
	--set Tag=$(shell cat $(runnerdockerpath)/.build/tag) --set HomeDir=$(HOME) --set Count=$(count) \
	| linkerd inject --proxy-bind-timeout 30s - \
	| kubectl apply -f - -n taask

logs/runner/docker:
	kubectl logs deployment/runner-docker runner-docker -n taask -f

logs/runner/docker/search:
	kubectl logs deployment/runner-docker runner-docker -n taask -f | grep $(search)

uninstall/runner/docker:
	kubectl delete deployment runner-docker -n taask

tag/runner/docker/dev:
	mkdir -p $(runnerdockerpath)/.build
	date +%s | openssl sha256 | base64 | head -c 12 > $(runnerdockerpath)/.build/tag

runner/docker/sandbox/build:
	docker build $(runnerdockerpath)/ops/sandbox -t taask/docker-sandbox:latest

runner/docker/sandbox/push:
	docker push taask/docker-sandbox:latest

runner/docker/echo/build:
	docker build $(runnerdockerpath)/ops/echo -t taask/echo