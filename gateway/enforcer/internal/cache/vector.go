package cache

import (
	"context"
	"fmt"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

// QueryResult represents a search result with text, score, and associated answer.
type QueryResult struct {
	Text   string
	Score  float32
	Answer string
}

// VectorProvider is an interface for querying and storing vector embeddings.
type VectorProvider interface {
	QueryEmbedding(emb []float32) ([]QueryResult, error)
	UploadAnswerAndEmbedding(queryString string, queryEmb []float32, answer string) error
}

// MilvusVectorProvider is an implementation of VectorProvider using Milvus as the storage backend.
type MilvusVectorProvider struct {
	client client.Client
	ctx    context.Context
}

const milvusHost = "host.docker.internal"
const milvusPort = "19530"
const collectionName = "cache"
const diamentions = 384

// NewMilvusVectorProvider initializes and returns a new instance of MilvusVectorProvider.
func NewMilvusVectorProvider() (VectorProvider, error) {

	ctx := context.Background()
	// Connect to Milvus
	c, err := client.NewGrpcClient(ctx, fmt.Sprintf("%s:%s", milvusHost, milvusPort))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Milvus: %v", err)
	}

	has, err := c.HasCollection(ctx, collectionName)
	if err != nil {
		return nil, fmt.Errorf("failed to check collection existence: %v", err)
	}

	if has {
		errDrop := c.DropCollection(ctx, collectionName)
		if errDrop != nil {
			return nil, fmt.Errorf("failed to drop collection: %v", errDrop)
		}
	}

	schema := entity.NewSchema().WithDynamicFieldEnabled(false).
		WithField(entity.NewField().WithName("id").WithIsAutoID(true).WithDataType(entity.FieldTypeInt64).WithIsPrimaryKey(true)).
		WithField(entity.NewField().WithName("text").WithDataType(entity.FieldTypeVarChar).WithMaxLength(4096)).
		WithField(entity.NewField().WithName("embedding").WithDataType(entity.FieldTypeFloatVector).WithDim(diamentions)).
		WithField(entity.NewField().WithName("answer").WithDataType(entity.FieldTypeVarChar).WithMaxLength(4096))

	schema.CollectionName = collectionName
	schema.Description = "Cache for query embeddings and answers"

	// fields := []*entity.Field{
	// 	{
	// 		Name:        "id",
	// 		DataType:    entity.FieldTypeInt64,
	// 		Description: "Query text",
	// 		PrimaryKey:  true,
	// 		AutoID:      true,
	// 	},
	// 	{
	// 		Name:        "text",
	// 		DataType:    entity.FieldTypeVarChar,
	// 		Description: "Query text",
	// 	},
	// 	{
	// 		Name:        "embedding",
	// 		DataType:    entity.FieldTypeFloatVector,
	// 		Description: "Query embedding",
	// 		TypeParams: map[string]string{
	// 			"dim": "384",
	// 		},
	// 	},
	// 	{
	// 		Name:        "answer",
	// 		DataType:    entity.FieldTypeVarChar,
	// 		Description: "Answer to the query",
	// 	},
	// }

	err = c.CreateCollection(ctx, schema, 2)

	if err != nil {
		return nil, fmt.Errorf("failed to create collection: %v", err)
	}

	// create index
	idxHnsw, errIdx := entity.NewIndexHNSW(entity.COSINE, 8, 200)
	if errIdx != nil {
		return nil, fmt.Errorf("failed to new index: %v", errIdx)
	}
	errIndex := c.CreateIndex(ctx, collectionName, "embedding", idxHnsw, false, client.WithIndexName("myIdx"))
	if errIndex != nil {
		return nil, fmt.Errorf("failed to create index: %s", errIndex)
	}

	return &MilvusVectorProvider{
		client: c,
		ctx:    ctx,
	}, nil
}

// QueryEmbedding performs a vector search in Milvus and returns the closest matches.
func (mvp *MilvusVectorProvider) QueryEmbedding(queryEmb []float32) ([]QueryResult, error) {
	// Perform vector search query
	sp, err := entity.NewIndexHNSWSearchParam(74)
	if err != nil {
		return nil, fmt.Errorf("failed to new hnsw search params: %v", err)
	}

	// load collection
	Loaderr := mvp.client.LoadCollection(mvp.ctx, collectionName, false, client.WithReplicaNumber(1))
	if Loaderr != nil {
		return nil, fmt.Errorf("failed to load collection: %v", Loaderr)
	}

	searchResultList, err := mvp.client.Search(
		mvp.ctx,                    // ctx
		collectionName,             // CollectionName
		[]string{},                 // partitionNames
		"",                         // expr
		[]string{"text", "answer"}, // outputFields
		[]entity.Vector{entity.FloatVector(queryEmb)}, // vectors
		"embedding",   // vectorField
		entity.COSINE, // metricType
		5,             // topK
		sp,            // sp
	)
	if err != nil {
		return nil, fmt.Errorf("fail to search collection: %v", err)
	}

	queryResults := make([]QueryResult, 0) // No need to pre-allocate, append will handle it

	for _, searchResult := range searchResultList { // Iterate over SearchResults
		for i := 0; i < int(searchResult.ResultCount); i++ { // Iterate over entries in each SearchResult
			text, _ := searchResult.Fields[0].GetAsString(i)   // Correct way to get text
			answer, _ := searchResult.Fields[1].GetAsString(i) // Correct way to get answer
			score := searchResult.Scores[i]                    // Correct way to get score

			// Get embedding (if needed - it's usually not retrieved in a query)
			// embeddingColumn := searchResult.Fields[1]
			// embedding := make([]float32, embeddingColumn.Dim())
			// for j:=0; j < embeddingColumn.Dim(); j++ {
			//  embedding[j] = embeddingColumn.Get(i*embeddingColumn.Dim()+j).(float32) // Get the correct embedding
			// }

			queryResults = append(queryResults, QueryResult{
				Text: text,
				// Embedding: convertFloat32ToFloat64(embedding), // If you retrieve embedding
				Score:  score,
				Answer: answer,
			})
		}
	}

	return queryResults, nil
}

// UploadAnswerAndEmbedding stores the query string, its embedding, and the corresponding answer in Milvus.
func (mvp *MilvusVectorProvider) UploadAnswerAndEmbedding(queryString string, queryEmb []float32, answer string) error {
	_, err := mvp.client.Insert(mvp.ctx, collectionName, "",
		entity.NewColumnVarChar("text", []string{queryString}),
		entity.NewColumnFloatVector("embedding", diamentions, [][]float32{queryEmb}),
		entity.NewColumnVarChar("answer", []string{answer}),
	)
	if err != nil {
		return fmt.Errorf("failed to insert data: %v", err)
	}
	return nil
}
