# 📈 TrendSpire

TrendSpire is a full-stack project that tracks and displays the **Top-K Trending Posts** in real-time using Redis and PostgreSQL. It’s designed to help you understand how scalable systems handle heavy hitters and frequent item queries in real-time scenarios.

---

## 🧱 Tech Stack

### Backend
- **Go + Gin** – High-performance web framework
- **PostgreSQL** – Persistent storage for posts and likes
- **Redis** – Real-time Top-K tracking using Sorted Sets
- **Docker** – Containerized app for easy setup and deployment

### Frontend
- **React** – Simple interface to interact with posts and visualize trending content

---

## 🚀 Features

- Create and like posts
- Track Top-K most liked posts in real-time
- Redis-backed ranking system with sorted sets (`ZINCRBY`, `ZREVRANGE`)
- RESTful API design with Gin
- Sync between Redis and Postgres for reliability
- Fully Dockerized for easy local development

---
