#submits and runs the workflows / activities 

all: run-worker run-starter

help:      ## help for this makefile.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

run-worker: ##will attempt to download the build artifacts as defined in vars.json "artifactProvider".
	@go run ./worker/main.go & sleep 5

run-starter: ##will attempt to build a container using the Dockerfile artifacts from --get-artifacts
	@go run ./start/main.go 



