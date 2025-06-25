package controllers

import (
    "hospital_system/config"
    "hospital_system/models"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

// GET /patients - List all patients
func GetPatients(c *gin.Context) {
    rows, err := config.DB.Query("SELECT id, name, age, gender, phone, address FROM patients")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
        return
    }
    defer rows.Close()

    var patients []models.Patient
    for rows.Next() {
        var p models.Patient
        rows.Scan(&p.ID, &p.Name, &p.Age, &p.Gender, &p.Phone, &p.Address)
        patients = append(patients, p)
    }
    c.JSON(http.StatusOK, patients)
}

// GET /patients/:id - Get one patient by ID
func GetPatientByID(c *gin.Context) {
    id := c.Param("id")
    var p models.Patient
    err := config.DB.QueryRow("SELECT id, name, age, gender, phone, address FROM patients WHERE id = $1", id).
        Scan(&p.ID, &p.Name, &p.Age, &p.Gender, &p.Phone, &p.Address)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
        return
    }
    c.JSON(http.StatusOK, p)
}

// PUT /patients/:id - Update patient
func UpdatePatient(c *gin.Context) {
    id := c.Param("id")
    var p models.Patient
    if err := c.ShouldBindJSON(&p); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    _, err := config.DB.Exec("UPDATE patients SET name=$1, age=$2, gender=$3, phone=$4, address=$5 WHERE id=$6",
        p.Name, p.Age, p.Gender, p.Phone, p.Address, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Patient updated"})
}

// DELETE /patients/:id - Delete patient
func DeletePatient(c *gin.Context) {
    id := c.Param("id")
    _, err := config.DB.Exec("DELETE FROM patients WHERE id=$1", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Patient deleted"})
}
