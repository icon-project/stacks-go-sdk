#!/bin/bash

# ----------------------------
# Configuration
# ----------------------------
REPO_ROOT="$(git rev-parse --show-toplevel)"

# GitHub repository details
OWNER="stacks-network"
REPO="stacks-core"
BRANCH="master"  # Change to 'main' if the default branch is 'main'
FOLDER_PATH="docs/rpc"
API_URL="https://api.github.com/repos/$OWNER/$REPO/contents/$FOLDER_PATH?ref=$BRANCH"

# Destination directory
DEST_DIR="$REPO_ROOT/pkg/rpc_client"
DOCS_DIR="$DEST_DIR/docs"

# GitHub Personal Access Token (Optional)
# Uncomment and set your token to increase rate limits
# TOKEN="your_personal_access_token"

# ----------------------------
# Functions
# ----------------------------

# Function to clean up the API repository
cleanup() {
  echo "Cleaning up: Removing $DOCS_DIR directory..."
  rm -rf "$DOCS_DIR"
}

# Ensure cleanup happens on script exit, regardless of success or failure
trap cleanup EXIT

# Function to download contents recursively
download_contents() {
    local url="$1"
    local current_dir="$2"

    # Make API request
    if [ -n "$TOKEN" ]; then
        response=$(curl -s -H "Authorization: token $TOKEN" "$url")
    else
        response=$(curl -s "$url")
    fi

    # Iterate over each item in the response
    echo "$response" | jq -c '.[]' | while read -r item; do
        type=$(echo "$item" | jq -r '.type')
        name=$(echo "$item" | jq -r '.name')
        download_url=$(echo "$item" | jq -r '.download_url')

        if [ "$type" == "file" ]; then
            # Create the directory if it doesn't exist
            mkdir -p "$current_dir"

            # Download the file
            echo "Downloading $name..."
            if [ -n "$TOKEN" ]; then
                curl -s -H "Authorization: token $TOKEN" -o "$current_dir/$name" "$download_url"
            else
                curl -s -o "$current_dir/$name" "$download_url"
            fi
        elif [ "$type" == "dir" ]; then
            # Recursively download the directory
            new_dir="$current_dir/$name"
            mkdir -p "$new_dir"
            new_url=$(echo "$item" | jq -r '.url')
            download_contents "$new_url" "$new_dir"
        fi
    done
}

# ----------------------------
# Main Script
# ----------------------------

# Create destination directory
mkdir -p "$DOCS_DIR"

# Start downloading
download_contents "$API_URL" "$DOCS_DIR"

echo "Download completed. Files are saved in '$DOCS_DIR'."

# Run openapi-generator
echo "Generating Stacks RPC Go Client using openapi-generator..."
openapi-generator generate \
  -i "$DOCS_DIR/openapi.yaml" \
  -g go \
  -o "$DEST_DIR" \
  --additional-properties=packageName=rpc_client,gitUserId=icon-project,gitRepoId=stacks-go-sdk \
  --skip-validate-spec

echo "Stacks RPC Go Client generation completed successfully."