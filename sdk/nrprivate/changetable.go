package nrprivate

type ChangeTable  struct {
	CollectionName string `json:"collectionName,omitempty"`
	Collection    interface{} `json:"collection,omitempty"`
	Method         string `json:"method,omitempty"`
}