# Aperture Frontend

Next.js 14 (App Router) + TypeScript + Tailwind CSS + shadcn/ui

## Structure

```
frontend/src/
├── app/
│   ├── (auth)/          Login, register pages
│   ├── (customer)/      Customer dashboard, chat, consent, documents, apply
│   ├── (rm)/            RM pipeline, lead cards, next best action
│   ├── (compliance)/    Consent health, audit flags, overrides
│   ├── (operations)/    Funnel analytics, conversion charts
│   └── (admin)/         User management, agent config, system health
├── components/
│   ├── ui/              shadcn/ui base components
│   ├── shared/          Navbar, Sidebar, ChatWidget, ScoreCard, etc.
│   └── charts/          Recharts wrappers for dashboards
├── features/            Feature-scoped components + hooks
├── hooks/               Custom React hooks
├── services/            API client functions (React Query)
├── store/               Zustand global state
├── types/               TypeScript type definitions
├── styles/              Global CSS, Tailwind config
└── lib/                 Utility functions, constants
```

## Running locally

```bash
cd frontend
npm install
npm run dev
# App starts on :3000
```
