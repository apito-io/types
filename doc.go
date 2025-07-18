package types

// DefaultDocumentStructure is the default document structure for the database
type DefaultDocumentStructure struct {
	Key      string                 `json:"_key,omitempty" firestore:"_key,omitempty" bson:"_key,omitempty"`
	ID       string                 `json:"id,omitempty" firestore:"id,omitempty" bson:"id,omitempty"`
	XID      ID                     `json:"_id,omitempty" firestore:"_id,omitempty" bson:"_id,omitempty"`
	Type     string                 `json:"type,omitempty" firestore:"type,omitempty" bson:"type,omitempty"`
	Data     map[string]interface{} `json:"data,omitempty" firestore:"data,omitempty" bson:"data,omitempty"`
	Meta     *MetaField   `json:"meta,omitempty" firestore:"meta,omitempty" bson:"meta,omitempty"`
	ExpireAt string                 `json:"expire_at,omitempty" firestore:"expire_at,omitempty" bson:"expire_at,omitempty"`

	TenantID    ID     `json:"tenant_id,omitempty" firestore:"tenant_id,omitempty" bson:"tenant_id,omitempty"`
	TenantModel string `json:"tenant_model,omitempty" firestore:"tenant_model,omitempty" bson:"tenant_model,omitempty"`

	// this is used in table data fetching for relation doc id, and used in disconnection of the relation
	// this is the id of the relation doc
	RelationDocID string `json:"relation_doc_id,omitempty" firestore:"relation_doc_id,omitempty" bson:"relation_doc_id,omitempty"`

	// if revision is enabled, this is the id of the last revision doc
	LastRevisionDocID string `json:"last_revision_doc_id,omitempty" firestore:"last_revision_doc_id,omitempty" bson:"last_revision_doc_id,omitempty"`
}