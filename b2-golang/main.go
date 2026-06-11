package main

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin" 
)

// Model sesuai spesifikasi soal
type LeaveRequest struct {
	ID         int    `json:"id"`
	EmployeeID string `json:"employee_id"`
	StartDate  string `json:"start_date"` // format: YYYY-MM-DD
	EndDate    string `json:"end_date"`
	Reason     string `json:"reason"`
	Status     string `json:"status"` // "Pending", "Approved", "Rejected"
}

var (
	leaveRequests = []LeaveRequest{}
	currentID     = 1
	mu            sync.Mutex // Menghindari race condition pada in-memory slice
)

func main() {
	r := gin.Default()

	// Routing Endpoint
	r.POST("/leave", createLeave)
	r.GET("/leave", getAllLeave)
	r.PUT("/leave/:id/approve", approveLeave)
	r.PUT("/leave/:id/reject", rejectLeave)
	r.GET("/leave/employee/:emp_id", getEmployeeLeave)

	// Server berjalan di port 8080
	r.Run(":8080")
}

// 1. POST /leave - Mengajukan cuti baru dengan validasi ketat
func createLeave(c *gin.Context) {
	var req LeaveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid"})
		return
	}

	layout := "2006-01-02"
	start, errStart := time.Parse(layout, req.StartDate)
	end, errEnd := time.Parse(layout, req.EndDate)
	
	// Ambil tanggal hari ini tanpa jam/menit/detik
	todayStr := time.Now().Format(layout)
	today, _ := time.Parse(layout, todayStr)

	// Validasi Bisnis 1: Format tanggal
	if errStart != nil || errEnd != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format tanggal harus YYYY-MM-DD"})
		return
	}
	// Validasi Bisnis 2: Start date tidak boleh lewat dari hari ini
	if start.Before(today) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date tidak boleh lebih awal dari hari ini"})
		return
	}
	// Validasi Bisnis 3: End date tidak boleh sebelum start date
	if end.Before(start) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "end_date tidak boleh sebelum start_date"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// Validasi Bisnis 4: Cek Overlap (Tanggal Tabrakan)
	for _, lr := range leaveRequests {
		if lr.EmployeeID == req.EmployeeID && lr.Status != "Rejected" {
			sExisting, _ := time.Parse(layout, lr.StartDate)
			eExisting, _ := time.Parse(layout, lr.EndDate)

			// Rumus Overlap: Start1 <= End2 DAN End1 >= Start2
			if !start.After(eExisting) && !end.Before(sExisting) {
				c.JSON(http.StatusConflict, gin.H{
					"error": "karyawan " + req.EmployeeID + " sudah memiliki pengajuan cuti yang overlap pada tanggal tersebut",
				})
				return
			}
		}
	}

	// Assign ID dan default status
	req.ID = currentID
	currentID++
	req.Status = "Pending"
	leaveRequests = append(leaveRequests, req)

	c.JSON(http.StatusCreated, req)
}

// 2. GET /leave - Mengambil semua list cuti
func getAllLeave(c *gin.Context) {
	c.JSON(http.StatusOK, leaveRequests)
}

// Helper untuk Update Status (Approve/Reject)
func updateStatus(c *gin.Context, targetStatus string) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, lr := range leaveRequests {
		if lr.ID == id {
			// Validasi Bisnis 5: Jika sudah Approved/Rejected tidak bisa diubah lagi
			if lr.Status == "Approved" || lr.Status == "Rejected" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Pengajuan sudah diproses sebelumnya dengan status: " + lr.Status})
				return
			}
			leaveRequests[i].Status = targetStatus
			c.JSON(http.StatusOK, leaveRequests[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Pengajuan cuti tidak ditemukan"})
}

func approveLeave(c *gin.Context) { updateStatus(c, "Approved") }
func rejectLeave(c *gin.Context)  { updateStatus(c, "Rejected") }

// 5. GET /leave/employee/:emp_id - List cuti per karyawan
func getEmployeeLeave(c *gin.Context) {
	empID := c.Param("emp_id")
	res := []LeaveRequest{}
	for _, lr := range leaveRequests {
		if lr.EmployeeID == empID {
			res = append(res, lr)
		}
	}
	c.JSON(http.StatusOK, res)
}