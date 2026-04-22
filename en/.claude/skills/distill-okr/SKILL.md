---
name: distill-okr
description: Craft thoughtful annual OKRs through dialogue. Includes multi-perspective interviewing, quantification support, and built-in failure prevention.
allowed-tools: Read, Grep, Glob, Write, Edit, AskUserQuestion, EnterPlanMode, ExitPlanMode
effort: high
---

# Distill OKR

A skill that distills OKRs into truly reliable metrics through dialogue.
Rather than filling in templates, it digs deep into challenges, mission, and expectations to forge OKRs that boost motivation and engagement.

## Assumptions

- Users may not provide all information upfront, so proactively interview them
- Maximum **3 Objectives**. Force focus
- All Key Results **must be quantified**. Even qualitative goals must be converted to numbers using proxy metrics, etc.
- KRs must describe outcomes (results), not outputs (deliverables)
- Claude proposes structure. **The final wording belongs to the user**

## Save Location

- `.local/docs/okr/{year}.md` (e.g., `.local/docs/okr/2026.md`)

## Flow

### [Interview & Draft Phase] — Plan mode

Use incomplete drafts as a medium for dialogue. Rather than bombarding with questions, show a rough OKR early to elicit feedback.

#### 1. User: Enter fiscal year

#### 2. Enter Plan mode

**Tool**: EnterPlanMode

#### 3. Review existing documents

Investigate under `.local/docs/` and check for:
- Mission definition
- Previous year's OKR (under `.local/docs/okr/`)
- Other related documents

**Tool**: Read, Glob, Grep

#### 4. Initial interview (minimal)

Ask only the following 3 points first. Deeper exploration happens after presenting the draft.

**Tool**: AskUserQuestion

**What to ask:**
- Your mission (reason for existence) and the ideal state one year from now
- The biggest challenges you currently face (up to 3)
- If previous year's OKR exists: achievement level and causes of shortfalls

Hearing principles and Good/Bad examples: see [../../guidelines/processes/hearing.md](../../guidelines/processes/hearing.md).

#### 5. Present a rough OKR draft

Based on the initial interview and existing documents, write an **incomplete OKR draft** to the plan file and show it to the user.

⚠️ **Important**: Don't aim for perfection. Mark insufficient or unclear areas with `[❓ needs confirmation]` and present as-is. The purpose of the draft is to elicit the user's reaction.

```markdown
# OKR Draft (v1)

## Deliverable
- Type: OKR document (markdown)
- Save location: .local/docs/okr/{year}.md

## Objective Candidates

### O1: {candidate}
- KR candidate: {quantitative metric} [❓ needs confirmation: current value / target value]
- KR candidate: {quantitative metric} [❓ needs confirmation: measurement infrastructure availability]

### O2: {candidate}
- ...

## Missing Information
- [❓ needs confirmation] Stakeholder expectations (leadership, etc.)
- [❓ needs confirmation] Resources & constraints (headcount, budget, technical)
- [❓ needs confirmation] {other gaps detected by Claude}
```

**Tool**: AskUserQuestion

After showing the draft, ask: **"Is this the right direction? What's different or missing?"**

Also include a question that encourages disagreement: **"Is there anything here that feels wrong to you?"**

#### 6. Iterative refinement (repeat)

Refine the draft based on user feedback. Ask targeted questions about missing information (stakeholder expectations, resource constraints, measurement infrastructure availability, etc.).

**Perspectives to confirm during iterations:**
- Contradictions with stakeholder expectations
- Resources & constraints (headcount, budget, technical aspects, major events during the fiscal year)
- Measurement infrastructure availability for each KR (→ A/B/C: [quantification.md](quantification.md))
- If there are 4 or more Objective candidates, narrow down to 3 or fewer through prioritization discussion

**Convergence signal**: Move to the next step when the user responds with "this is roughly right" regarding the direction of Objectives and KRs.

#### 7. Self-review and exit Plan mode

Self-review the plan file and request user approval with ExitPlanMode.

**Tool**: ExitPlanMode

### [OKR Creation Phase] — Normal mode (autonomous after ExitPlanMode approval)

After obtaining direction approval via ExitPlanMode, proceed to completion without asking for further approvals.

#### 8. Create OKR document

Based on the refined draft in the plan file, create the formal OKR document.

**Save location**: `.local/docs/okr/{year}.md`

**Template:**

```markdown
# OKR Fiscal Year {year}

## Overview

- **Period**: {start} – {end}
- **Created**: {date}
- **Mission**: {summary of mission}

## Objective 1: {qualitative and ambitious goal}

{Rationale for setting this Objective. Connection to challenges and stakeholder expectations.}

| # | Key Result | Current Value | Target Value | Measurement Method | Data Source | Measurement Infrastructure |
|---|-----------|---------------|--------------|-------------------|-------------|---------------------------|
| 1 | {KR description} | {value} | {target} | {method} | {source} | {A/B/C} |
| 2 | {KR description} | {value} | {target} | {method} | {source} | {A/B/C} |
| 3 | {KR description} | {value} | {target} | {method} | {source} | {A/B/C} |

## Objective 2: {qualitative and ambitious goal}

(Same structure)

## Objective 3: {qualitative and ambitious goal}

(Same structure. Maximum 3 Objectives)

## Assumptions & Constraints

- {Resource constraints and risks identified during interviews}

## Quantification Rationale

Conversion rationale for qualitative goals that were quantified:
- {KR}: {why this metric and value were chosen}

## Objective Candidates Deferred to Out of Scope

| Candidate | Reason for Deferral | Priority for Next Year |
|-----------|--------------------|-----------------------|
| {candidate} | {reason} | {High/Medium/Low} |

## Appendix: Operations Guide

- **Review cadence**: Monthly progress check recommended
- **Review method**: {specific review procedures based on each KR's measurement method}
- **Mid-term review**: Re-evaluate target value validity at the half-year mark
- **Kickoff**: Internalize the OKRs. Articulate how each Objective connects to your daily actions
```

**KR writing rules:**

- **Write as outcomes**: Not "release XX" but "XX improves YY" ([outcome-vs-output.md](outcome-vs-output.md))
- **Current value, target value, measurement method, and data source are mandatory**: measurement-foundation classification (A/B/C), quantification techniques, and rationale are in [quantification.md](quantification.md)
- **Don't include numbers in Objectives**: Numbers belong in KRs

**Tool**: Write, Edit, AskUserQuestion

#### 9. Rewrite in your own words

After completing the OKR draft, ask:

**"Reading this OKR back, does it feel like it's written in your own words? Are there any expressions you'd like to rephrase?"**

- Replace Claude's proposed expressions with the user's own words
- Change jargon and framework terminology to phrases you use daily
- If no changes are needed, proceed as-is

**Tool**: AskUserQuestion, Edit

#### 10. Completion report

Conduct a self-review of the OKR document (using the review perspectives described later), then perform the following **failure scenario check** and report the results to the user.

**Failure Scenario Check:**

"If this OKR were to fail, the following patterns are conceivable:"

- {Pattern 1}: {description}
- {Pattern 2}: {description}
- ...

Do not request approval. Apply any revision instructions the user provides ([../../guidelines/processes/review-flow.md](../../guidelines/processes/review-flow.md)).

## Self-Review Perspectives

### Interview & Draft Phase
- Were mission, challenges, expectations, and constraints comprehensively covered?
- Was previous year's OKR achievement level and causes of shortfalls confirmed?
- Were contradictions among stakeholders detected?

### OKR Document
- Are Objectives qualitative and ambitious (no numbers included)?
- **Are all KRs outcomes (not outputs)?**
- Are all KRs quantified (current value, target value, measurement method, and data source all present)?
- **Do all KRs have a measurement infrastructure rank (A/B/C)?**
- **For rank B items, is a baseline measurement deadline specified?**
- Does achieving the KRs indicate Objective achievement (alignment)?
- **Are there 3 or fewer Objectives?**
- Is quantification rationale documented?
- **Are deferred Objective candidates recorded?**
- **Is it written in your own words (not Claude-speak)?**
- **Has the failure scenario check been conducted and acknowledged by the user?**

Adjust perspectives based on content. If anything is unclear, interview the user.
