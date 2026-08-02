#  Cipher Project

## Overview
Cipher Project is a secure web‑based application built in Golang that allows users to **encode and decode text using custom cipher mappings**.  
It features authentication, encryption logic, and PostgreSQL integration, designed with a futuristic cyber‑security theme.

---

## Key Achievements
- Implemented **JWT authentication** with secure cookie handling.  
- Built **password hashing** using bcrypt for strong security.  
- Designed a **custom cipher engine** to encode/decode text with digit‑to‑symbol mappings.  
- Integrated **PostgreSQL databases** for user management and cipher mappings.  
- Developed a modular backend with Go, ensuring scalability and maintainability.  

---

## Tech Highlights
- **Backend:** Golang (Gin framework)  
- **Database:** PostgreSQL (User DB + Cipher DB)  
- **Authentication:** JWT + bcrypt  
- **Frontend:** HTML5/CSS3 templates (Signup, Login, Encode, Decode, Home)  
- **Deployment:** Dockerized for portability  

---

## Skills Demonstrated
- Secure backend development (auth middleware, JWT, cookies)  
- Cryptography basics (custom cipher mapping, encoding/decoding logic)  
- Database design and integration with PostgreSQL  
- Error handling and fallback mechanisms in Go  
- Full‑stack development with modular architecture  

---

## 📂 Project Structure
<details>
<summary>Click to expand</summary>
CIPHER PROJECT/
│── assets/
│   ├── decode.html
│   ├── encode.html
│   ├── home.html
│   ├── login.html
│   ├── signup.html
│── database/
│   ├── cipher_code.sql
│   ├── cipher_mapping_1.json
│   ├── cipher_num.sql
│   ├── user.sql
│── auth.go
│── cipher.go
│── database.go
│── main.go
│── sqlconnect.go
│── go.mod
│── go.sum


</details>

---

## Installation & Setup
### Clone the Repository
`git clone https://github.com/gurumdlrdev-hub/cipher_project.git`
`cd cipher_project`
### Build Docker Image
`docker build -t cipher-project .`
### Run Container
`docker run -p 7777:7777 cipher-project`
### Access via Browser
`http://localhost:7777`

---

## Architecture Diagram

┌──────────────────────────────────────────┐
│           Client Side Browser            │
└────────────────────┬─────────────────────┘
                     │
              (HTTPS Request)
                     │
                     ▼
┌──────────────────────────────────────────┐
│            Docker Container              │
│                                          │
│  ┌────────────────────────────────────┐  │
│  │   Go App (Gin Router + Cipher API) │  │
│  │                                    │  │
│  │  ├── Auth Engine (JWT + bcrypt)    │  │
│  │  ├── Encode Module (cipher.go)     │  │
│  │  ├── Decode Module (cipher.go)     │  │
│  │  └── Database Connector            │  │
│  └────────────────────────────────────┘  │
└────────────────────┬─────────────────────┘
                     │
        (Secure Database Connection)
                     │
                     ▼
┌──────────────────────────────────────────┐
│        PostgreSQL Cloud Databases        │
│   ├── app_user (User Accounts)           │
│   └── cipher_project (Cipher Mappings)   │
└──────────────────────────────────────────┘

---

## 🔑 Example: Encoding

**Input:** Guru Golang Developer
**Output (Cipher):** *!?# #!\^ !!!! &@&? $!%* %^@& !!!! $&?$ ##^\ \*^^ %!& **&&
---
## 🔑 Example: Decoding

**Input (Cipher):** ```@*#* !!!! %\? !!!! \?*^ $@@# !!!! #^%& &!%# !!!! *?@$ ?%^! \*@?```
**Output (Decoded):** I am Loki from Ascard

---

## 📸 Screenshots

### Sign Up Page
![Signup Screenshot](https://res.cloudinary.com/ygllvtyg/image/upload/v1785701928/Screenshot_2026-08-03_001056_cmscxg.png)

### Log in Page
![Login Screenshot](https://res.cloudinary.com/ygllvtyg/image/upload/v1785701928/Screenshot_2026-08-03_001132_hi7dtf.png)

### Home Page
![Home Page Screenshot](https://res.cloudinary.com/ygllvtyg/image/upload/v1785701928/Screenshot_2026-08-03_001211_iexpga.png)

### Encode Page
![Encode Screenshot](https://res.cloudinary.com/ygllvtyg/image/upload/v1785702611/Screenshot_2026-08-03_015633_rbotjb.png)

### Decode Page
![Decode Screenshot](https://res.cloudinary.com/ygllvtyg/image/upload/v1785702272/Screenshot_2026-08-03_015416_fybynj.png)

---

## Usage
* Sign Up / Login to create a secure account.
*  Navigate to the Home Dashboard.
*  Use the Encode Page to convert plain text into cipher symbols.
*  Use the Decode Page to convert cipher symbols back into readable text.
*  Logout securely with JWT session termination.
---

## Author
Guru – Software Engineering Intern
Focused on backend development, cryptography basics, and scalable systems.
