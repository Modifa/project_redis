package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	models "github.com/Modifa/project_redis.git/models"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
)

func SaveTest(User models.User) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDISSERVER_HOST") + ":" + os.Getenv("REDISSERVER_PORT"),
		Password:   os.Getenv("REDISSERVER_PASSWORD"), // no password set
		DB:         0,
		MaxConnAge: 0,
	})
	key := User.PhoneNumber
	b, err := json.Marshal(User)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	/**/
	err = rdb.Set(ctx, "USER:"+key, b, 0).Err()
	return err
}

func SaveTransactions(TransactionDetails models.DonationsResponse) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDISSERVER_HOST") + ":" + os.Getenv("REDISSERVER_PORT"),
		Password:   os.Getenv("REDISSERVER_PASSWORD"), // no password set
		DB:         0,                                 /*LookUP*/
		MaxConnAge: 0,
	})

	b, err := json.Marshal(TransactionDetails)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// err = rdb.Set(ctx, "PORTFOLIO:TRANSACTIONS:"+TransactionId, b, 0).Err()
	err = rdb.LPush(ctx, "PORTFOLIO:TRANSACTIONS", b).Err()

	if err != nil {
		fmt.Println(err)
	}

	return err
}
