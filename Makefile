BINARY_NAME=sudoku_coloring

all: build test done
build:
	@echo "> Building..."
	go build -o $(BINARY_NAME)
test:
	@echo "> Testing..."
	./$(BINARY_NAME) < samples/2x2.txt | diff samples/2x2_solved.txt -
	./$(BINARY_NAME) < samples/easy.txt | diff samples/easy_solved.txt -
	./$(BINARY_NAME) < samples/expert.txt | diff samples/expert_solved.txt -
done:
	@echo "\n> We are all good!"
