#! /bin/bash

IMAGE_NAME=robocar-objects-detection

TAG=$(git describe)
FULL_IMAGE_NAME=docker.io/cyrilix/${IMAGE_NAME}:${TAG}

# Remove old images
podman manifest rm localhost/${IMAGE_NAME}:${TAG} ${FULL_IMAGE_NAME}

podman build . --platform linux/amd64,linux/arm64 --manifest ${IMAGE_NAME}:${TAG}
podman manifest push --all --rm localhost/${IMAGE_NAME}:${TAG} ${FULL_IMAGE_NAME}
