package product_handler

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProductPartialiById(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id") 
	numID, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var req domain.UpdateProductReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if req.ProductName == nil && req.Url == nil && req.Quantity == nil {
		utils.WriteResponse(w, http.StatusBadRequest, "No fields provided for update")
		return
	}

	updatedUser, err := h.svc.PatchProduct(numID, req)
	if err != nil {
		utils.WriteResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.WriteResponse(w, http.StatusOK, updatedUser)
}
