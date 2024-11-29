#!/bin/bash

# A script to fetch, merge, and push changes to the remote repository.

# Ensure we are in a Git repository
if [ ! -d .git ]; then
    echo "Error: Not a Git repository. Please run this script inside a Git project."
    exit 1
fi

# Define the remote and branch (customize if needed)
REMOTE="origin"
BRANCH="main"

echo "Fetching changes from remote repository..."
git fetch $REMOTE

echo "Pulling changes from the remote branch..."
git pull $REMOTE $BRANCH

if [ $? -ne 0 ]; then
    echo "Error: Failed to pull changes. Resolve any conflicts before continuing."
    exit 1
fi

echo "Pushing local changes to the remote repository..."
git push $REMOTE $BRANCH

if [ $? -eq 0 ]; then
    echo "Successfully synchronized with the remote repository!"
else
    echo "Error: Failed to push changes. Please check for issues."
    exit 1
fi
