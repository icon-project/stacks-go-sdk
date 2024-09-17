#!/bin/bash
set -e

# Get the absolute path to the repository root
REPO_ROOT="$(git rev-parse --show-toplevel)"

# Define variables relative to the repository root
API_REPO="https://github.com/hirosystems/stacks-blockchain-api.git"
API_DIR="$REPO_ROOT/stacks-blockchain-api"
OPENAPI_JSON="docs/openapi.json"
GENERATED_DIR="$REPO_ROOT/pkg/stacks_blockchain_api_client"

# Function to clean up the API repository
cleanup() {
  echo "Cleaning up: Removing $API_DIR directory..."
  rm -rf "$API_DIR"
}

# Ensure cleanup happens on script exit, regardless of success or failure
trap cleanup EXIT

# Clone or update the stacks-blockchain-api repository
if [ -d "$API_DIR" ]; then
  echo "Updating existing $API_DIR repository..."
  cd "$API_DIR"
  git fetch origin
  cd ..
else
  echo "Cloning stacks-blockchain-api repository..."
  git clone "$API_REPO" "$API_DIR"
fi

# Install dependencies (only if node_modules doesn't exist)
if [ ! -d "$API_DIR/node_modules" ]; then
  echo "Installing npm dependencies..."
  cd "$API_DIR"
  npm install
  cd ..
else
  echo "npm dependencies already installed."
fi

# Build the project
echo "Building the stacks-blockchain-api project..."
cd "$API_DIR"
npm run build
cd ..

# Generate OpenAPI files
echo "Generating OpenAPI specifications..."
cd "$API_DIR"
node lib/openapi-generator.js
cd ..

# Check if openapi.json was generated
if [ ! -f "$API_DIR/$OPENAPI_JSON" ]; then
  echo "Error: $OPENAPI_JSON not found. Aborting."
  exit 1
fi

# Run openapi-generator
echo "Generating Stacks Blockchain API Go Client using openapi-generator..."
openapi-generator generate \
  -i "$API_DIR/$OPENAPI_JSON" \
  -g go \
  -o "$GENERATED_DIR" \
  --additional-properties=packageName=stacks_blockchain_api_client,gitUserId=icon-project,gitRepoId=stacks-go-sdk \
  --skip-validate-spec

echo "Stacks Blockchain API Go Client generation completed successfully."

# Cleanup is handled by the trap and cleanup function
