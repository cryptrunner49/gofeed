FROM node:23.7.0-bullseye

WORKDIR /app

COPY . .

RUN npm install -g pnpm

RUN pnpm install

EXPOSE 80
CMD ["pnpm", "dev", "--host", "0.0.0.0", "--port", "80"]
