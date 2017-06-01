.DEFAULT_GOAL := help

help:       ## Show available options with this Makefile
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY : test
test:       ## Run all the tests
test:
	./test.sh

clean:      ## Clean the application
	@go clean -i ./...
	@rm -rf ./goupload


build:      ## Build the application
build:	clean
	@go build github.com/ansrivas/goupload


.PHONY : app_help
app_help:   ## Display flags accepted by the application
APP_HELP = "$(shell ./goupload --help)"
app_help: build
	@echo $(APP_HELP)

.PHONY: test_run
test_run:   ## Run the application in a test mode
test_run:	clean build
	@echo "Running now.."
	@./goupload --configPath ./config.yaml 
