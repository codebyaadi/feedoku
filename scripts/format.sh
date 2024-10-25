#!/bin/bash

# Find all Go files and format them one by one
echo "Formatting Go files..."
find . -name '*.go' -print0 | while IFS= read -r -d '' file; do
  echo "Formatting $file"
  gofmt -s -w "$file"
done

# Check for any remaining formatting issues
echo "Checking for formatting issues..."
unformatted=$(gofmt -l .)

if [ -z "$unformatted" ]; then
  echo "All files are properly formatted."
else
  echo "The following files still need formatting:"
  echo "$unformatted"
  exit 1
fi

echo "Formatting complete."
