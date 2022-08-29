package models

// Article model
type Visit struct {
	ID        int    `json:"id" binding:"required"`
	Type      int    `json:"type" binding:"required"`
	Ip        string `json:"ip"  binding:"required"`
	VisitTime string `json:"visit_time"`
}
