package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"ez-code-run/executors"
	"ez-code-run/models"
)

func GetSubmissions(w http.ResponseWriter, r *http.Request) {
	problemIDStr := r.PathValue("id")
	problemID, err := strconv.ParseUint(problemIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid problem ID", http.StatusBadRequest)
		return
	}
	userID := r.Context().Value("user_id").(uint)
	submissions, err := models.GetSubmissionsByProblemIDAndUserID(uint(problemID), uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(submissions)
}

func CreateSubmission(w http.ResponseWriter, r *http.Request) {
	var submission models.Submission
	err := json.NewDecoder(r.Body).Decode(&submission)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	problemIDStr := r.PathValue("id")
	problemID, err := strconv.ParseUint(problemIDStr, 10, 32)
	fmt.Println(submission)
	fmt.Println(problemID)
	if err != nil {
		http.Error(w, "Invalid problem ID", http.StatusBadRequest)
		return
	}
	submission.ProblemID = uint(problemID)
	userID := r.Context().Value("user_id").(uint)
	submission.UserID = userID
	fmt.Println(submission)
	fmt.Println(userID)

	problem, err := models.GetProblemByID(submission.ProblemID)
	if err != nil {
		http.Error(w, "Problem not found", http.StatusNotFound)
		return
	}

	submission.ResultStatus = "Running"
	err = models.CreateSubmission(&submission)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ここでコードを実行
	executors.ExecuteCode(&submission, &problem)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(submission)
}
