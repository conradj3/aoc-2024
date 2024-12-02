#!/bin/bash

# Function to read the file and parse it into two lists
read_file() {
    local filePath="$1"
    leftList=()
    rightList=()

    while IFS= read -r line; do
        # Split the line into two parts
        read -r left right <<< "$line"
        leftList+=("$left")
        rightList+=("$right")
    done < "$filePath"
}

# Function to calculate total distance (Part 1)
get_total_distance() {
    local sortedLeft=($(printf "%s\n" "${leftList[@]}" | sort -n))
    local sortedRight=($(printf "%s\n" "${rightList[@]}" | sort -n))

    local totalDistance=0
    for i in "${!sortedLeft[@]}"; do
        local diff=$((sortedLeft[i] - sortedRight[i]))
        if (( diff < 0 )); then
            diff=$(( -diff ))
        fi
        totalDistance=$((totalDistance + diff))
    done

    echo "$totalDistance"
}

# Function to calculate similarity score (Part 2)
get_similarity_score() {
    declare -A rightCounts
    local similarityScore=0

    # Count occurrences in the right list
    for num in "${rightList[@]}"; do
        ((rightCounts["$num"]++))
    done

    # Calculate similarity score based on left list
    for num in "${leftList[@]}"; do
        if [[ -n ${rightCounts["$num"]} ]]; then
            similarityScore=$((similarityScore + num * rightCounts["$num"]))
        fi
    done

    echo "$similarityScore"
}

# Main script logic
filePath="../../day1/input.txt"  # Change this to your input file path

# Step 1: Read the input file
read_file "$filePath"

# Step 2: Calculate total distance (Part 1)
totalDistance=$(get_total_distance)
echo "Part 1 - Total Distance: $totalDistance"

# Step 3: Calculate similarity score (Part 2)
similarityScore=$(get_similarity_score)
echo "Part 2 - Similarity Score: $similarityScore"
