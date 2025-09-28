# CS-OPN

CS:OPN is a CS2 case opening simulator where users can open up cases and see what they can pull and add those items into their own inventories. Giving users the feeling dopamine of pull golds from cases, without having to spend a ton of money

## 🚀 Tech Stack

### Frontend

- [Next.js](https://nextjs.org/) (App Router, React Server Components)
- [Tailwind CSS](https://tailwindcss.com/) (utility-first styling)
- [shadcn/ui](https://ui.shadcn.com/) (prebuilt, accessible components)
- [Framer Motion](https://www.framer.com/motion/) (smooth animations for case opening)
- [React Query](https://tanstack.com/query) (API calls + caching)
- Deployment: [Vercel](https://vercel.com/)

### Backend

- [Go](https://go.dev/) (fast, lightweight backend)
- Web framework: [Fiber](https://gofiber.io/) or [Chi](https://github.com/go-chi/chi)
- ORM / DB Layer: [GORM](https://gorm.io/) or [SQLC](https://sqlc.dev/)
- Database: PostgreSQL (scalable)
- Authentication: JWT-based (optional)
- Deployment: [Railway](https://railway.app/), [Render](https://render.com/)

## 📂 Project Structure

(coming soon)

## 🔑 Core Features

- 🎨 **Case Opening Animation** – smooth roll animations using Framer Motion
- 📦 **Item Pool & Rarity System** – define skins, rarities, and probabilities in DB
- 📊 **Inventory Tracking** – (optional) user keeps opened items
- ⚙️ **API Endpoints**:
  - `POST /api/open-case` → simulate case opening
  - `GET /api/inventory/:userId` → fetch player inventory
