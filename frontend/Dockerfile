FROM node:8.16.2

RUN npm install -g @vue/cli
ENV BACKEND_URL 'http://goservice:8080/'
#COPY src .
#COPY public .
#COPY package.json .
#COPY package-lock.json .
#COPY babel.config.js .
COPY . .

EXPOSE 8080
ENTRYPOINT ["npm", "run", "serve"]