# First Stage Build: Frontend (Vite + React) into static folder (dist)
FROM node:23 AS frontend-builder

WORKDIR /frontend

COPY frontend/package*.json ./
RUN npm install

COPY frontend .
RUN npm run build

# Second Stage Build: Backend (Go) 
FROM golang:1.24 AS backend-builder

WORKDIR /backend

COPY backend/go.mod backend/go.sum ./
RUN go mod tidy

COPY backend .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o asuka-fan-controller ./cmd/server

# Final Stage: Runtime
FROM alpine:latest

WORKDIR /app

# Install ipmitool and minimal dependencies
RUN apk add --no-cache ipmitool tzdata

# Copy Go binary and static assets
COPY --from=backend-builder /backend/asuka-fan-controller .
COPY --from=frontend-builder /frontend/dist ./static

EXPOSE 8080

CMD ["./asuka-fan-controller"]