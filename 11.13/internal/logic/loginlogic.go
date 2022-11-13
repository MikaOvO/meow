package logic

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"11.13/internal/svc"
	"11.13/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("meow").Collection("user")

	var results []types.User

	filter := bson.D{{"name", req.Name}}

	res, _ := collection.Find(context.TODO(), filter)

	_ = res.All(context.TODO(), &results)

	if len(results) == 0 {
		return &types.LoginResponse{
			Result: "user has not registered!",
		}, nil
	}

	user := results[0]

	if req.Password != user.Password {
		return &types.LoginResponse{
			Result: "passwords do not matched!",
		}, nil
	}
	return &types.LoginResponse{
		Result: "login success!",
	}, nil
}
