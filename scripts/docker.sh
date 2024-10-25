#!/bin/bash

# Set the image name and tag
IMAGE_NAME="codebyaadi/rss-feed-agg"
IMAGE_TAG="latest"

# Build the Docker image
echo "Building Docker image: $IMAGE_NAME:$IMAGE_TAG"
docker build -t $IMAGE_NAME:$IMAGE_TAG .

# Check if the build was successful
if [ $? -eq 0 ]; then
  echo "Docker image built successfully: $IMAGE_NAME:$IMAGE_TAG"
else
  echo "Docker build failed. Please check the Dockerfile and try again."
  exit 1
fi

# Optionally, you can add more steps here, like pushing the image to a registry
docker push $IMAGE_NAME:$IMAGE_TAG

# End of script
