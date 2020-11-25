FROM node:14.8.0-alpine

WORKDIR /app

COPY package.json yarn.lock ./

RUN yarn 

COPY . .

RUN yarn install

CMD ["yarn", "start"]

