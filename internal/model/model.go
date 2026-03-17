package model

type Profile struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

type Todo struct {
	ID        int    `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
}

type TodoResponse struct {
	Todos []Todo `json:"todos"`
}

type DashboardResponse struct {
	ID               int     `json:"id"`
	FullName         string  `json:"full_name"`
	Status           string  `json:"status"`
	PendingTaskCount int     `json:"pending_task_count"`
	NextUrgentTask   string  `json:"next_urgent_task"`
	ErrorWarning     *string `json:"error_warning"`
}