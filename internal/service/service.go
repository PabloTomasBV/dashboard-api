package service

import (
	"dashboard-api/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type DashboardService struct {
	client  *http.Client
	baseURL string
}

func NewDashboardService(client *http.Client, baseURL string) *DashboardService {
	return &DashboardService{
		client:  client,
		baseURL: baseURL,
	}
}

func (s *DashboardService) GetDashboard(id string) (*model.DashboardResponse, error) {
	var profile *model.Profile
	var todo *model.TodoResponse

	var profileErr error
	var todoErr error

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		profile, profileErr = s.getUser(id)
	}()

	go func() {
		defer wg.Done()
		todo, todoErr = s.getTodo(id)
	}()

	wg.Wait()
	var warning *string

	if profileErr != nil {
		return nil, fmt.Errorf("Error fetching profile: %w", profileErr)
	}

	if todoErr != nil {
		msg := "Todos Unavailable"
		warning = &msg
	}

	result := buildDashboard(profile, todo)
	result.ErrorWarning = warning

	return &result, nil
}

func (s *DashboardService) getUser(id string) (*model.Profile, error) {
	url := fmt.Sprintf("%s/users/%s", s.baseURL, id)

	resp, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user model.Profile

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *DashboardService) getTodo(id string) (*model.TodoResponse, error) {
	url := fmt.Sprintf("%s/todos/user/%s", s.baseURL, id)

	resp, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var todos model.TodoResponse

	if err := json.NewDecoder(resp.Body).Decode(&todos); err != nil {
		return nil, err
	}

	return &todos, nil
}

func buildDashboard(profile *model.Profile, todo *model.TodoResponse) model.DashboardResponse {
	fullName := profile.FirstName + " " + profile.LastName
	pendingCount := 0
	nextTask := ""

	status := "Rookie"
	if profile.Age > 50 {
		status = "Veteran"
	}

	if todo != nil {
		for _, t := range todo.Todos {
			if !t.Completed {
				pendingCount++
				if nextTask == "" {
					nextTask = t.Todo
				}
			}
		}
	}

	return model.DashboardResponse{
		ID:               profile.ID,
		FullName:         fullName,
		Status:           status,
		PendingTaskCount: pendingCount,
		NextUrgentTask:   nextTask,
	}
}
