package requests

type Header struct {
	KanbanContainer     int   `json:"KanbanContainer"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}
