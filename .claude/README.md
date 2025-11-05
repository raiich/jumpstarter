# Claude Code Conversation Logs

This directory contains settings for recording conversations with Claude Code.

## Configuration

The hooks defined in `settings.json` automatically log the following:

- **User Prompts**: Questions and instructions you send to Claude Code
- **Tool Executions**: Tools used by Claude Code (file reads, edits, command executions, etc.)

## Log Files

Conversation logs are saved to `.claude/logs/conversation.log`.

```bash
# View logs
cat .claude/logs/conversation.log

# Monitor logs in real-time
tail -f .claude/logs/conversation.log

# Search logs for a specific date
grep "2025-11-01" .claude/logs/conversation.log
```

## How to Use Logs

### 1. Search Past Conversations
```bash
# Search for conversations containing a specific keyword
grep -A 5 "keyword" .claude/logs/conversation.log
```

### 2. Create Daily Reports
```bash
# Output today's conversations to a file
grep "$(date '+%Y-%m-%d')" .claude/logs/conversation.log > today.log
```

### 3. Conversation Statistics
```bash
# Count number of user prompts
grep "USER PROMPT" .claude/logs/conversation.log | wc -l

# Check types of tools used
grep "TOOL:" .claude/logs/conversation.log | sort | uniq -c
```

## Important Notes

- Log files (`.claude/logs/`) are added to `.gitignore` and will not be committed to Git
- Be careful when sharing as they may contain personal conversation history
- Log files are not automatically rotated, so regular cleanup is recommended

## Customization

You can change the log format and recorded information by editing `settings.json`.
For details, see the [Claude Code Hooks documentation](https://docs.claude.com/en/docs/claude-code/hooks).

### About Settings Files

- `.claude/settings.json`: Project-specific settings (used in this repository)
- `~/.claude/settings.json`: Global settings (applied to all projects)
- `.claude/settings.local.json`: Local settings (personal settings not committed to Git)
