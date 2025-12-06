.PHONY: build test test-simple test-numeric test-reverse test-unique test-keyfield test-month test-human test-check test-files clean

BINARY := sort
TEST_DIR := test/testfiles

build:
	@echo "Building binary..."
	@go build -o $(BINARY) ./cmd/sort

test: build
	@echo "=== Running all tests ==="
	@make test-simple
	@make test-numeric
	@make test-reverse
	@make test-unique
	@make test-keyfield
	@make test-month
	@make test-human
	@make test-check
	@make test-files
	@echo "=== All tests passed ==="

test-simple: build
	@echo -n "Testing simple sort... "
	@TEMP_DIR=$$(mktemp -d); \
	printf "c\na\nb\n" | ./$(BINARY) > $$TEMP_DIR/our.txt 2>&1; \
	printf "c\na\nb\n" | sort > $$TEMP_DIR/unix.txt 2>&1; \
	diff -q $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt > /dev/null && echo "PASS" || (echo "FAIL" && diff $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt && rm -rf $$TEMP_DIR && exit 1); \
	rm -rf $$TEMP_DIR

test-numeric: build
	@echo -n "Testing numeric sort... "
	@TEMP_DIR=$$(mktemp -d); \
	printf "3\n1\n2\n10\n" | ./$(BINARY) -n > $$TEMP_DIR/our.txt 2>&1; \
	printf "3\n1\n2\n10\n" | sort -n > $$TEMP_DIR/unix.txt 2>&1; \
	diff -q $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt > /dev/null && echo "PASS" || (echo "FAIL" && diff $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt && rm -rf $$TEMP_DIR && exit 1); \
	rm -rf $$TEMP_DIR

test-reverse: build
	@echo -n "Testing reverse sort... "
	@TEMP_DIR=$$(mktemp -d); \
	printf "a\nc\nb\n" | ./$(BINARY) -r > $$TEMP_DIR/our.txt 2>&1; \
	printf "a\nc\nb\n" | sort -r > $$TEMP_DIR/unix.txt 2>&1; \
	diff -q $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt > /dev/null && echo "PASS" || (echo "FAIL" && diff $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt && rm -rf $$TEMP_DIR && exit 1); \
	rm -rf $$TEMP_DIR

test-unique: build
	@echo -n "Testing unique sort... "
	@TEMP_DIR=$$(mktemp -d); \
	printf "a\na\nb\nb\nc\n" | ./$(BINARY) -u > $$TEMP_DIR/our.txt 2>&1; \
	printf "a\na\nb\nb\nc\n" | sort -u > $$TEMP_DIR/unix.txt 2>&1; \
	diff -q $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt > /dev/null && echo "PASS" || (echo "FAIL" && diff $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt && rm -rf $$TEMP_DIR && exit 1); \
	rm -rf $$TEMP_DIR

test-keyfield: build
	@echo -n "Testing keyfield sort... "
	@TEMP_DIR=$$(mktemp -d); \
	printf "1\tz\n2\ta\n3\tb\n" | ./$(BINARY) -k 2 > $$TEMP_DIR/our.txt 2>&1; \
	printf "1\tz\n2\ta\n3\tb\n" | sort -k 2 > $$TEMP_DIR/unix.txt 2>&1; \
	diff -q $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt > /dev/null && echo "PASS" || (echo "FAIL" && diff $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt && rm -rf $$TEMP_DIR && exit 1); \
	rm -rf $$TEMP_DIR

test-numeric-reverse: build
	@echo -n "Testing numeric reverse sort... "
	@TEMP_DIR=$$(mktemp -d); \
	printf "3\n1\n2\n10\n" | ./$(BINARY) -n -r > $$TEMP_DIR/our.txt 2>&1; \
	printf "3\n1\n2\n10\n" | sort -n -r > $$TEMP_DIR/unix.txt 2>&1; \
	diff -q $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt > /dev/null && echo "PASS" || (echo "FAIL" && diff $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt && rm -rf $$TEMP_DIR && exit 1); \
	rm -rf $$TEMP_DIR

test-unique-numeric: build
	@echo -n "Testing unique numeric sort... "
	@TEMP_DIR=$$(mktemp -d); \
	printf "3\n1\n3\n2\n1\n" | ./$(BINARY) -n -u > $$TEMP_DIR/our.txt 2>&1; \
	printf "3\n1\n3\n2\n1\n" | sort -n -u > $$TEMP_DIR/unix.txt 2>&1; \
	diff -q $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt > /dev/null && echo "PASS" || (echo "FAIL" && diff $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt && rm -rf $$TEMP_DIR && exit 1); \
	rm -rf $$TEMP_DIR

test-month: build
	@echo -n "Testing month sort... "
	@TEMP_DIR=$$(mktemp -d); \
	./$(BINARY) -M $(TEST_DIR)/month.txt > $$TEMP_DIR/our.txt 2>&1; \
	sort -M $(TEST_DIR)/month.txt > $$TEMP_DIR/unix.txt 2>&1; \
	diff -q $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt > /dev/null && echo "PASS" || (echo "FAIL" && diff $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt && rm -rf $$TEMP_DIR && exit 1); \
	rm -rf $$TEMP_DIR

test-human: build
	@echo -n "Testing human sort... "
	@TEMP_DIR=$$(mktemp -d); \
	./$(BINARY) -h $(TEST_DIR)/human.txt > $$TEMP_DIR/our.txt 2>&1; \
	sort -h $(TEST_DIR)/human.txt > $$TEMP_DIR/unix.txt 2>&1; \
	diff -q $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt > /dev/null && echo "PASS" || (echo "FAIL" && diff $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt && rm -rf $$TEMP_DIR && exit 1); \
	rm -rf $$TEMP_DIR

test-check: build
	@echo -n "Testing check sorted... "
	@TEMP_DIR=$$(mktemp -d); \
	printf "1\n2\n3\n" | ./$(BINARY) -c > $$TEMP_DIR/our.txt 2>&1; our_exit=$$?; \
	printf "1\n2\n3\n" | sort -c > $$TEMP_DIR/unix.txt 2>&1; unix_exit=$$?; \
	if [ $$our_exit -eq $$unix_exit ]; then echo "PASS"; else echo "FAIL" && rm -rf $$TEMP_DIR && exit 1; fi; \
	echo -n "Testing check not sorted... "; \
	printf "3\n1\n2\n" | ./$(BINARY) -c > $$TEMP_DIR/our.txt 2>&1; our_exit=$$?; \
	printf "3\n1\n2\n" | sort -c > $$TEMP_DIR/unix.txt 2>&1; unix_exit=$$?; \
	if [ $$our_exit -eq $$unix_exit ]; then echo "PASS"; else echo "FAIL" && rm -rf $$TEMP_DIR && exit 1; fi; \
	rm -rf $$TEMP_DIR

test-files: build
	@echo "=== Testing with files ==="
	@TEMP_DIR=$$(mktemp -d); \
	echo -n "Testing simple file... "; \
	./$(BINARY) $(TEST_DIR)/simple.txt > $$TEMP_DIR/our.txt 2>&1; \
	sort $(TEST_DIR)/simple.txt > $$TEMP_DIR/unix.txt 2>&1; \
	diff -q $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt > /dev/null && echo "PASS" || (echo "FAIL" && diff $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt && rm -rf $$TEMP_DIR && exit 1); \
	echo -n "Testing numeric file... "; \
	./$(BINARY) -n $(TEST_DIR)/numeric.txt > $$TEMP_DIR/our.txt 2>&1; \
	sort -n $(TEST_DIR)/numeric.txt > $$TEMP_DIR/unix.txt 2>&1; \
	diff -q $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt > /dev/null && echo "PASS" || (echo "FAIL" && diff $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt && rm -rf $$TEMP_DIR && exit 1); \
	echo -n "Testing duplicates file... "; \
	./$(BINARY) -u $(TEST_DIR)/duplicates.txt > $$TEMP_DIR/our.txt 2>&1; \
	sort -u $(TEST_DIR)/duplicates.txt > $$TEMP_DIR/unix.txt 2>&1; \
	diff -q $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt > /dev/null && echo "PASS" || (echo "FAIL" && diff $$TEMP_DIR/our.txt $$TEMP_DIR/unix.txt && rm -rf $$TEMP_DIR && exit 1); \
	rm -rf $$TEMP_DIR

clean:
	@rm -f $(BINARY)
