FROM node:8-buster

WORKDIR /app
RUN apt update && apt install -y ffmpeg libffi-dev libsodium-dev build-essential

COPY . .
RUN npm install

CMD ["node", "./index.js"]
