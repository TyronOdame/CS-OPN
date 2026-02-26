# CS-OPN

CS:OPN is a CS2 case-opening simulator where users can:
- browse and buy cases,
- open them with an animated reel,
- receive skins with weighted drop rates,
- track inventory and transactions.

This project is for simulation/entertainment and is not affiliated with Valve.

## Tech Stack

### Frontend
- Next.js 15 (App Router)
- React 19
- Tailwind CSS
- shadcn/ui + Radix UI
- TypeScript

### Backend
- Go + Gin
- GORM
- PostgreSQL
- JWT authentication

## Project Structure

```text
CS-OPN/
├── frontend/              # Next.js app (UI, pages, components)
├── backend/               # Go API server
│   ├── handlers/          # HTTP route handlers
│   ├── models/            # GORM models
│   ├── middleware/        # Auth middleware
│   ├── database/          # DB connect/migrate/seed helpers
│   └── seed/              # Seed data + image sync logic
└── README.md
```

## Core Features

- Case shop and case purchase flow
- Case opening animation with centered winner marker
- Weighted drop logic by rarity
- Inventory management (including selling items)
- User transaction history
- AI mock price-check endpoint
- Automatic image URL sync from SteamApis for seeded skins/cases

## Prerequisites

- Node.js 18+ (recommended: current LTS)
- npm
- Go (matching `backend/go.mod`)
- PostgreSQL running locally or remotely

## Environment Variables

### Backend (`backend/.env`)

Create `backend/.env`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=csopn
JWT_SECRET=replace_with_secure_secret
PORT=8080
FRONTEND_URL=http://localhost:3000
```

Notes:
- `DB_PASSWORD` and `JWT_SECRET` are required.
- `PORT` defaults to `8080` if not set.

### Frontend (`frontend/.env.local`)

Create `frontend/.env.local`:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080
```

If omitted, frontend defaults to `http://localhost:8080`.

## Local Development

Run backend and frontend in separate terminals.

### 1) Start backend

```bash
cd backend
go run .
```

Backend health check:
- `http://localhost:8080/health`

On startup, backend will:
- connect to PostgreSQL,
- run migrations,
- seed data (when needed),
- sync image URLs for cases/skins.

### 2) Start frontend

```bash
cd frontend
npm install
npm run dev
```

Frontend:
- `http://localhost:3000`

## Frontend Scripts

From `frontend/`:

- `npm run dev` - start local dev server
- `npm run build` - production build
- `npm run start` - run production build
- `npm run lint` - run ESLint

## API Routes (Current)

### Public
- `GET /health`
- `POST /auth/register`
- `POST /auth/login`
- `GET /cases`
- `GET /cases/:id`

### Protected (JWT required)
- `GET /user/profile`
- `PUT /user/profile`
- `POST /cases/:id/buy`
- `POST /cases/:id/open`
- `GET /inventory`
- `POST /inventory/:id/sell`
- `GET /inventory/cases`
- `POST /inventory/cases/:id/open`
- `GET /transactions`
- `POST /ai/price-check`

## Troubleshooting

- `ERR_CONNECTION_REFUSED` from frontend:
  - backend is not running on `localhost:8080`; start `go run .` in `backend/`.
- Broken images:
  - restart backend so image sync runs and updates stored URLs.
- CORS errors:
  - ensure frontend runs on `http://localhost:3000` and backend on `http://localhost:8080`.
