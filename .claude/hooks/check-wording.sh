#!/usr/bin/env bash
# PostToolUse(Write|Edit) hook: flag reconsider-worthy words in the edited file
# against references/writing-word-check.md. Advisory — exit 2 reports, not a block.
set -euo pipefail

list="$(dirname "${BASH_SOURCE[0]}")/../references/writing-word-check.md"
file=$(jq -r '.tool_input.file_path // empty')

# Skip a non-regular path (empty / dir): not a file we check. A missing $list
# is left to fail loudly in awk below, which names the path.
[ -f "$file" ] || exit 0
# Files that quote the flagged words as examples (the word list, the style
# rule) are skipped by path, so those files carry no in-file opt-out marker.
case "$file" in
  *.claude/references/writing-word-check.md|*.claude/rules/writing-style.instructions.md) exit 0 ;;
esac

content=$(<"$file")

# Each list line is `pattern ::: message`, grouped in ```exact / ```regex blocks.
scan() {  # $1 = block tag, $2 = grep flag (-F | -E)
  awk -v t="$1" '$0=="```"t{b=1;next} b&&$0=="```"{b=0} b' "$list" |
    while IFS= read -r line; do
      pat=${line%% ::: *}
      [ -n "$pat" ] || continue
      # `if` (not `&&`) so a non-match returns 0 and set -e never trips here.
      if grep -q "$2" -- "$pat" <<<"$content"; then
        printf '  - "%s": %s\n' "$pat" "${line#* ::: }"
      fi
    done
}

hits=$(scan exact -F; scan regex -E)
[ -n "$hits" ] || exit 0
printf 'wording-check (%s): consider these replacements.\n%s\n' "$file" "$hits" >&2
exit 2
