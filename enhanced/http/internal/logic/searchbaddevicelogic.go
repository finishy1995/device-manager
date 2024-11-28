package logic

import (
	"context"
	"encoding/json"

	"finishy1995/device-manager/enhanced/http/internal/svc"
	"finishy1995/device-manager/enhanced/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SearchBadDeviceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchBadDeviceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchBadDeviceLogic {
	return &SearchBadDeviceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchBadDeviceLogic) SearchBadDevice(req *types.SearchBadDeviceReq) (resp *types.SearchBadDeviceResp, err error) {
	data := map[string]interface{}{}
	if req.Type == 0 || req.Type == 1 {
		// search by Gradual Drift
		query := bson.A{
			bson.M{
				"$vectorSearch": bson.M{
					"index":         "cosine_index",
					"path":          "vector_angle",
					"queryVector":   GradualDriftEuclideanDistance,
					"numCandidates": 50,
					"limit":         req.Limit,
				},
			},
		}

		// Access the collection and execute the aggregation query
		cursor, errResp := l.svcCtx.MongoClient.Database("test").Collection("device_vector_data").Aggregate(l.ctx, query)
		if errResp != nil {
			l.Error("Failed to aggregate data", err)
		} else {
			var results []primitive.M
			if errResp := cursor.All(l.ctx, &results); err != nil {
				l.Error("Failed to decode data", errResp)
				return
			}
			for _, result := range results {
				delete(result, "vector_euclidean_distance")
				delete(result, "vector_angle")
			}
			data["Gradual Drift"] = results
		}
	}
	if req.Type == 0 || req.Type == 2 {
		// search by Fixed
		query := bson.A{
			bson.M{
				"$vectorSearch": bson.M{
					"index":         "euclidean_index",
					"path":          "vector_euclidean_distance",
					"queryVector":   FixedEuclideanDistance,
					"numCandidates": 50,
					"limit":         req.Limit,
				},
			},
		}

		// Access the collection and execute the aggregation query
		cursor, errResp := l.svcCtx.MongoClient.Database("test").Collection("device_vector_data").Aggregate(l.ctx, query)
		if errResp != nil {
			l.Error("Failed to aggregate data", err)
		} else {
			var results []primitive.M
			if errResp := cursor.All(l.ctx, &results); err != nil {
				l.Error("Failed to decode data", errResp)
				return
			}
			for _, result := range results {
				delete(result, "vector_euclidean_distance")
				delete(result, "vector_angle")
			}
			data["Fixed"] = results
		}
	}
	// if req.Type == 0 || req.Type == 3 {
	// 	// search by Non-physical Rotations
	// 	query := bson.A{
	// 		bson.M{
	// 			"$vectorSearch": bson.M{
	// 				"index":         "cosine_index",
	// 				"path":          "vector_angle",
	// 				"queryVector":   RandomAngleBetweenVectors,
	// 				"numCandidates": 50,
	// 				"limit":         req.Limit,
	// 			},
	// 		},
	// 	}

	// 	// Access the collection and execute the aggregation query
	// 	cursor, errResp := l.svcCtx.MongoClient.Database("test").Collection("device_vector_data").Aggregate(l.ctx, query)
	// 	if errResp != nil {
	// 		l.Error("Failed to aggregate data", err)
	// 	} else {
	// 		var results []primitive.M
	// 		if errResp := cursor.All(l.ctx, &results); err != nil {
	// 			l.Error("Failed to decode data", errResp)
	// 			return
	// 		}
	// 		data["Non-physical Rotations"] = results
	// 	}
	// }

	jsonData, err := json.Marshal(data)
	if err != nil {
		l.Error("Failed to marshal data", err)
		return
	}
	return &types.SearchBadDeviceResp{
		Code: 200,
		Data: string(jsonData),
	}, nil
}
