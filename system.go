package types

type SystemUser struct {
	XKey string `json:"_key,omitempty" firestore:"_key,omitempty" bson:"_key,omitempty"`
	ID   string `bun:"type:uuid,pk" json:"id,omitempty" firestore:"id,omitempty" bson:"_id,omitempty"`

	FirstName string `json:"first_name,omitempty" firestore:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty" firestore:"last_name,omitempty" bson:"last_name,omitempty"`
	Role      string `json:"role,omitempty" firestore:"role,omitempty" bson:"role,omitempty"`
	Username  string `json:"username,omitempty" firestore:"username,omitempty" bson:"username,omitempty"`
	Email     string `json:"email,omitempty" firestore:"email,omitempty" bson:"email,omitempty"`
	
	IsActive bool `json:"is_active,omitempty" firestore:"is_active,omitempty" bson:"is_active,omitempty"`
	IsProjectUser bool `json:"is_project_user,omitempty" firestore:"is_project_user,omitempty" bson:"is_project_user,omitempty"`
}

type MetaField struct {
	SourceID string `json:"source_id,omitempty" firestore:"source_id,omitempty" bson:"source_id,omitempty"`

	CreatedAt      string      `json:"created_at,omitempty" firestore:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt      string      `json:"updated_at,omitempty" firestore:"updated_at,omitempty" bson:"updated_at,omitempty"`
	CreatedBy      *SystemUser `json:"created_by,omitempty" firestore:"title,omitempty" bson:"created_by,omitempty"`
	LastModifiedBy *SystemUser `json:"last_modified_by,omitempty" firestore:"created_by,omitempty" bson:"last_modified_by,omitempty"`

	Status string `json:"status,omitempty" firestore:"status,omitempty" bson:"status,omitempty"`
	//TenantId       string `json:"tenant_id,omitempty" firestore:"tenant_id,omitempty"` move to root removed from meta
	RootRevisionID string `json:"root_revision_id,omitempty" firestore:"root_revision_id,omitempty" bson:"root_revision_id,omitempty"`
	Revision       bool   `json:"revision,omitempty" firestore:"revision,omitempty" bson:"revision,omitempty"`
	RevisionAt     string `json:"revision_at,omitempty" firestore:"revision_at,omitempty" bson:"revision_at,omitempty"`

	// used in filterAbsentStudent where multiple record is processed but we need to return only attendance id
	ResourceID string `json:"resource_id,omitempty" firestore:"resource_id,omitempty" bson:"resource_id,omitempty"`
}