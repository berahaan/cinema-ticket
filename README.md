# 📽️ Cinema Schedule and Ticketing Website  

## 📋 Project Description  
This project is a fully functional cinema scheduling and ticketing platform, allowing users to:  
- Browse movies with schedules and average ratings.  
- Purchase tickets securely via Chapa.  
- Bookmark movies, download tickets, and receive success emails.  
- Manage their profiles.  

Admins can manage movies, schedules, and ticket sales directly.  

---

## ✨ Features  
- **Payment Integration**: Secure payments with Chapa.  
- **Dynamic Movie Listings**: View movies with average ratings and detailed schedules.  
- **Admin Features**: Manage movies, schedules, and ticket sales.  
- **User Features**: Bookmark,rate movies, and receive success emails with attachment.  
- **Responsive Design**: Optimized for both mobile and desktop.  

---

## 🚀 Technologies Used  
**Frontend**: Vue, Nuxt 3, Tailwind CSS  
**Backend**: Go (Golang), Hasura GraphQL Engine  
**Database**: PostgreSQL  
**Payments**: Chapa  
**Other Tools**: Apollo GraphQL, Pinia (state management), Postman  

---

## 🚀 Backend Powered by Golang  
The backend of this project is built using Golang, offering:  
- **High Performance**: Handles concurrent requests efficiently for real-time ticketing and scheduling.  
- **Scalability**: Supports increasing traffic with ease.  
- **Simplicity**: Clean and reliable codebase.  

### Responsibilities of the Golang Backend:  
- Payment processing and webhook handling for Chapa.  
- File uploads, authentication, and JWT generation using Hasura GraphQL Actions.  
- Ticket generation and email notifications.  

---

## 🐳 Docker for Containerization  
This project uses Docker to ensure a smooth setup and consistent environment.  

### Why Docker?  
- Simplifies setup by containerizing the app and its dependencies.  
- Provides consistency across local, staging, and production environments.  
- Enables easy scaling and deployment (e.g., Vercel or cloud providers).  

---

## 🛠️ Setup Instructions  

### Frontend Setup:  
1. Clone the repository:  
   ```bash  
   git clone https://github.com/berahaan/cinema-ticket.git  

2.Navigate to the project folder
```bash
cd cinema-frontend
```
3.Install dependencies:
for frontend 
```bash
npm install
npm run dev
```
4.Set up environment variables
 create .env if not exists already and add
```bash
 GRAPHQL_ENDPOINT ="Your_Graphql_Endpoint"
```
---
5.Access the project
Open http://localhost:3000 in your browser

🛠️ Setup Instructions for Backend
1.Navigate to project folder
```bash
cd cinema-backend
```
-for golang run
```bash 
go mod tidy 
```
so that install dependency
2.Set up environment variables
Then create .env and add this 
```bash
JWT_SECRET=""
JWT_REFRESH_SECRET=""
GRAPHQL_ENDPOINT =""
CHAPA_INITIALIZE_PAYMENT=""
CHAPA_SECRET_KEY =""
HASURA_GRAPHQL_ADMIN_SECRET= ""
CALL_BACK_URL=""
RETURN_URL =""
WEB_HOOK_SECRET_KEY =""
CHAPA_VERIFY_ENDPOINT=""
FEATURED_IMAGE_PATH =""
OTHER_IMAGE_PATH =""
PROFILE_IMAGE_PATH =""
FEATURED_FILE_PATH =""
OTHERIMG_FILE_PATH =""
PROFILE_FILE_PATH=""
EMAIL =""
GOOGLE_PASSWORD="your App Passwords "
```
for google_passwords you can generate it for your Emails under your google Accounts then search for App passwords then Generate it by adding Gmail as you app then copy paste it 
3.Running the containers 
```bash
docker-compose up -d
```
4.after running containers you can access Hasura at
```bash
http://localhost:8080
```
- wait for a few minutes if you machines is a bit older after setted by also you need wsl for windows to run docker locally 
- before running the above commands make sure you start docker engine
5.in current directory i.e cinema0backend run
  ```bash
  go run main.go
  ```
5.You can stop Containers by using  
```bash   
docker-compose down
```
📄 Key Endpoints
Payment Webhook: Confirm payments and generate tickets.
Image Uploads: Manage movie and profile images.
📬 Contact
For questions or feedback, feel free to reach out at my telegram  https://t.me/bekind2yourself
