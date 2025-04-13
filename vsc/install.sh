#!/bin/bash

# Get the directory where the script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Get the script's own filename
SCRIPT_NAME="$(basename "$0")"

# Set the destination directory (VS Code user settings path)
DEST_DIR="$HOME/AppData/Roaming/Code/User"

# Find files to be copied (not directories)
FILES_TO_COPY=$(find "$SCRIPT_DIR/Settings" -maxdepth 1 -type f ! -name "$SCRIPT_NAME")

echo ==== SETTINGS ====
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

echo
echo ====== EXTENSIONS ========

if [ -f "$SCRIPT_DIR/extension.txt" ]; then
    echo "The following extensions will be installed"
    while IFS= read -r line; do 
    echo "$line"
    done <"$SCRIPT_DIR/extension.txt"

    read -p "Install extension? [y/N]: " confirm

    if [[ "$confirm" =~ ^[Yy]$ ]]; then
        while IFS= read -r line; do 
            if [[ -n "$line" ]]; then
                echo "Installing: $line"
                code --install-extension "$line"
            fi
        done < "$SCRIPT_DIR/extension.txt"
    else
        echo "Aborted."
    fi

fi
