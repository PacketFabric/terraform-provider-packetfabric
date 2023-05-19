#!/usr/bin/env bash
#set -x # uncomment for debug
set -e

# Define the tags we want to count
tags=("all" "smoke" "resource" "datasource" "cloud_router" "hosted_cloud" "dedicated_cloud" "core" "marketplace" "other")

# Initialize an array to store the counts
declare -a counts

# Initialize the counts for each tag to zero
for ((i=0; i<${#tags[@]}; i++)); do
    counts[i]=0
done

# Iterate over the test files
for file in internal/provider/*_test.go; do
    # Read the first line of each file
    read -r first_line < "$file"
    # Extract the tags from the first line
    tags_line="${first_line#*\/\/go:build }"
    # Iterate over the defined tags
    for ((i=0; i<${#tags[@]}; i++)); do
        tag=${tags[i]}
        # Check if the tag is present in the first line
        if [[ $tags_line == *"$tag"* ]]; then
            # Increment the count for the tag
            ((counts[i]++))
        fi
    done
done

# Print the tag counts
for ((i=0; i<${#tags[@]}; i++)); do
    echo "Tag '${tags[i]}' count: ${counts[i]}"
done

