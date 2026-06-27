# Aperture AI Agents

Python 3.11 + FastAPI + LangChain — 8 independently deployable AI agents

## Agent Ports

| Agent | Port | Responsibility |
|-------|------|---------------|
| intent-agent | 8101 | Session signals → Intent score + category |
| qualification-agent | 8102 | Profile → 5-dimension eligibility score |
| personalisation-agent | 8103 | Profile → Personalised offer message + language |
| conversation-agent | 8104 | Chat turn → Response + field extraction |
| document-agent | 8105 | Document image → OCR extracted fields + confidence |
| compliance-agent | 8106 | Action → Policy check (consent, minimisation) |
| explainability-agent | 8107 | Decision → Human-readable explanation |
| audit-agent | 8108 | Event → Immutable audit log entry |

## Each Agent Structure

```
agents/{agent-name}/
├── config/         agent_config.json — thresholds, policy refs, model config
├── prompts/        prompt templates (versioned YAML)
├── service/        Core agent logic (chains, tools, memory)
├── tools/          LangChain tool definitions
├── memory/         Memory class implementations
└── tests/          Unit + integration tests
```

## LLM

All agents use Claude claude-sonnet-4-6 via the Anthropic SDK. The `ANTHROPIC_API_KEY` must be set in the environment.

## Running an agent

```bash
cd agents/intent-agent
pip install -r requirements.txt
uvicorn service.main:app --port 8101 --reload
```
