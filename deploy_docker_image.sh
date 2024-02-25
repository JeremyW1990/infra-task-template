#!/bin/bash

# Check if a version is passed as an argument
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <version>"
    exit 1
fi

# Assign the version to a variable
VERSION=$1

# Define your Docker image name
IMAGE_NAME="dataset-cleaner"

# Define your GCP Docker repository
GCP_REPO="us-central1-docker.pkg.dev/magic-gcp-test/magic-repo"

# Build the Docker image
docker build -t $IMAGE_NAME .

# Tag the Docker image with the version
docker tag $IMAGE_NAME $GCP_REPO/$IMAGE_NAME:v$VERSION

# Push the Docker image to the GCP repository
docker push $GCP_REPO/$IMAGE_NAME:v$VERSION

echo "Docker image $IMAGE_NAME:v$VERSION has been pushed to $GCP_REPO successfully."
