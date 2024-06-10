build:
	docker build -t distribuida-tuple-space .
server:
	docker run -it --rm --network=host distribuida-tuple-space
all: build server