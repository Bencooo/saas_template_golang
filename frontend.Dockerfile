# FROM node:22.0.0-alpine AS builder

# COPY ./frontend. ./
# RUN npm install
# RUN npm run build

# FROM node:22.0.0-alpine AS runner
# COPY --from=builder /app/frontend.dist /app
# ENTRYPOINT ["node /app/index.html"]

# frontend.Dockerfile (à la racine du projet)
FROM node:18-alpine AS builder

WORKDIR /app

COPY . . 
RUN npm install
RUN npm run build

FROM node:18-alpine AS runner

WORKDIR /app

COPY --from=builder /app ./
EXPOSE 8080
CMD ["npm", "start"]
