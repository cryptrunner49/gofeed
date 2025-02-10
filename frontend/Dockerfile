FROM node:23.7.0-bullseye

WORKDIR /app

COPY . .

RUN npm install -g bun

RUN bun install

EXPOSE 80
CMD ["bun", "run", "dev", "--host", "0.0.0.0", "--port", "80"]
