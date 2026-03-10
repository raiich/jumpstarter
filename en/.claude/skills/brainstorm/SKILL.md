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

### 3. Generate Random Constraints

Intuitively pick **2** from the list below. Avoid repeating the same combination as last time.
Prefer surprising combinations over those too closely related to the topic.
You may also invent constraints not on this list.

- What if the budget were 1/10 of the current one?
- What if the target audience had the opposite attributes?
- What if this technology became unavailable?
- What if a competitor launched the same thing in 3 months?
- What if the scale were 100x larger?
- What if the team were cut in half?
- What if the deadline were doubled?
- What if regulations changed?
- What if a key customer left?
- What if the core technology's patent expired?

Pass these constraints to `stakeholder-sim` if it is launched.

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
Random constraints: [Constraints generated in step 3]
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

Integrate outputs from all rounds (R1 + R2 + R3) and structure as follows.
Merge overlapping findings to maximize information density.
Prioritize content deepened or reinforced in R2.

**Important**: This is an AI-generated analysis. Use it in combination with feedback from actual humans.

#### Output Format

Include only sections that correspond to actual subagent output.
Omit empty sections.

```markdown
## Brainstorm Results: [Topic]

### Analysis Process
- **Rounds**: [2 or 3]
- **Participating agents**: [Agent name list]
- **Selection rationale**: [Brief reason]
- **R3 judgment**: [Trigger if executed / "Not needed" if skipped]

### Unverified Assumptions
Things this idea implicitly assumes. Verification needed.
- [ ] Assumption 1: [Description] — Verification method: [How to confirm]

### Risk List (by priority)
- [ ] Risk 1: [Description] (Probability: High/Med/Low, Impact: High/Med/Low)
  - Preventive measure: [Explanation]
  - Early warning indicator: [Signal]

### Problem Structure Analysis
Known patterns this problem maps to, and their known solutions.
- Pattern: [Name] — Application to this topic: [Explanation]

### Hints from Analogies
Transferable patterns from structurally similar problems in other domains.
- [Domain]: [Case] → Transfer: [Explanation]

### Beginner's Questions
Naive questions that experts tend to overlook.
- [ ] Question: [Description] — Why it matters: [Explanation]

### TRIZ Insights
Ideas derived from inventive principles.
- [Principle name]: [Concrete insight for the topic]

### Opportunities and Strengths
Points to leverage and hidden opportunities. With evidence.
- Opportunity 1: [Description] — Evidence: [Explanation]

### Questions to Validate
Questions to confirm with actual users, markets, or stakeholders.
- [ ] Question 1: [Description] — Ask: [Who to ask]

### Discovered Alternative Approaches
- Option 1: [Summary] — Source: [Which analysis]

### Recommended Next Actions
1. [Top priority to validate]
2. [Next to validate]
3. [Then validate]
```
