#!/usr/bin/env bash
# PostToolUse(Write|Edit) hook: flag reconsider-worthy words in the edited
# file against writing-word-check.md. Advisory — exit 2 reports, not a block.
set -euo pipefail

list="$(dirname "${BASH_SOURCE[0]}")/../references/writing-word-check.md"
file=$(jq -r '.tool_input.file_path // empty')

# All written/edited files are checked, incl. code comments. An empty or
# non-regular path (cat would hang on stdin / error on a dir) is not ours.
[ -f "$file" ] || exit 0
# A missing list is a setup error, not a clean pass — report it instead of
# crashing in awk (which set -e would surface as a confusing block).
if [ ! -f "$list" ]; then
  printf 'wording-check: word list (%s) not found.\n' "$list" >&2
  exit 2
fi
content=$(cat "$file")
# Opt-out marker must be the first line, so files that merely quote it as an
# example in their body (e.g. docs teaching this hook) stay checked.
[ "${content%%$'\n'*}" = '<!-- wording-check: skip -->' ] && exit 0

# Each list line is `pattern ::: message`, grouped in ```exact / ```regex blocks.
scan() {  # $1 = block tag, $2 = grep flag (-F | -E)
  awk -v t="$1" '$0=="```"t{b=1;next} b&&$0=="```"{b=0} b' "$list" |
    while IFS= read -r line; do
      pat=${line%% ::: *}
      [ -n "$pat" ] || continue
      if grep -q "$2" -- "$pat" <<<"$content"; then
        printf '  - "%s": %s\n' "$pat" "${line#* ::: }"
      fi
    done
}

hits=$(scan exact -F; scan regex -E)
[ -n "$hits" ] || exit 0
printf 'wording-check (%s): consider these replacements.\n%s\n' "$file" "$hits" >&2
exit 2
