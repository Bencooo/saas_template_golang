# SaaS Template – Golang & Next.js

Projet full-stack SaaS basé sur :
- **Go** (Backend REST API)
- **Next.js** (Frontend)
- **MongoDB** (Base de données)
- **Docker** pour l'environnement de dev

---

## Fonctionnalités actuelles

- Inscription utilisateur (`/register`)
- Connexion utilisateur (`/login`)
- Authentification via JWT
- Route protégée `/profile`
- Frontend responsive (formulaires stylés)
- Dockerisation complète (Mongo, Backend, Frontend)

---

## Démarrage rapide

### 1. Cloner le projet
```bash
git clone https://github.com/Bencooo/saas_template_golang
cd saas_template_golang
```

### 2. Lancer les services
```bash
docker compose up --build
```

### 3. Accéder aux services

| Service     | URL                                              |
|-------------|--------------------------------------------------|
| Frontend    | [http://localhost:8080](http://localhost:8080)   |
| Backend API | [http://localhost:3000](http://localhost:3000)   |
| MongoDB     | mongodb://root:1234@mongo:27017/saas_db?authSource=admin |



## Authentification

- Token **JWT** généré automatiquement lors de la connexion (`/login`)
- Le token est **stocké dans le `localStorage`** côté frontend
- Un **middleware backend** (`jwtMiddleware`) protège les routes sensibles (ex : `/profile`)

## À faire ensuite

- [ ] Implémenter un **CRUD complet** 
- [ ] Utiliser une **goroutine avec channel** 

