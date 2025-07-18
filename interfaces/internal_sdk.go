package interfaces

import (
	"context"
	"github.com/apito-io/types"
)

// InternalSDKOperation defines the interface that this SDK implements
// This matches the interface from the main Apito Engine
type InternalSDKOperation interface {
	// GenerateTenantToken generates a new tenant token for the specified tenant ID
	GenerateTenantToken(ctx context.Context, token string, tenantID string) (string, error)

	// GetProjectDetails retrieves project details for the given project ID
	//GetProjectDetails(ctx context.Context, projectID string) (*protobuff.Project, error)

	// GetSingleResource retrieves a single resource by model and ID, with optional single page data
	GetSingleResource(ctx context.Context, model, _id string, singlePageData bool) (*types.DefaultDocumentStructure, error)

	// SearchResources searches for resources in the specified model using the provided filter
	SearchResources(ctx context.Context, model string, filter map[string]interface{}, aggregate bool) (*types.SearchResult, error)

	// GetRelationDocuments retrieves related documents for the given ID and connection parameters
	GetRelationDocuments(ctx context.Context, _id string, connection map[string]interface{}) (*types.SearchResult, error)

	// CreateNewResource creates a new resource in the specified model with the given data and connections
	CreateNewResource(ctx context.Context, request *types.CreateAndUpdateRequest) (*types.DefaultDocumentStructure, error)

	// UpdateResource updates an existing resource by model and ID, with optional single page data, data updates, and connection changes
	UpdateResource(ctx context.Context, request *types.CreateAndUpdateRequest) (*types.DefaultDocumentStructure, error)

	// DeleteResource deletes a resource by model and ID
	DeleteResource(ctx context.Context, model, _id string) error

	// Debug is used to debug the plugin, you can pass data here to debug the plugin
	Debug(ctx context.Context, stage string, data ...interface{}) (interface{}, error)
}