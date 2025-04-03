FROM node:22.0.0-alpine AS builder

COPY ./frontend. ./
RUN npm install
RUN npm run build

FROM node:22.0.0-alpine AS runner
COPY --from=builder /app/frontend.dist /app
ENTRYPOINT ["node /app/index.html"]