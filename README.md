# IoT Cloud Orchestrator (EO LM Simulation)

A TLS-enabled Go microservice that simulates Containerized Network Function (CNF) lifecycle management — inspired by Ericsson's EO Lifecycle Manager.

## 🔧 Features

- Secure Go REST API (`/device/update`, `/device/status`)
- TLS + Token-based Authentication
- Dockerized & Kubernetes-ready
- CNF-like devices with auto-recovery logic
- Deployed and tested in Minikube

## 🧪 API Endpoints

- `POST /device/update`  
  Authenticated. Accepts:
  ```json
  {
    "id": "sensor01",
    "status": "online"
  }