package handler

import (
	"net/http"
	"strconv"

	"github.com/deikioveca/TheRedDevilsData/api/helper"
	"github.com/deikioveca/TheRedDevilsData/api/service"
)

type Handler struct {
	service	service.Service
}


func NewHandler(s service.Service) *Handler {
	return &Handler{service: s}
}


func (h *Handler) GetCountries(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetCountries()
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	helper.WriteJSON(w, http.StatusOK, data)
}


func (h *Handler) GetCountryByName(w http.ResponseWriter, r *http.Request) {
	countryName := r.PathValue("name")

	data, err := h.service.GetCountryByName(countryName)
	if err != nil {
		switch err {
		case service.ErrCountryNotFound:
			helper.WriteError(w, http.StatusNotFound, err.Error())
		default:
			helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	helper.WriteJSON(w, http.StatusOK, data)
}


func (h *Handler) GetLeagues(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetLeagues()
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	helper.WriteJSON(w, http.StatusOK, data)
}


func (h *Handler) GetTeam(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetTeam()
	if err != nil {
		switch err {
		case service.ErrManchesterUnitedNotFound:
			helper.WriteError(w, http.StatusNotFound, err.Error())
		default:
			helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	helper.WriteJSON(w, http.StatusOK, data)
}


func (h *Handler) GetTeamStatsGames(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("season")
	season, err := strconv.Atoi(pathValue)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "incorrect path variable for 'season'")
		return
	}

	data, err := h.service.GetTeamStatsGames(season)
	if err != nil {
		switch err {
		case service.ErrTeamStatsNotFound:
			helper.WriteError(w, http.StatusNotFound, err.Error())
		default:
			helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	helper.WriteJSON(w, http.StatusOK, data)
}


func (h *Handler) GetTeamStatsGoals(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("season")
	season, err := strconv.Atoi(pathValue)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "incorrect path variable for 'season'")
		return
	}

	data, err := h.service.GetTeamStatsGoals(season)
	if err != nil {
		switch err {
		case service.ErrTeamStatsNotFound:
			helper.WriteError(w, http.StatusNotFound, err.Error())
		default:
			helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	helper.WriteJSON(w, http.StatusOK, data)
}


func (h *Handler) GetTeamStatsStreak(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("season")
	season, err := strconv.Atoi(pathValue)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "incorrect path variable for 'season'")
		return
	}

	data, err := h.service.GetTeamStatsStreak(season)
	if err != nil {
		switch err {
		case service.ErrTeamStatsNotFound:
			helper.WriteError(w, http.StatusNotFound, err.Error())
		default:
			helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	helper.WriteJSON(w, http.StatusOK, data)
}


func (h *Handler) GetTeamStatsBiggest(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("season")
	season, err := strconv.Atoi(pathValue)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "incorrect path variable for 'season'")
		return
	}

	data, err := h.service.GetTeamStatsBiggest(season)
	if err != nil {
		switch err {
		case service.ErrTeamStatsNotFound:
			helper.WriteError(w, http.StatusNotFound, err.Error())
		default:
			helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	helper.WriteJSON(w, http.StatusOK, data)
}


func (h *Handler) GetTeamStatsCleanSheet(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("season")
	season, err := strconv.Atoi(pathValue)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "incorrect path variable for 'season'")
		return
	}

	data, err := h.service.GetTeamStatsCleanSheet(season)
	if err != nil {
		switch err {
		case service.ErrTeamStatsNotFound:
			helper.WriteError(w, http.StatusNotFound, err.Error())
		default:
			helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	helper.WriteJSON(w, http.StatusOK, data)
}


func (h *Handler) GetTeamStatsFailedToScore(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("season")
	season, err := strconv.Atoi(pathValue)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "incorrect path variable for 'season'")
		return
	}

	data, err := h.service.GetTeamStatsFailedToScore(season)
	if err != nil {
		switch err {
		case service.ErrTeamStatsNotFound:
			helper.WriteError(w, http.StatusNotFound, err.Error())
		default:
			helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	helper.WriteJSON(w, http.StatusOK, data)
}


func (h *Handler) GetTeamStatsPenalty(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("season")
	season, err := strconv.Atoi(pathValue)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "incorrect path variable for 'season'")
		return
	}

	data, err := h.service.GetTeamStatsPenalty(season)
	if err != nil {
		switch err {
		case service.ErrTeamStatsNotFound:
			helper.WriteError(w, http.StatusNotFound, err.Error())
		default:
			helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	helper.WriteJSON(w, http.StatusOK, data)
}


func (h *Handler) GetTeamStatsCards(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("season")
	season, err := strconv.Atoi(pathValue)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "incorrect path variable for 'season'")
		return
	}

	data, err := h.service.GetTeamStatsCards(season)
	if err != nil {
		switch err {
		case service.ErrTeamStatsNotFound:
			helper.WriteError(w, http.StatusNotFound, err.Error())
		default:
			helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	helper.WriteJSON(w, http.StatusOK, data)
}


func (h *Handler) GetTeamStatsLineups(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("season")
	season, err := strconv.Atoi(pathValue)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "incorrect path variable for 'season'")
		return
	}

	data, err := h.service.GetTeamStatsLineup(season)
	if err != nil {
		switch err {
		case service.ErrTeamStatsNotFound, service.ErrLineupNotFound:
			helper.WriteError(w, http.StatusNotFound, err.Error())
		default:
			helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	helper.WriteJSON(w, http.StatusOK, data)
}


func (h *Handler) GetVenues(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetVenues()
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	helper.WriteJSON(w, http.StatusOK, data)
}


func(h *Handler) GetVenuesByCity(w http.ResponseWriter, r *http.Request) {
	city := r.PathValue("city")

	data, err := h.service.GetVenuesByCity(city)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	helper.WriteJSON(w, http.StatusOK, data)
}


func (h *Handler) GetVenuesBiggestAndSmallest(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetVenuesBiggestAndSmallest()
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	helper.WriteJSON(w, http.StatusOK, data)
}


func (h *Handler) GetStandingsBySeason(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("season")
	season, err := strconv.Atoi(pathValue)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "incorrect path variable for 'season'")
		return
	}

	data, err := h.service.GetStandingsBySeason(season)
	if err != nil {
		switch err {
		case service.ErrStandingNotFound:
			helper.WriteError(w, http.StatusNotFound, err.Error())
		default:
			helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	helper.WriteJSON(w, http.StatusOK, data)
}


func (h *Handler) GetFixturesBySeason(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("season")
	season, err := strconv.Atoi(pathValue)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "incorrect path variable for 'season'")
		return
	}

	data, err := h.service.GetFixturesBySeason(season)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	helper.WriteJSON(w, http.StatusOK, data)
}


func (h *Handler) GetInjuriesBySeason(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("season")
	season, err := strconv.Atoi(pathValue)
	if err != nil {
		helper.WriteError(w, http.StatusBadRequest, "incorrect path variable for 'season'")
		return
	}

	data, err := h.service.GetInjuriesBySeason(season)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	helper.WriteJSON(w, http.StatusOK, data)
}


func (h *Handler) GetSquad(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetSquad()
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	helper.WriteJSON(w, http.StatusOK, data)
}