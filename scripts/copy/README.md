# copy-to-pods

A utility script for copying files to Kubernetes pods that match a specified regex pattern.

## Description

This tool allows you to easily copy a file to multiple Kubernetes pods that match a given regular expression pattern. It also ensures that the target directory exists in each pod before copying, creating it if necessary.

## Usage

```bash
./copy-to-pods -p <pod_pattern> -s <source_file> -d <destination_path>
```

### Required Flags

- `-p, --pattern`: Regex pattern to filter pods by name
- `-s, --source`: Path of the source file to copy
- `-d, --destination`: Destination file path in the pod

## Example

```bash
# Copy config.yaml to all pods whose names contain "backend"
./copy-to-pods -p "backend" -s "./config.yaml" -d "/app/config/config.yaml"
```

## Requirements

- kubectl must be configured and accessible in your PATH
- You must have appropriate permissions to exec into pods and copy files

## How it works

1. Filters available Kubernetes pods based on the provided regex pattern
2. For each matching pod:
   - Creates the target directory if it doesn't exist
   - Copies the source file to the specified destination
   - Provides feedback on successful operations

If any step fails, the script will exit with an error message providing details about the failure.