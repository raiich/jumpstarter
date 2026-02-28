#!/bin/bash
# Logs stdin to .local/claude/conversation.log with a timestamp and header in markdown format
# Usage: echo "content" | log-to-file.sh "HEADER TEXT"

set -euo pipefail

HEADER_TEXT="${1:-}"

if [ -z "$HEADER_TEXT" ]; then
  echo "Error: Header text is required" >&2
  echo "Usage: $0 'HEADER TEXT'" >&2
  exit 1
fi

CONTENT=$(cat)

# Skip if stdin is empty
if [ -z "$CONTENT" ]; then
  exit 0
fi

REPO_ROOT=$(git rev-parse --show-toplevel 2>/dev/null || pwd)
LOG_FILE="$REPO_ROOT/.local/claude/conversation.log"
mkdir -p "$(dirname "$LOG_FILE")"

{
  echo ""
  echo "## [$(date "+%Y-%m-%d %H:%M:%S")] $HEADER_TEXT"
  echo ""
  echo "$CONTENT"
  echo ""
} >> "$LOG_FILE"
