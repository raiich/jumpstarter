---
name: brainstorm
description: Multi-perspective brainstorming combining critical and constructive analysis. Runs parallel subagents with diverse thinking frameworks across multiple rounds.
disable-model-invocation: true
---

# Brainstorm

A skill that validates ideas from multiple angles using diverse thinking frameworks.
Launches subagents in parallel and structures results through up to 3 rounds of iterative discussion.

## Flow Overview

```
1-4: Prep → R1: Independent Analysis (diverge) → R2: Cross-Review (intersect) → R3: Convergence (conditional) → Integration
```

## Flow

### 1. Confirm the Topic

Confirm the topic provided by the user: `$ARGUMENTS`

If the topic is unclear, use AskUserQuestion to clarify.

### 2. Select Subagents

Select **2–3** subagents to launch based on the nature of the topic.
Not all need to be launched every time.

**Selection principle**: Include both critical and constructive perspectives.

| Subagent | Perspective | Selection Criteria |
|---|---|---|
| `critical-analyst` | Critical | Almost always useful. Counterarguments and failure prediction |
| `framework-analyst` | Exploratory | When you want to understand problem structure. Analogies, problem patterns, naive questions |
| `constructive-analyst` | Constructive | Opportunity discovery, strength amplification. When you don't want to stop at criticism |
| `stakeholder-sim` | Multi-faceted | When multiple stakeholders are involved, or for market-facing products |

Report selection rationale together with results (proceed directly to launch without user confirmation here).

### 3. Select Random Constraints

When launching `stakeholder-sim`, intuitively pick **2** from the "Random Constraints" candidates in [stakeholder-sim.md](../../agents/stakeholder-sim.md).

- Avoid repeating the same combination as last time
- Prefer surprising combinations over those too closely related to the topic
- You may also invent constraints not on this list

### 4. Gather Context (if needed)

Only if the topic is related to the codebase, briefly review relevant files and existing designs.

**Tools**: Read, Grep, Glob

### 5. Round 1: Independent Analysis (Diverge)

Launch selected subagents **in parallel** (do not run sequentially).

Provide each subagent with:

```
Topic: [User's topic]
Context: [Collected relevant information, if any]
```

For `stakeholder-sim`, also add random constraints:
```
Random constraints: [Constraints selected in step 3]
```

**Tools**: Agent

### 6. Round 2: Cross-Review (Intersect)

Re-launch the **same agents** from R1 **in parallel**.
Pass each agent the R1 outputs from all agents (including their own) and instruct cross-review.

#### R2 Prompt Structure

```
Topic: [Topic]
Context: [Original context]

--- Round 1 Analysis Results from All Agents ---

## [Agent Name A]'s Analysis:
[R1 output]

## [Agent Name B]'s Analysis:
[R1 output]

## [Agent Name C]'s Analysis:
[R1 output]

---

Cross-review instructions:
Based on the other agents' analyses, from your specialized perspective:
1. Point out overlooked aspects
2. Refute or reinforce other analyses (with evidence)
3. New insights emerging from combining multiple analyses

Important: Do not repeat your own Round 1 analysis. Report only new findings.
```

#### Summarizing R1 Output

If the total R1 output is lengthy, summarize each agent's output to approximately 5–7 key points before passing to R2.

**Tools**: Agent

### 7. Round 3: Judgment and Convergence (Conditional)

Evaluate R2 output and execute convergence **only if any of the following apply**.

| Trigger | Criteria |
|---|---|
| Significant contradiction | Agents reached opposite conclusions on the same point in R2 |
| Major new insight | A substantially new idea or perspective emerged in R2 that wasn't in R1 |
| Unresolved trade-off | Risk vs. opportunity or stakeholder conflicts remain unresolved |

**If triggered**: The orchestrator performs the following (no subagent launch):
- Resolve contradictions or clarify conditions for coexistence
- Identify points of agreement
- Prioritize

**If not triggered**: Skip to step 8.

### 8. Integrate and Structure Results

Reorganize outputs from all rounds (R1 + R2 + R3) **along the reader's decision flow**.
Do not arrange by technique; structure chapters by "what the reader should decide and in what order."
Merge duplicates. Prioritize content deepened or reinforced in R2.
Add a source tag `(from: agent-name / technique)` to each item to preserve traceability.

**Important**: This is an AI-generated analysis. Use it in combination with feedback from actual humans.

#### Output Format

Include only sections that have actual content. Omit empty sections.
Top sections (TL;DR, Next Actions, Key Tensions) must be synthesized, not raw R1 output pasted in.

```markdown
## Brainstorm Results: [Topic]

### TL;DR
3–5 lines summarizing the core. Only decision-relevant conclusions.

### Next Actions
List in execution order. Note verification cost.
1. [Top priority action] — Verification cost: Low/Med/High
2. ...
3. ...

### Key Tensions
Points where agents disagreed, or unresolved trade-offs. Primarily reflects R2 / R3 output.
- Tension: [Description]
  - Position A: [Claim] (from: [source])
  - Position B: [Claim] (from: [source])
  - Deciding condition: [What determines which to pick]

### What to Validate
Merge unverified assumptions and questions to ask.
- [ ] Assumption: [Description] — Verification method: [How to confirm] (from: [source])
- [ ] Question: [Description] — Ask: [Who] (from: [source])

### Risks and Opportunities
- Risk: [Description] (Probability: High/Med/Low, Impact: High/Med/Low) — Prevention: [Explanation] / Early warning: [Signal] (from: [source])
- Opportunity: [Description] — Evidence: [Explanation] (from: [source])

### Alternative Approaches
- Option: [Summary] — Pros/Cons: [Explanation] (from: [source])

### Appendix: Source Notes
Material from individual techniques that didn't fit into the upper sections, or that is worth preserving as-is.
- **Problem Structure Pattern**: [Name] → [Application to this topic]
- **Analogy**: [Domain]'s [Case] → [Transfer]
- **Naive Question**: [Description] → [Why it matters]
- **TRIZ Principle**: [Principle name] → [Insight]

### Analysis Process
- Rounds: [2 or 3]
- Participating agents: [List]
- Selection rationale: [Brief]
- R3 judgment: [Trigger name / Not needed]
```
