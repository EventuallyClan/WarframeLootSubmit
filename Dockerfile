FROM node

ENV PORT 8080

EXPOSE 8080

COPY package.json package.json
RUN npm install

COPY . .
RUN npm run build

CMD ["node", "dist/"]

