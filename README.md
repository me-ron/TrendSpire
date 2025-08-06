# 📈 TrendSpire

TrendSpire is a full-stack project that tracks and displays the **Top-K Trending Posts** in real-time using Redis and PostgreSQL. It’s designed to help understand how scalable systems handle heavy hitters and frequent item queries in real-time scenarios.

---

## 🧱 Tech Stack

### Backend
- **Go + Gin**
- **PostgreSQL**
- **Redis**
- **Docker**

### Frontend
- **React** 
---

## 🚀 Features

- Create and like posts
- Track Top-K most liked posts in real-time
- Redis-backed ranking system with sorted sets (`ZINCRBY`, `ZREVRANGE`)
- RESTful API design with Gin
- Sync between Redis and Postgres for reliability
- Fully Dockerized for easy local development

---
