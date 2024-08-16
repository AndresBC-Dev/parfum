package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AndresBC-Dev/parfum/entity"
	"github.com/AndresBC-Dev/parfum/service"
)

type PerfumeController struct {
	PerfumeService service.PerfumeService
}

func NewPerfumeController(perfumeService service.PerfumeService) *PerfumeController {
	return &PerfumeController{PerfumeService: perfumeService}
}

func (c *PerfumeController) CreatePerfumeHandler(w http.ResponseWriter, r *http.Request) {
	//ejecucion
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo no permitido: ", http.StatusBadRequest)
	}

	var parfum entity.Parfum

	err := json.NewDecoder(r.Body).Decode(&parfum)
	if err != nil {
		http.Error(w, "Error al procesar el JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	parfum.CreatedAt = time.Now()

	err = json.NewDecoder(r.Body).Decode(&parfum)
	if err != nil {
		http.Error(w, "Error al crear el perfume: "+err.Error(), http.StatusBadRequest)
		return
	}
	err = c.PerfumeService.Create(parfum)
	if err != nil {
		http.Error(w, "Error al crear el perfume: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Perfume creado exitosamente!"))
}
