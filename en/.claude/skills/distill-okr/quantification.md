# OKR Quantification Guide

Techniques for quantifying Key Results, classification of measurement foundations, and how to document rationale.

## Quantification Techniques

When quantifying a qualitative goal, pick a suitable technique and agree on it with the user.

| Technique | When to use | Example |
|---|---|---|
| Direct metric | Measurable target exists | Revenue, users, response time |
| Proxy metric | Direct measurement is hard | "Improve work environment" → eNPS |
| Milestone completion rate | Project-completion-type | "Release new features" → X of 4 done |
| Frequency / count | Behavior / process improvement | "Enable tech sharing" → X events/month |
| Time reduction | Efficiency | "Faster deploys" → Lead time X days → Y days |
| Ratio / percentage | Quality improvement | "Reduce bugs" → Defect rate ≤ X% |
| Satisfaction score | Subjective quality | "Improve CX" → NPS X → Y |
| Binary outcome | Clear achieved/not | "Auth platform overhaul" → All services on new platform |

### Principles

- Every KR has a baseline (current value) and a target value
- Measurement method and data source are recorded
- Without a measurement foundation, follow the A/B/C classification (below)
- When stuck on development metrics, see [SPACE framework](space-framework.md)
- Write outcomes, not outputs ([outcome-vs-output.md](outcome-vs-output.md))

## Measurement Foundation A/B/C Classification

Classify each KR candidate by the status of the measurement foundation.

| Rank | Status | Action |
|---|---|---|
| 🟢 **A** | Data source exists, measurable now | Adopt as-is |
| 🟡 **B** | No data source / uninstrumented | Add "Complete baseline measurement in Q1" as a prerequisite. Mark current value as `[TBD Q1]` and set a tentative target |
| 🔴 **C** | Measurement is impractical | Replace with a proxy metric, or drop from candidates |

If 3+ KRs are rank B, propose making "building the measurement foundation" itself one KR.

## Rationale

When quantifying a qualitative goal, record the conversion rationale in the OKR document.

```markdown
## Quantification Rationale

- {KR}: {why this metric / number was chosen}
```

Without rationale, "why this number?" cannot be traced, and mid-cycle reviews lose their basis.
