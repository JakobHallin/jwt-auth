# JWT-Go

A backend authenticatio built with Go, PostgreSQL, and JWT. Token-based authentication, test coverage, and containerized deployment using Docker and Kubernetes.

---

## Features

- ✅ User signup and login
- 🔐 Password hashing with `bcrypt`
- 🪪 JWT token generation and validation
- 🔒 Protected routes with token-based access control
- 🐳 Docker + docker-compose setup
- ☸️ Kubernetes support for app, database, and test jobs

---

## 📂 Project Structure
```bash
.
├── main.go                 # App entrypoint
├── user.go                # User logic (create, auth)
├── jwt.go                 # JWT creation & validation
├── handlers.go            # HTTP endpoints
├── *_test.go              # Unit & handler tests
├── Dockerfile             # Production app image
├── Dockerfile.test        # Kubernetes test runner
├── docker-compose.yml     # Local dev & test environment
├── k8s/*.yaml                 # Kubernetes manifests
└── README.md              # This file
```
## ⚙️ Getting Started (Docker)
1. Run Locally with Docker Compose
```bash
.
docker-compose up --build
```
App: http://localhost:8080

Adminer: http://localhost:9090

2. Run Tests Locally
Need to export the testdatabase
```
.
export TEST_DATABASE_URL="postgres://myuser:example@localhost:5433/testdb?sslmode=disable"
```
run the test
```
.
go test . -v
```
3. Try It Out (Signup, Login, Protected Access)
```
curl -i -X POST   -H "Content-Type: application/json"   -d '{"username":"alice","password":"secret"}'   http://localhost:8080/signup

curl -i -X POST   -H "Content-Type: application/json"   -d '{"username":"alice","password":"secret"}'   http://localhost:8080/login

curl -i -H "Authorization: Bearer <token>"   http://localhost:8080/protected
```
## ☸️ Kubernetes
Deploy the App
Need to add the dockerfile
```
.
eval $(minikube docker-env)
docker build -f Dockerfile

kubectl apply -f postgres-deployment.yaml
kubectl apply -f postgres-service.yaml

kubectl apply -f app-deployment.yaml
kubectl apply -f app-service.yaml
```
Deploy Test Database
```
.
kubectl apply -f postgres-test-deployment.yaml
kubectl apply -f postgres-test-service.yaml
```
Build and Run Test Job (in Minikube)
```
eval $(minikube docker-env)
docker build -f Dockerfile.test -t jwt-go-test .

kubectl apply -f test-job.yaml
kubectl logs -l job-name=run-go-tests
```
## API Endpoints
```bash
.
Method	Path	Description
POST	/signup	Create a new user
POST	/login	Get JWT token
GET	/protected	Requires valid token
```
## Why This Project?
This project was built to:

Understand JWT authentication.

Learn Kubernetes.

Learn how to build tests.
