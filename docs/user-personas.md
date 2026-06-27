# Aperture — User Personas

**Version:** 1.0  
**Date:** 2026-06-27

---

## Persona 1 — Arjun Sharma (New Prospect)

**Age:** 28  
**Occupation:** Software Engineer, Infosys Pune  
**Income:** ₹1.2 LPA  
**Location:** Pune, Maharashtra  
**Tech Literacy:** High  
**Banking:** Has HDFC savings account, no SBI relationship  

### Scenario
Arjun is looking to buy his first flat in Pune. He browses the SBI website for home loan rates, spends 7 minutes on the home loan page, uses the EMI calculator, and then closes the tab.

### Goals
- Find the best home loan rate
- Understand eligibility quickly
- Complete application without visiting a branch

### Pain Points
- Doesn't know his exact eligibility
- Finds bank forms confusing and long
- Worried about submitting wrong documents
- Has previously abandoned applications at document upload stage

### Aperture Touchpoints
1. Intent Agent detects his high-intent session
2. Qualification Agent scores him (income ₹1.2L, age 28, no existing SBI relationship)
3. Personalisation Agent crafts: "You may be eligible for ₹60L home loan at 8.4%"
4. Conversation Agent guides him through the application
5. Document Agent extracts his PAN and salary slip
6. Consent captured, application submitted

---

## Persona 2 — Meena Pillai (Existing SBI Customer – Cross-Sell)

**Age:** 42  
**Occupation:** School Teacher, Government of Kerala  
**Income:** ₹65,000/month  
**Location:** Thiruvananthapuram, Kerala  
**Tech Literacy:** Medium  
**Banking:** SBI savings account (12 years), no loan product  
**Language:** Prefers Malayalam  

### Scenario
Meena logs into YONO to check her balance. She has been browsing the mutual funds section for the past 3 sessions. Aperture detects this cross-sell opportunity.

### Goals
- Grow her savings for retirement
- Wants a safe, trusted product from SBI
- Needs explanation in her language

### Pain Points
- Doesn't understand financial jargon
- Worried about losing her savings
- Prefers to speak in Malayalam

### Aperture Touchpoints
1. Intent Agent detects repeated mutual fund exploration in YONO
2. Qualification Agent uses her existing SBI profile (salary credit, account age, no existing investment product)
3. Personalisation Agent generates Malayalam message
4. Conversation Agent guides her in Malayalam through SBI mutual fund onboarding
5. Pre-filled application from existing SBI KYC

---

## Persona 3 — Rohit Verma (RM – Relationship Manager)

**Age:** 35  
**Role:** Personal Banking Relationship Manager, SBI Mumbai Branch  
**Portfolio:** 200+ HNI/mass-affluent customers  
**Daily Task:** Follow up on leads, cross-sell products, handle escalations

### Goals
- Quickly identify which prospects are worth calling today
- See what product each customer is most likely to take
- Reduce time spent on cold calls
- Get context before a customer meeting

### Pain Points
- No intelligence on which leads are hot vs cold
- Has to manually search CBS for customer data
- No conversational context from digital touchpoints
- No next-best-action guidance

### Aperture Touchpoints
1. RM Dashboard shows prioritised lead list (hot → warm → cold)
2. Each lead card shows: intent signals, qualification score, product recommendation, conversation summary
3. Next Best Action: "Call Arjun today — home loan enquiry, income verified"
4. After call: RM logs outcome, system updates prospect record

---

## Persona 4 — Kavitha Rajan (Branch Employee)

**Age:** 29  
**Role:** Customer Service Officer, SBI Branch Hyderabad  
**Daily Task:** Walk-in customer assistance, account opening, document collection

### Goals
- Open accounts faster with minimal manual data entry
- Guide customers who are confused about digital process
- Reduce document rejection

### Pain Points
- Customers arrive with incomplete documents
- Form filling is slow and error-prone
- No visibility into customer's digital journey before branch visit

### Aperture Touchpoints
1. Branch tablet shows customer's Aperture profile when customer arrives with appointment
2. Pre-filled application reduces data entry
3. Document Agent has already extracted fields — branch employee only verifies
4. Conversation context from digital onboarding visible to branch staff

---

## Persona 5 — Vikram Nair (Compliance Officer)

**Age:** 48  
**Role:** Chief Compliance Officer, SBI Digital Banking Division  
**Responsibility:** Ensure RBI, DPDP Act, and internal policy compliance

### Goals
- Verify every AI decision is explainable and auditable
- Ensure consent was properly obtained before data use
- Flag any AI output that violates policy
- Generate compliance reports for RBI inspection

### Pain Points
- AI systems are black boxes — no visibility into decisions
- Consent records scattered across systems
- No unified audit trail for digital acquisitions
- Manual compliance checks don't scale

### Aperture Touchpoints
1. Compliance Dashboard: consent health, audit flags, policy violations
2. Explainability panel for every AI decision
3. Ability to override any AI decision with reason
4. Export full audit trail as PDF for RBI submission

---

## Persona 6 — Priya Das (Operations Team)

**Age:** 32  
**Role:** Digital Operations Manager, SBI Digital Hub  
**Responsibility:** Monitor acquisition funnel, resolve operational issues, track KPIs

### Goals
- Track daily/weekly acquisition funnel metrics
- Identify where prospects drop off
- Ensure SLA on application processing
- Escalation management

### Pain Points
- No real-time visibility into digital funnel
- Drop-off analysis requires manual data pulls
- No single view of customer across digital channels

### Aperture Touchpoints
1. Operations Dashboard: funnel visualisation, drop-off analysis, conversion rates
2. Application processing SLA tracker
3. Escalation queue for applications requiring human review
4. Volume trends and performance KPIs
