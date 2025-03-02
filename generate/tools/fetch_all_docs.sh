#!/bin/bash

# fetch_all_docs.sh - Downloads documentation for all Apple frameworks
# Usage: ./fetch_all_docs.sh [output_dir]

OUTPUT_DIR=${1:-"generate/apidocs"}
FRAMEWORKS=(
    "appkit"
    "foundation"
    "corefoundation" 
    "coregraphics"
    "coredata"
    "coreaudio"
    "coremedia"
    "avfoundation"
    "webkit"
    "metal"
    "cloudkit"
    "gamekit"
    "healthkit"
    "homekit"
    "mapkit"
    "scenekit"
    "security"
    "spritekit"
    "uikit"
    "vision"
)

mkdir -p "$OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR/analysis"

echo "Starting documentation download for ${#FRAMEWORKS[@]} frameworks"
echo "Output directory: $OUTPUT_DIR"
echo "-----------------------------------------"

for framework in "${FRAMEWORKS[@]}"; do
    echo "Fetching documentation for $framework..."
    
    # Call the fetch_apple_docs tool
    go run generate/tools/fetch_apple_docs.go "$framework"
    
    # Wait a bit to be nice to Apple's servers
    sleep 1
done

echo "-----------------------------------------"
echo "Documentation download complete"
echo "Running documentation analyzer..."

# Run the docs analyzer
go run generate/tools/docs_analyzer.go --outdir="$OUTPUT_DIR/analysis"

echo "Analyzing darwinkit coverage..."
# Run the coverage analyzer
go run generate/tools/analyze_darwinkit_coverage.go --apple-docs="$OUTPUT_DIR" --out="$OUTPUT_DIR/analysis/coverage_report.json"

echo "All documentation processing complete"