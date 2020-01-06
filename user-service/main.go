package main

import (
	"log"

	pb "github.com/ab22/shippy/user-service/proto/user"
	micro "github.com/micro/go-micro"
)

func main() {
	log.Println("Starting server...")
	log.Println("Creating database connection...")
	db, err := CreatePostgresConnection()

	if err != nil {
		log.Fatalln("Could not establish connection to postgres database:", err)
	}

	log.Println("Auto migrating models...")
	defer db.Close()
	db.AutoMigrate(&pb.User{})

	log.Println("Creating microservice...")
	var (
		srv = micro.NewService(
			micro.Name("shippy.user.service"),
			micro.Version("latest"),
		)
		repo         = &UserRepository{db}
		tokenService = &TokenService{repo}
	)

	srv.Init()
	pb.RegisterUserServiceHandler(srv.Server(), &handler{repo, tokenService})

	log.Println("Running...")
	if err := srv.Run(); err != nil {
		log.Println("Error running microservice:", err)
	}
}
