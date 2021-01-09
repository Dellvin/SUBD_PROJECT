package main

import (
	"SUBD_PROJECT/conf"
	"SUBD_PROJECT/internal/app/forum/Repository/forum"
	fde "SUBD_PROJECT/internal/app/forum/delivery"
	ude "SUBD_PROJECT/internal/app/user/delivery"
	ure "SUBD_PROJECT/internal/app/user/repository"
	"SUBD_PROJECT/internal/pkg/json"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/jackc/pgx"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

func main() {
	r := router.New()
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable port=%s",
		conf.Postgres.User,
		conf.Postgres.Password,
		conf.Postgres.DBName,
		conf.Postgres.Port)
	pgxConn, err := pgx.ParseConnectionString(connStr)
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}
	pgxConn.PreferSimpleProtocol = true
	config := pgx.ConnPoolConfig{
		ConnConfig:     pgxConn,
		MaxConnections: 100,
		AfterConnect:   nil,
		AcquireTimeout: 0,
	}

	connPool, err := pgx.NewConnPool(config)
	if err != nil {
		log.Error().Msgf(err.Error())
	}

	userRepo := ure.New(connPool)
	forumRepo := Postgres.New(connPool, userRepo)

	ude.New(r, userRepo, forumRepo)
	fde.New(r, forumRepo, userRepo)
	port := ":5000"
	fmt.Println("Starts on:", port)
	log.Error().Msgf(fasthttp.ListenAndServe(port, json.AppJson(r.Handler)).Error())
}
