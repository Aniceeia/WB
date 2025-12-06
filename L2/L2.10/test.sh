#!/bin/bash

BINARY="./sort"
TEST_DIR="test/testfiles"
TEMP_DIR=$(mktemp -d)
trap "rm -rf $TEMP_DIR" EXIT

PASSED=0
FAILED=0

test_case() {
    local name="$1"
    local input="$2"
    local args="$3"
    
    echo -n "Testing $name... "
    
    local our_output="$TEMP_DIR/our_${name}.txt"
    local unix_output="$TEMP_DIR/unix_${name}.txt"
    
    echo -e "$input" | $BINARY $args > "$our_output" 2>&1 || true
    echo -e "$input" | sort $args > "$unix_output" 2>&1 || true
    
    if diff -q "$our_output" "$unix_output" > /dev/null 2>&1; then
        echo "PASS"
        PASSED=$((PASSED + 1))
    else
        echo "FAIL"
        echo "  Our output:"
        cat "$our_output" | sed 's/^/    /'
        echo "  Unix output:"
        cat "$unix_output" | sed 's/^/    /'
        echo "  Diff:"
        diff "$our_output" "$unix_output" | sed 's/^/    /' || true
        FAILED=$((FAILED + 1))
    fi
}

test_file() {
    local name="$1"
    local file="$2"
    local args="$3"
    
    echo -n "Testing file $name... "
    
    local our_output="$TEMP_DIR/our_${name}.txt"
    local unix_output="$TEMP_DIR/unix_${name}.txt"
    
    $BINARY $args "$file" > "$our_output" 2>&1 || true
    sort $args "$file" > "$unix_output" 2>&1 || true
    
    if diff -q "$our_output" "$unix_output" > /dev/null 2>&1; then
        echo "PASS"
        PASSED=$((PASSED + 1))
    else
        echo "FAIL"
        echo "  Our output:"
        cat "$our_output" | sed 's/^/    /'
        echo "  Unix output:"
        cat "$unix_output" | sed 's/^/    /'
        echo "  Diff:"
        diff "$our_output" "$unix_output" | sed 's/^/    /' || true
        FAILED=$((FAILED + 1))
    fi
}

if [ ! -f "$BINARY" ]; then
    echo "Building binary..."
    go build -o "$BINARY" ./cmd/sort
fi

echo "=== Running tests ==="
echo

test_case "simple" "c\na\nb" ""
test_case "numeric" "3\n1\n2\n10" "-n"
test_case "reverse" "a\nc\nb" "-r"
test_case "unique" "a\na\nb\nb\nc" "-u"
test_case "numeric_reverse" "3\n1\n2\n10" "-n -r"
test_case "unique_numeric" "3\n1\n3\n2\n1" "-n -u"

test_file "simple_file" "$TEST_DIR/simple.txt" ""
test_file "numeric_file" "$TEST_DIR/numeric.txt" "-n"
test_file "duplicates_file" "$TEST_DIR/duplicates.txt" "-u"

test_case "keyfield" "1\tz\n2\ta\n3\tb" "-k 2"
test_case "numeric_keyfield" "10\tother\n2\tother\n1\tother" "-n -k 1"
test_case "ignore_blanks" "a  \nb\nc " "-b"

test_file "keyfield_file" "$TEST_DIR/keyfield.txt" "-k 2"
test_file "month_file" "$TEST_DIR/month.txt" "-M"
test_file "human_file" "$TEST_DIR/human.txt" "-h"

test_case_check() {
    local name="$1"
    local input="$2"
    local args="$3"
    
    echo -n "Testing $name... "
    
    local our_output="$TEMP_DIR/our_${name}.txt"
    local unix_output="$TEMP_DIR/unix_${name}.txt"
    
    echo -e "$input" | $BINARY $args > "$our_output" 2>&1; our_exit=$?
    echo -e "$input" | sort $args > "$unix_output" 2>&1; unix_exit=$?
    
    our_normalized=$(grep -o "disorder" "$our_output" || echo "")
    unix_normalized=$(grep -o "disorder" "$unix_output" || echo "")
    
    if [ "$our_exit" = "$unix_exit" ] && [ "$our_normalized" = "$unix_normalized" ]; then
        echo "PASS"
        PASSED=$((PASSED + 1))
    else
        echo "FAIL"
        echo "  Our exit code: $our_exit, Unix exit code: $unix_exit"
        echo "  Our output:"
        cat "$our_output" | sed 's/^/    /'
        echo "  Unix output:"
        cat "$unix_output" | sed 's/^/    /'
        FAILED=$((FAILED + 1))
    fi
}

test_case_check "check_sorted" "1\n2\n3" "-c"
test_case_check "check_not_sorted" "3\n1\n2" "-c"

echo
echo "=== Results ==="
echo "Passed: $PASSED"
echo "Failed: $FAILED"

if [ $FAILED -eq 0 ]; then
    echo "All tests passed!"
    exit 0
else
    echo "Some tests failed!"
    exit 1
fi

