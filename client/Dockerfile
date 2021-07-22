FROM node:12

ENV PORT 3000

# Create app directory
WORKDIR /next/app

# Installing dependencies
COPY package*.json .
RUN yarn install

# Copying source files
COPY . .

# Building app
RUN yarn build
EXPOSE 3000

# Running the app
CMD ["yarn", "dev"]