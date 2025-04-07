#!/bin/bash

# Get the directory where the script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Get the script's own filename
SCRIPT_NAME="$(basename "$0")"

# Set the destination directory (VS Code user settings path)
DEST_DIR="$HOME/AppData/Roaming/Code/User"

# Find files to be copied (not directories)
FILES_TO_COPY=$(find "$SCRIPT_DIR" -maxdepth 1 -type f ! -name "$SCRIPT_NAME")

# Dry run preview
echo "The following files will be copied to:"
echo "$DEST_DIR"
echo

for file in $FILES_TO_COPY; do
    echo "  $(basename "$file")"
done

echo
read -p "Proceed with copying? [y/N]: " confirm

if [[ "$confirm" =~ ^[Yy]$ ]]; then
    mkdir -p "$DEST_DIR"
    for file in $FILES_TO_COPY; do
        cp "$file" "$DEST_DIR"
    done
    echo "Files copied successfully."
else
    echo "Aborted."
fi
