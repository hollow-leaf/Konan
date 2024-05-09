FROM node:18-alpine
WORKDIR /
COPY . .
RUN npm install -g pnpm@8.9
RUN pnpm install
ENV NEXT_PUBLIC_PROJECT_ID $project_id
RUN pnpm build
CMD ["pnpm", "dev"]