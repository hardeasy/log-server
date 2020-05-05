package dto
type Index struct {
	Health string `json:"health"`
	Status string `json:"status"`
	Index string `json:"index"`
	Uuid string `json:"uuid"`
	Pri int `json:"pri"`
	Rep int `json:"rep"`
	DocsCount int `json:"docs_count"`
	DocsDeleted int `json:"docs_deleted"`
	StoreSize string `json:"store_size"`
	PriStoreSize string `json:"pri_store_size"`
}
