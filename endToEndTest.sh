#!/bin/bash
# First attemt for an end to end test
# use a special test file numbered 0, target number = 50
# we do 3 tries:  49 (expect Higher), 51 (expect Lower), 50 (expect Won)

#create a trap mechanism that restores the values in the test JSON after the test is ran
set -euo pipefail
JSON_FILE="./DB/0"
BACKUP_FILE="./DB/0_backup"

# Step1 - Backup the file before any modifications are done to it
cp "$JSON_FILE" "$BACKUP_FILE"

# Step2 - Define cleanup function that restores the original file
cleanup() {
    echo "Restoring game state..."
    mv "$BACKUP_FILE" "$JSON_FILE"
}

# Step 3 - Ensure cleanup always runs, even on error or CTRL+C
trap cleanup EXIT

id=0
set 49 51 50

echo "All params right now are as follows: $*"

curl -s "http://localhost:8080/api/guess/$id/?userguess=$1" | grep  -q "Higher"

if [[ $? -ne 0 ]]; then
    echo "Error. The result should contain  'Higher' "
else
    echo "Correct. The result contains 'Higher' "
fi

curl -s "http://localhost:8080/api/guess/0/?userguess=$2" | grep  -q "Lower"

if [[ $? -ne 0 ]]; then
    echo "Error. The result should contain  'Lower' "
else
    echo "Correct. The result contains 'Lower' "
fi

echo "DEBUG: \$3 = $3"

curl -s "http://localhost:8080/api/guess/0/?userguess=$3" | grep  -q "WON"
if [[ $? -ne 0 ]]; then
    echo "Error. The result should contain  'Won' "
else
    echo "Correct. The result contains 'Won' "
fi