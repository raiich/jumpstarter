# Efficient Hearing Principles

Principles for extracting requirements and information from users. Avoid broad questions; ask pinpointed questions based on prior investigation.

## Principles

- Do not ask about things that can be understood from code or existing docs
- Present specific options informed by constraints and patterns identified in investigation
- Proactively confirm aspects the user has not mentioned (edge cases, consistency with existing features, stakeholder expectations, resource constraints)
- Pair positive questions with negative ones ("Which of these feels *wrong* to you?")

## Examples

### Case 1: Hearing after codebase investigation

User input: "I want to add a notification feature"

❌ Bad: "What kind of notification feature?" (too broad)

✅ Good (options informed by investigation):
- "Do you envision event-driven using the existing EventBus (events/bus.go), or direct invocation?"
- "Are there notification targets beyond Slack webhook (slack_url already exists in config)?"

### Case 2: Hearing on qualitative goals

User input: "I want to improve development productivity"

❌ Bad: "What kind of productivity improvement?" (too broad)

✅ Good:
- "Among lead-time reduction, deploy frequency, and MTTR improvement, which has the biggest impact?"
- "What is the current deploy frequency? Do you have a target value in mind?"
