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

**Principles for efficient interviewing:**
- Don't ask about things that can be learned from existing documents
- Ask by presenting specific options
- Proactively check perspectives the user hasn't mentioned

**Example: Efficient interviewing**

User input: "We want to improve development productivity"

❌ Bad question: "What kind of productivity improvement?" (too broad)

✅ Good questions:
- "Among lead time reduction, deployment frequency improvement, and incident recovery time reduction, which would have the most impact?"
- "What is your current deployment frequency? Do you have a target value in mind?"

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
- Measurement infrastructure availability for each KR (→ A/B/C classification, described later)
- If there are 4 or more Objective candidates, narrow down to 3 or fewer through prioritization discussion

**Convergence signal**: Move to the next step when the user responds with "this is roughly right" regarding the direction of Objectives and KRs.

#### 7. Self-review and exit Plan mode

Self-review the plan file and request user approval with ExitPlanMode.

**Tool**: ExitPlanMode

### [OKR Creation Phase] — Normal mode (approval gate)

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

- **Write as outcomes**: Not "release XX" but "XX improves YY"
- **Current value and target value are mandatory**: Don't set goals without a baseline (for measurement infrastructure rank B, follow the rules described later)
- **Measurement method and data source are mandatory**: Confirm "can we measure this right now"
- **Don't include numbers in Objectives**: Numbers belong in KRs

**Tool**: Write, Edit, AskUserQuestion

#### 9. Rewrite in your own words

After completing the OKR draft, ask:

**"Reading this OKR back, does it feel like it's written in your own words? Are there any expressions you'd like to rephrase?"**

- Replace Claude's proposed expressions with the user's own words
- Change jargon and framework terminology to phrases you use daily
- If no changes are needed, proceed as-is

**Tool**: AskUserQuestion, Edit

#### 10. Completion check

Conduct a self-review of the OKR document (using the review perspectives described later), then perform the following **failure scenario check** and present the results to the user.

**Failure Scenario Check:**

"If this OKR were to fail, the following patterns are conceivable:"

- {Pattern 1}: {description}
- {Pattern 2}: {description}
- ...

**Completion criteria**: The user has acknowledged the risks and decided to "proceed."

⛔ **Do not mark as complete without user approval**

## Measurement Infrastructure A/B/C Classification

Classify the measurement infrastructure status when KR candidates emerge.

| Rank | Status | Action |
|------|--------|--------|
| 🟢 **A** | Data source already exists, measurable immediately | Adopt as KR as-is |
| 🟡 **B** | No data source, or measurement setup is not in place | Note "Complete baseline measurement by Q1" as a prerequisite. Record current value as `[Q1 measurement planned]` and set a provisional target value |
| 🔴 **C** | Measurement is not realistic | Replace with a different proxy metric or exclude from KR candidates |

If there are 3 or more rank B items, propose making measurement infrastructure development itself one of the KRs.

## KR Quantification Guide

When users express goals in qualitative terms, propose an appropriate technique from the following and quantify with mutual agreement.

| Technique | When to Apply | Example |
|-----------|--------------|---------|
| Direct metric | A measurable target exists | Revenue, user count, response time |
| Proxy metric | Direct measurement is difficult | "Improve work environment" → eNPS score |
| Milestone completion rate | Project completion type | "Release new feature" → X out of 4 features completed |
| Frequency / count | Behavior / process improvement | "Activate knowledge sharing" → Hold X study sessions per month |
| Time reduction | Efficiency type | "Speed up deployment" → Lead time from X days to Y days |
| Percentage / ratio | Quality improvement type | "Reduce bugs" → Defect rate below X% |
| Satisfaction score | Subjective quality | "Improve customer experience" → NPS from X to Y |
| Binary judgment | Clear achieved/not achieved | "Revamp auth infrastructure" → All services running on new infrastructure |

**Quantification principles:**
- Include current value (baseline) and target value for all KRs
- Include measurement method and data source alongside
- If measurement infrastructure doesn't exist, handle according to A/B/C classification
- When struggling to choose development-related metrics, refer to the [SPACE Framework](space-framework.md)

## Outcome Conversion Guide

When a KR candidate is an output (deliverable), convert it to an outcome (result).

**Example: Output → Outcome conversion**

| Output (❌) | Outcome (✅) | Conversion rationale |
|------------|-------------|---------------------|
| Build a CI/CD pipeline | Increase deployment frequency from weekly to daily | The purpose of building the pipeline is faster deployment |
| Create a monitoring dashboard | Reduce time from incident detection to response start from X min to Y min | The purpose of the dashboard is early incident detection |
| Improve documentation | Reduce new member onboarding period from X weeks to Y weeks | The purpose of documentation is to support ramp-up |
| Achieve 80% test coverage | Reduce post-release defect rate from X% to Y% | Coverage is a means; quality is the goal |

**Conversion prompt**: "If you achieve that, what changes for you or your customers?"

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
