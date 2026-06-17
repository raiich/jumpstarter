#!/usr/bin/env bash
set -euo pipefail

# input: https://code.claude.com/docs/en/statusline
input=$(cat)

MODEL=$(jq -r '.model.display_name // "?"' <<<"$input")
DIR=$(jq -r '.workspace.current_dir // .cwd // "."' <<<"$input")
# Usage (%): prefer used_percentage, else derive from remaining. Treat as 0 before the first API response
USED=$(jq -r '(.context_window.used_percentage // (100 - (.context_window.remaining_percentage // 100))) | floor' <<<"$input")
[ "$USED" -lt 0 ] && USED=0; [ "$USED" -gt 100 ] && USED=100

CYAN='\033[36m'; GREEN='\033[32m'; YELLOW='\033[33m'; RED='\033[31m'; DIM='\033[2m'; RESET='\033[0m'

# Render usage as a width-10 bar (1/8 per cell: fractional remainder drawn with a partial block). Red when high
if   [ "$USED" -ge 90 ]; then BAR_COLOR="$RED"
elif [ "$USED" -ge 70 ]; then BAR_COLOR="$YELLOW"
else BAR_COLOR="$GREEN"; fi
WIDTH=10
BLOCKS=(' ' '▏' '▎' '▍' '▌' '▋' '▊' '▉' '█')   # index 0 = blank, 8 = full block
FULL=$(( USED * WIDTH / 100 ))                  # filled cells
FRAC=$(( (USED * WIDTH % 100) * 8 / 100 ))      # quantize remainder to 0..7
printf -v FILL "%${FULL}s"; BAR="${FILL// /█}"
if [ "$FULL" -lt "$WIDTH" ]; then
  BAR+="${BLOCKS[FRAC]}"
  printf -v PAD "%$(( WIDTH - FULL - 1 ))s"; BAR+="${PAD// /░}"
fi

BRANCH=$(git -C "$DIR" branch --show-current 2>/dev/null || true)
GIT_SEG=""
[ -n "$BRANCH" ] && GIT_SEG=" 🌿 $BRANCH"

# Last-update marker: script run time (updates on events, frozen while idle). Show TZ too so prompt time can be derived from response time - duration
TS=$(TZ=Asia/Tokyo date '+%H:%M %Z')

printf '%b\n' "${CYAN}${MODEL}${RESET}${GIT_SEG} ${DIM}·${RESET} ${BAR_COLOR}${BAR} ${USED}%${RESET} ctx ${DIM}· 🕐 ${TS}${RESET}"
