package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql/graph/generated"
	"graphql/graph/model"
	"graphql/mgo"
	"log"
	"math/rand"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context, count *int) ([]*model.Todo, error) {
	var vv []*model.Todo
	var a = 1
	var c = *count
	//var idx *int
	for _, v := range r.todos {
		if a > c {
			break
		}
		vv = append(vv, v)
		a++
	}
	return vv, nil
}

func (r *queryResolver) GetDashboard(ctx context.Context, condition *model.DashboardQueryCondition) ([]*model.Dashboard, error) {
	env := *condition.Env
	cdType := *condition.CdType
	keyword := *condition.Keyword
	cluster := *condition.Cluster
	namespace := *condition.Namespace
	startTime := *condition.StartDatetime
	endTime := *condition.EndDatetime
	status := *condition.Status
	page := *condition.Page
	pagesize := *condition.Pagesize

	if env == "" {
		env = "prod"
	}
	name := env + "_monitoring"
	collectionName := strings.ReplaceAll(name, "-", "_")
	collection := mgo.MongoDB.Collection(collectionName)

	var filter []bson.M
	filter = append(filter, bson.M{"isdelete": bson.M{"$ne": "y"}})

	if cdType != "" {
		filter = append(filter, bson.M{"cdtype": bson.M{"$eq": cdType}})
	}

	if keyword != "" {
		filterKeyword := bson.M{
			"$or": []interface{}{
				bson.M{"application": bson.M{"$regex": keyword, "$options": "$i"}},
				bson.M{"helmpackage": bson.M{"$regex": keyword, "$options": "$i"}},
			},
		}
		filter = append(filter, filterKeyword)
	}

	if cluster != "" {
		filter = append(filter, bson.M{"cluster": bson.M{"$eq": cdType}})
	}

	if namespace != "" {
		filter = append(filter, bson.M{"namespace": bson.M{"$eq": namespace}})
	}

	// UTC
	dateFormat := "2006-01-02 15:04:05"
	if startTime != "" {
		startTimeStampT, _ := time.Parse(dateFormat, startTime)
		startTimeStamp := startTimeStampT.Unix()
		filter = append(filter, bson.M{"datetime": bson.M{"$gte": startTimeStamp}})
	}

	if endTime != "" {
		endTimeStampT, _ := time.Parse(dateFormat, endTime)
		endTimeStamp := endTimeStampT.Unix()
		filter = append(filter, bson.M{"datetime": bson.M{"$lte": endTimeStamp}})
	}

	if status != "" {
		filter = append(filter, bson.M{"status": bson.M{"$eq": namespace}})
	}

	options := options.Find()
	if page > 0 && pagesize > 0 {
		offset := int64(pagesize * (page - 1))
		options.SetSkip(offset)
		options.SetLimit(int64(pagesize))
	}
	options.SetSort(bson.D{{"datetime", -1}})
	cur, err := collection.Find(context.Background(), bson.M{"$and": filter}, options)

	if err != nil {
		fmt.Printf("get deploy history list error")
	}

	var dashboardList []*model.Dashboard
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		dashboard := model.Dashboard{}
		err := cur.Decode(&dashboard)
		if err != nil {
			log.Printf("decode deploy history item error")
		}

		dashboardList = append(dashboardList, &dashboard)
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}

	return dashboardList, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
