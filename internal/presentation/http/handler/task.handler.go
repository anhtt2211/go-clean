package handler

import (
	"My-Clean/internal/application/use-cases"
	"My-Clean/internal/domain/types"
	"My-Clean/internal/presentation/http/inputs"
	"My-Clean/internal/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TaskHandler struct {
	TaskUseCase *use_cases.TaskUseCase
}

func NewTaskHandler(taskUseCase *use_cases.TaskUseCase) *TaskHandler {
	return &TaskHandler{TaskUseCase: taskUseCase}
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Fetch task by ID from use case
	task, err := h.TaskUseCase.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Failed to retrieve task: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if task was found
	if task == nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	// Respond with task data
	utils.RespondWithJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters into PaginateOptions
	options := types.PaginateOptions{}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid query parameters", http.StatusBadRequest)
		return
	}

	if val, ok := r.Form["keyword"]; ok && len(val) > 0 {
		options.Keyword = val[0]
	}
	if val, ok := r.Form["limit"]; ok && len(val) > 0 {
		limit, err := strconv.Atoi(val[0])
		if err != nil {
			http.Error(w, "Invalid limit value", http.StatusBadRequest)
			return
		}
		options.Limit = limit
	}
	if val, ok := r.Form["page"]; ok && len(val) > 0 {
		page, err := strconv.Atoi(val[0])
		if err != nil {
			http.Error(w, "Invalid page value", http.StatusBadRequest)
			return
		}
		options.Page = page
	}
	if val, ok := r.Form["order_by"]; ok && len(val) > 0 {
		options.OrderBy = val[0]
	}
	if val, ok := r.Form["order"]; ok && len(val) > 0 {
		options.Order = val[0]
	}
	if val, ok := r.Form["filters"]; ok && len(val) > 0 {
		filters := make(map[string]interface{})
		err := json.Unmarshal([]byte(val[0]), &filters)
		if err != nil {
			http.Error(w, "Invalid filters value", http.StatusBadRequest)
			return
		}
		options.Filters = filters
	}

	// Fetch tasks from use case
	tasks, err := h.TaskUseCase.Gets(options)
	if err != nil {
		http.Error(w, "Failed to retrieve tasks: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with tasks data
	utils.RespondWithJSON(w, http.StatusOK, tasks)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var taskInput inputs.CreateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&taskInput); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	// Convert CreateTaskInput to CreateTaskDto
	taskDto, err := taskInput.ToCreateTaskDto()
	if err != nil {
		http.Error(w, "Invalid input data: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Call the use case to create the task
	err = h.TaskUseCase.Create(taskDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the created task DTO
	utils.RespondWithJSON(w, http.StatusCreated, taskDto)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Decode the JSON request body into an UpdateTaskInput
	var taskInput inputs.UpdateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&taskInput); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	// Convert UpdateTaskInput to UpdateTaskDto
	taskDto, err := taskInput.ToUpdateTaskDto()
	if err != nil {
		http.Error(w, "Invalid input data: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Set the ID from the URL parameters
	taskDto.ID = uint(id)

	// Call the use case to update the task
	err = h.TaskUseCase.Update(taskDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the updated task DTO
	utils.RespondWithJSON(w, http.StatusOK, taskDto)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	err = h.TaskUseCase.Delete(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
