FROM node:23-alpine AS frontend

# Set working directory and copy over all files
WORKDIR /app
COPY . .

# Clean install node modules then build
RUN npm ci
RUN npm run build

# Delete intermediate files no longer needed
RUN rm -rf src/ static/ docker-compose.yml

USER node:node
CMD ["node","build/index.js"]