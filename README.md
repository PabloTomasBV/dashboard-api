# Dashboard-API

# Description
Simple Go API that fetches dashboard data from https://dummyjson.com.

# Requirements
- Go 1.25 (or 1.20+)
- make

# Run locally
1. Clone the project

2. Run:
   ```bash
   make run
   ```
3. You should see:
   ```text
   Listening and serving HTTP on :8080
   ```
4. Try the endpoint:
   ```bash
   curl http://localhost:8080/dashboard/1
   ```
   
# Test:
1. Run
  ```bash
  make test
  ```