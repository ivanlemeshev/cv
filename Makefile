generate-test-cv:
	@echo "Generating test CV"
	go run cmd/cv/main.go -file data/test/cv.yml -output test-cv.pdf
.PHONY: generate-test-cv
