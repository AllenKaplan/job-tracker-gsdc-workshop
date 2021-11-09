go-build:
	go build -o app ./cmd

run: build
	./app

build:
	DOCKER_BUILDKIT=1 docker build -t job-tracker .
	
up:
	docker run  --rm -d -p 8080:8080 -e PORT='8080' \
		--name job-tracker job-tracker

push:
	docker tag job-tracker gcr.io/job-tracker-331623/job-tracker
	docker push gcr.io/job-tracker-331623/job-tracker


kill:
	docker kill job-tracker
	