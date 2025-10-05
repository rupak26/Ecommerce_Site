package user_handler

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateUserPatialById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id") 
	numID, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var req repo.UpdateUserReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if req.Name == nil && req.Age == nil && req.Email == nil && req.Password == nil && req.Occupation == nil {
		utils.WriteResponse(w, http.StatusBadRequest, "No fields provided for update")
		return
	}

	updatedUser, err := h.userRepo.PatchUser(numID, req)
	if err != nil {
		utils.WriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.WriteResponse(w, http.StatusOK, updatedUser)
}
