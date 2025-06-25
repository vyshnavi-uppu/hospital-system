🏥 Hospital Management System (Golang + Gin + PostgreSQL)
A simple yet powerful backend API system designed for hospitals to manage patients via receptionist and doctor portals. Built using Go (Gin), PostgreSQL, and JWT authentication.

🚀 Features
✅ JWT Authentication for login
🧑‍💼 Receptionist Portal:

Register new patients

Perform full CRUD (Create, Read, Update, Delete) on patients

🩺 Doctor Portal:

View patient details

Update patient diagnosis/info

🔐 Role-Based Access Control (RBAC)
📦 PostgreSQL as the backend database
🧪 Unit testing support
🌐 Clean and modular directory structure
📄 API documented using Postman / Swagger

📂 Project Structure
pgsql
Copy
Edit
hospital_system/
├── config/         → Database connection (config/db.go)
├── controllers/    → Business logic (auth.go, patient.go)
├── middleware/     → JWT auth, role-based access
├── models/         → Structs and DB schema models
├── routes/         → Route grouping for APIs
├── utils/          → Helper functions (optional)
├── main.go         → Entry point
├── go.mod / go.sum → Go module files
🛠️ Technologies Used
Golang 1.20+

Gin web framework

PostgreSQL

JWT (github.com/golang-jwt/jwt/v5)

pgAdmin (for DB management)

Postman (for testing APIs)

🔐 Authentication
Single login API: POST /login

Returns a JWT token

User roles: receptionist or doctor

Routes are protected based on user role

🧑‍💻 API Endpoints
Base URL: http://localhost:8080

Method	Endpoint	Description	Access
POST	/login	Login and receive JWT	Public
POST	/patients	Register new patient	Receptionist
GET	/patients	Get all patients	All roles
GET	/patients/:id	Get a specific patient	All roles
PUT	/patients/:id	Update patient info	Doctor
DELETE	/patients/:id	Delete a patient	Receptionist

📦 Sample Login Payload
json
Copy
Edit
POST /login
{
  "username": "reception1",
  "password": "pass123"
}
🧪 Testing with Postman
Import postman_collection.json from /docs/postman/

Use JWT token in Authorization header → Bearer Token

Test each role (doctor, receptionist) with different users

🧾 How to Run
Clone the repo:

bash
Copy
Edit
git clone https://github.com/yourusername/hospital_system.git
cd hospital_system
Set up PostgreSQL:

Create database: hospital_db

Create tables (use SQL migration or schema file if provided)

Run the app:

bash
Copy
Edit
go mod tidy
go run main.go
✅ You should see:

Connected to PostgreSQL

Listening on :8080




