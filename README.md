# 🔄 Data Transformation Platform

[![Go Version](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://golang.org)
[![React](https://img.shields.io/badge/React-18+-61DAFB.svg)](https://reactjs.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Docker](https://img.shields.io/badge/Docker-Ready-blue.svg)](docker-compose.yaml)

> A powerful, enterprise-grade data transformation platform for managing client data mappings and JSON transformations with real-time processing capabilities.

![Platform Screenshot](https://via.placeholder.com/800x400/1f2937/ffffff?text=Data+Transformation+Platform)

## 🌟 Overview

The Data Transformation Platform is a full-stack solution designed to handle complex data mapping and transformation scenarios. Built with Go and React, it provides a robust, scalable architecture for enterprise data integration needs.

### 🎯 Key Capabilities

- **🏢 Multi-Client Management**: Isolated environments for different clients with dedicated mapping rules
- **🔧 Advanced Mapping Engine**: Support for direct mapping, expressions, data validation, and custom transformations
- **📦 Bulk Operations**: Import/export mapping rules efficiently with JSON templates
- **⚡ Real-time Processing**: Instant data transformation with streaming support for large payloads
- **🔒 Enterprise Security**: JWT-based authentication with role-based access control
- **🎨 Modern UI**: Responsive, intuitive interface built with React and Tailwind CSS
- **📊 Expression Engine**: Dynamic transformations using custom expression language
- **🔄 Stream Processing**: Handle large datasets (5MB+) with memory-efficient streaming

## 📚 Documentation

| Resource | Description |
|----------|-------------|
| [📖 Bulk Mapping Guide](docs/BULK_MAPPING_GUIDE.md) | Complete guide for bulk import/export features |
| [🔧 API Documentation](#-api-reference) | Detailed API endpoint documentation |
| [🏗️ Architecture](#-project-structure) | System architecture and project structure |

## 🛠️ Technology Stack

| Component | Technology | Version |
|-----------|------------|---------|
| **Backend** | Go | 1.24+ |
| **Frontend** | React | 18+ |
| **Database** | PostgreSQL | 12+ |
| **Styling** | Tailwind CSS | 3+ |
| **Build Tool** | Vite | 4+ |
| **Authentication** | JWT | - |
| **Containerization** | Docker | 20+ |

## 🚀 Quick Start

### Prerequisites

- [Go 1.24+](https://golang.org/dl/)
- [Node.js 16+](https://nodejs.org/) & npm
- [PostgreSQL 12+](https://www.postgresql.org/)
- [Docker](https://www.docker.com/) (optional)

### 🔧 Installation & Setup

#### Option 1: Docker Deployment (Recommended)

```bash
# Clone the repository
git clone https://github.com/Aryan4GIT/Data-Transformation-Platform.git
cd Data-Transformation-Platform

# Start all services
docker-compose up --build

# Access the application
# Frontend: http://localhost:5173
# Backend: https://localhost:8080
```

#### Option 2: Manual Setup

**1. Backend Setup**

```bash
# Install Go dependencies
go mod download

# Set environment variables
export SERVER_PORT=8080
export DATABASE_URL="postgres://user:password@localhost:5432/data_mapping"
export JWT_SECRET="your_secure_secret_key"
export LOG_LEVEL="info"
export CERT_FILE_PATH="cert.pem"
export KEY_FILE_PATH="key.pem"

# Run the server
go run main.go
```

**2. Frontend Setup**

```bash
# Navigate to frontend directory
cd frontend

# Install dependencies
npm install

# Start development server
npm run dev

# Build for production
npm run build
```

## 🔐 Authentication

| Credential | Value |
|------------|-------|
| **Username** | `admin` |
| **Password** | `password` |

> ⚠️ **Security Notice**: Change default credentials before production deployment!

## 📡 API Reference

All protected endpoints require JWT authentication via `Authorization: Bearer <token>` header.

### 🔑 Authentication Endpoints
| Endpoint | Method | Auth | Description |
|----------|--------|------|-------------|
| `/login` | `POST` | ❌ | Authenticate user and receive JWT token |

### 👥 Client Management
| Endpoint | Method | Auth | Description |
|----------|--------|------|-------------|
| `/clients` | `GET` | ✅ | Retrieve all clients |
| `/clients` | `POST` | ✅ | Create a new client |
| `/clients/:id` | `DELETE` | ✅ | Delete a client and associated mappings |

### 🗂️ Mapping Rules
| Endpoint | Method | Auth | Description |
|----------|--------|------|-------------|
| `/clients/:client_id/mappings` | `GET` | ✅ | Get mapping rules for client |
| `/clients/:client_id/mappings` | `POST` | ✅ | Create mapping rules for client |
| `/mappings/:mapping_id` | `DELETE` | ✅ | Delete specific mapping rule |

### 🔄 Data Transformation
| Endpoint | Method | Auth | Description |
|----------|--------|------|-------------|
| `/clients/:client_id/transform` | `POST` | ✅ | Transform JSON using client mappings |

### 🔍 System Health
| Endpoint | Method | Auth | Description |
|----------|--------|------|-------------|
| `/health` | `GET` | ❌ | System health check |
| `/` | `GET` | ❌ | API information |

## 🏗️ Mapping Configuration

### Transformation Types

| Type | Description | Example |
|------|-------------|---------|
| **direct** | Simple field mapping | `sourcePath → destinationPath` |
| **expression** | Dynamic transformation | `income * 12` |
| **toString** | Convert to string | `123 → "123"` |
| **mapGender** | Gender mapping | `M → Male` |
| **formatDate** | Date formatting | `2023-01-01 → 01/01/2023` |

### Example Mapping Rule

```json
[
  {
    "sourcePath": "applicant.firstName",
    "destinationPath": "name.first",
    "required": true,
    "defaultValue": "",
    "transformType": "direct"
  },
  {
    "sourcePath": "applicant.income",
    "destinationPath": "financials.annualIncome", 
    "required": true,
    "transformType": "expression",
    "transformLogic": "income * 12"
  }
]
```

## ⚡ Advanced Features

### 🚀 Performance Optimizations
- **Streaming Support**: Handles large payloads (5MB+) with memory-efficient processing
- **Concurrent Processing**: Multi-threaded transformation for improved performance
- **Connection Pooling**: Optimized database connections for high throughput

### 🔒 Security Features
- **JWT Authentication**: Secure token-based authentication
- **HTTPS Support**: SSL/TLS encryption with certificate management
- **Input Validation**: Comprehensive request validation and sanitization
- **CORS Configuration**: Configurable cross-origin resource sharing

### 📊 Monitoring & Logging
- **Structured Logging**: JSON-formatted logs with multiple levels
- **Health Checks**: Built-in endpoint for monitoring system status
- **Error Tracking**: Comprehensive error handling and reporting

## 🏗️ Project Structure

```
📁 Data-Transformation-Platform/
├── 📄 main.go                     # Application entry point
├── 📄 go.mod                      # Go module definition
├── 📄 docker-compose.yaml         # Container orchestration
├── 📄 Dockerfile                  # Container build instructions
├── 🔐 cert.pem / key.pem          # SSL certificates
├── 📁 config/                     # Configuration management
│   └── 📄 config.go
├── 📁 database/                   # Database layer
│   ├── 📄 connection.go
│   └── 📁 migrations/
├── 📁 handlers/                   # HTTP request handlers
│   ├── 📄 client.go
│   ├── 📄 jwt.go
│   ├── 📄 mapping.go
│   └── 📄 transform.go
├── 📁 middleware/                 # HTTP middleware
│   ├── 📄 logging.go
│   └── 📄 security.go
├── 📁 models/                     # Data models
│   ├── 📄 models.go
│   ├── 📄 request.go
│   └── 📄 logs.go
├── 📁 utils/                      # Utility functions
│   ├── 📄 transform.go
│   └── 📄 validation.go
├── 📁 sample_data/                # Example data and templates
├── 📁 docs/                       # Documentation
└── 📁 frontend/                   # React application
    ├── 📄 package.json
    ├── 📄 vite.config.js
    ├── 📁 src/
    │   ├── 📄 App.jsx
    │   ├── 📁 components/
    │   ├── 📁 pages/
    │   ├── 📁 services/
    │   └── 📁 contexts/
    └── 📁 public/
```

## 🚨 Troubleshooting

### Common Issues

| Issue | Solution |
|-------|----------|
| **SSL Certificate Errors** | Accept self-signed certificates or replace with valid ones |
| **Database Connection Failed** | Verify PostgreSQL is running and credentials are correct |
| **JWT Token Expired** | Re-authenticate to get a new token |
| **CORS Errors** | Check frontend URL is allowed in CORS configuration |
| **Large Payload Timeouts** | Ensure streaming is enabled for payloads > 5MB |

### Environment Variables

```bash
# Required environment variables
SERVER_PORT=8080
DATABASE_URL=postgres://user:password@localhost:5432/data_mapping
JWT_SECRET=your_super_secure_secret_key_here
LOG_LEVEL=info
CERT_FILE_PATH=cert.pem
KEY_FILE_PATH=key.pem

# Optional environment variables
CORS_ALLOWED_ORIGINS=http://localhost:5173
MAX_PAYLOAD_SIZE=10485760  # 10MB
DB_MAX_CONNECTIONS=25
```

## Security Notes

- Replace the default credentials and JWT secret key before deploying to production.
- The current CORS configuration allows all origins in development mode. Configure appropriately for production.
#   D a t a - T r a n s f o r m a t i o n - P l a t f o r m 
 
 
## 🔒 Security Considerations

> ⚠️ **Important Security Notes**

- **Production Deployment**: Replace default credentials and JWT secret before production
- **CORS Configuration**: The current setup allows all origins in development. Configure appropriately for production
- **SSL Certificates**: Replace self-signed certificates with valid ones for production
- **Environment Variables**: Store sensitive data in environment variables, not in code
- **Database Security**: Use strong database passwords and limit database access

## 🤝 Contributing

We welcome contributions! Please follow these steps:

1. **Fork the repository**
2. **Create a feature branch** (`git checkout -b feature/amazing-feature`)
3. **Commit your changes** (`git commit -m 'Add amazing feature'`)
4. **Push to the branch** (`git push origin feature/amazing-feature`)
5. **Open a Pull Request**

### Development Guidelines

- Follow Go and React best practices
- Write comprehensive tests for new features
- Update documentation for API changes
- Ensure all CI checks pass

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built with ❤️ using Go and React
- Inspired by enterprise data integration needs
- Thanks to the open-source community

## 📞 Support

- **Issues**: [GitHub Issues](https://github.com/Aryan4GIT/Data-Transformation-Platform/issues)
- **Discussions**: [GitHub Discussions](https://github.com/Aryan4GIT/Data-Transformation-Platform/discussions)
- **Email**: [aryansingh73321@gmail.com](mailto:aryansingh73321@gmail.com)

---

<div align="center">

**[⬆ Back to Top](#-data-transformation-platform)**

Made with ❤️ by [Aryan Singh](https://github.com/Aryan4GIT)

⭐ **Star this repo if you find it helpful!** ⭐

</div>
