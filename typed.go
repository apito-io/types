package types

// TypedSearchResult represents a search result with typed data
type TypedSearchResult[T any] struct {
	Results []*TypedDocumentStructure[T] `json:"results"`
	Count   int                          `json:"count"`
}

// TypedDocumentStructure represents a document with typed data
type TypedDocumentStructure[T any] struct {
	Key           string     `json:"_key,omitempty" firestore:"_key,omitempty" bson:"_key,omitempty"`
	Data          T          `json:"data,omitempty" firestore:"data,omitempty" bson:"data,omitempty"`
	Meta          *MetaField `json:"meta,omitempty" firestore:"meta,omitempty" bson:"meta,omitempty"`
	ID            string     `json:"id,omitempty" firestore:"id,omitempty" bson:"id,omitempty"`
	ExpireAt      int64      `json:"expire_at,omitempty" firestore:"expire_at,omitempty" bson:"expire_at,omitempty"`
	RelationDocID string     `json:"relation_doc_id,omitempty" firestore:"relation_doc_id,omitempty" bson:"relation_doc_id,omitempty"`
	Type          string     `json:"type,omitempty" firestore:"type,omitempty" bson:"type,omitempty"`
	TenantID      string     `json:"tenant_id,omitempty" firestore:"tenant_id,omitempty" bson:"tenant_id,omitempty"`
	TenantModel   string     `json:"tenant_model,omitempty" firestore:"tenant_model,omitempty" bson:"tenant_model,omitempty"`
}

type SearchResult struct {
	Results []*DefaultDocumentStructure `json:"results"`
	Count   int                         `json:"count"`
}

// GraphQLResponse represents a generic GraphQL response
type GraphQLResponse struct {
	Data   interface{}    `json:"data,omitempty"`
	Errors []GraphQLError `json:"errors,omitempty"`
}

// GraphQLError represents a GraphQL error
type GraphQLError struct {
	Message    string                 `json:"message"`
	Locations  []GraphQLErrorLocation `json:"locations,omitempty"`
	Path       []interface{}          `json:"path,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

// GraphQLErrorLocation represents the location of a GraphQL error
type GraphQLErrorLocation struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

type CreateAndUpdateRequest struct {
	ID             string                 `json:"id"`
	Model          string                 `json:"model"`
	Payload        interface{}            `json:"payload"`
	Connect        map[string]interface{} `json:"connect"`
	Disconnect     map[string]interface{} `json:"disconnect"`
	SinglePageData bool                   `json:"single_page_data"`
	ForceUpdate    bool                   `json:"force_update"`
}
