# Day 1: Historian Hysteria

# Function to read the file and parse into two lists
function Read-File {
    param (
        [string]$filePath
    )
    # Initialize arrays to hold the left and right lists
    $leftList = @()
    $rightList = @()

    # Read file content
    Get-Content $filePath | ForEach-Object {
        $parts = $_ -split '\s+'
        if ($parts.Count -eq 2) {
            $leftList += [int]$parts[0]
            $rightList += [int]$parts[1]
        }
    }

    # Return both lists
    return @{ Left = $leftList; Right = $rightList }
}

# Function to calculate total distance (Part 1)
function Get-TotalDistance {
    param (
        [int[]]$leftList,
        [int[]]$rightList
    )

    # Sort the lists
    $sortedLeft = $leftList | Sort-Object
    $sortedRight = $rightList | Sort-Object

    # Calculate total distance
    $totalDistance = 0
    for ($i = 0; $i -lt $sortedLeft.Count; $i++) {
        $totalDistance += [math]::Abs($sortedLeft[$i] - $sortedRight[$i])
    }

    return $totalDistance
}

# Function to calculate similarity score (Part 2)
function Get-SimilarityScore {
    param (
        [int[]]$leftList,
        [int[]]$rightList
    )

    # Create a frequency dictionary for the right list
    $rightCounts = @{}
    foreach ($num in $rightList) {
        if ($rightCounts.ContainsKey($num)) {
            $rightCounts[$num]++
        } else {
            $rightCounts[$num] = 1
        }
    }

    # Calculate similarity score
    $similarityScore = 0
    foreach ($num in $leftList) {
        if ($rightCounts.ContainsKey($num)) {
            $similarityScore += $num * $rightCounts[$num]
        }
    }

    return $similarityScore
}

# Main Script Logic
$filePath = "../../day1/input.txt"  # Change this to your input file path

# Read the input file
$data = Read-File -filePath $filePath
$leftList = $data.Left
$rightList = $data.Right

# Part 1: Calculate total distance
$totalDistance = Get-TotalDistance -leftList $leftList -rightList $rightList
Write-Output "Part 1 - Total Distance: $totalDistance"

# Part 2: Calculate similarity score
$similarityScore = Get-SimilarityScore -leftList $leftList -rightList $rightList
Write-Output "Part 2 - Similarity Score: $similarityScore"
