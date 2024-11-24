package routes

import (
	"cinema-backend/auth"
	"cinema-backend/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// here we can handles routes logic here
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/login", handler.LoginHandler).Methods("POST")
	router.HandleFunc("/insertAdmin", handler.InsertAdminHandler).Methods("POST")
	router.HandleFunc("/signup", handler.Singuphandler).Methods("POST")
	router.HandleFunc("/PaymentInitiate", handler.HandlePaymentInitiate).Methods("POST")
	router.HandleFunc("/chapaWebhook", handler.ChapaWebhookHandler).Methods("POST")
	router.HandleFunc("/fileUpload", handler.UploadFilesHandles).Methods("POST")
	router.HandleFunc("/downloadTicket", handler.HandleDownloadTicket).Methods("POST")
	router.HandleFunc("/profileUpload", handler.HandleProfileUpload).Methods("POST")
	router.HandleFunc("/deleteFeaturedImage", handler.HandleImageDelete).Methods("POST")
	router.HandleFunc("/deleteOtherImage", handler.HandleOtherImagesDelete).Methods("POST")
	router.HandleFunc("/TokenAuth", auth.HandleToken).Methods("POST")
	featuredImages := http.FileServer(http.Dir("./movies/upload/featuredImage"))
	profileImages := http.FileServer(http.Dir("./movies/upload/profilePhoto"))
	otherImages := http.FileServer(http.Dir("./movies/upload/otherImages"))
	router.PathPrefix("/uploads/profilePhoto/").Handler(http.StripPrefix("/uploads/profilePhoto/", profileImages))
	router.PathPrefix("/uploads/featuredImage/").Handler(http.StripPrefix("/uploads/featuredImage/", featuredImages))
	router.PathPrefix("/uploads/otherImages/").Handler(http.StripPrefix("/uploads/otherImages/", otherImages))
	log.Println("Server running on port 4000")
	err := http.ListenAndServe(":4000", auth.EnableCORS(router))
	if err != nil {
		log.Fatal(err)
	}
	return router
}
