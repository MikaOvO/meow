package logic

import (
	"11.13/internal/svc"
	"11.13/internal/types"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
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

	if len(results) > 0 {
		return &types.RegisterResponse{
			Result: "user has registered!",
		}, nil
	}

	user := types.User{
		req.Name,
		req.Password,
	}

	insertResult, _ := collection.InsertOne(context.TODO(), user)

	fmt.Println("insert: ", insertResult)

	return &types.RegisterResponse{
		Result: "register success!",
	}, nil
}
