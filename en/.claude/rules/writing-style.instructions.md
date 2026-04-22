# Writing Style Guidelines

## Concise Writing Style

- Only necessary and sufficient information. Short sentences, clear terminology, use bullet points
- Avoid preambles, excessive politeness, and unnecessary confirmations
- Avoid over-explanation (assume the reader already knows or can look up the information)
- Provide timeline/effort estimates only when explicitly requested
- Do not add lengthy explanations to self-evident code changes
- Do not proactively explain things the user did not ask about
- Be honest about uncertainties

## Code Conciseness Principles

- Achieve the goal with minimal changes
- Do not abstract one-time-use logic
- Do not add options or settings that are not needed now
- Seven similar lines of code are better than a premature abstraction
- Do not design for hypothetical future requirements

## Style

- **File names**: lowercase-with-hyphens (kebab-case). Exceptions: conventionally uppercase files like README.md, CLAUDE.md
- **Diagrams**: Prefer mermaid for sequence diagrams, state diagrams, and flowcharts
- **Tone**: Use a polite, readable style appropriate for technical articles
- **Emoji**: Limit to key emphasis points, easy-to-miss items, and enumerated lists where items are hard to distinguish
- **Tables**: Keep cell content short (1-2 sentences max). Write longer explanations outside the table
- **Code blocks**: One subject per block. Separate different files and Before/After into different blocks
- Do not write lengthy prose for what code examples can explain sufficiently

## Terminology

- Do not call tool names (Docker, React, etc.) "technology"
- Describe the underlying methods and principles as "technology", and reference tools as their implementations

## Review

After writing or implementing, perform self-review per [self-review.instructions.md](self-review.instructions.md).
