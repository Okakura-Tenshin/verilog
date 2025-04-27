package models

import (
	"gorm.io/gorm"
	"time"
    "errors"
)

type Experiment struct {
    gorm.Model
    ExperimentName        string    `gorm:"type:varchar(255);not null"`     // 明确指定长度
    ExperimentDescription string    `gorm:"not null"`
    ExperimentStatus      bool      `gorm:"not null"`
    ExperimentStartTime   time.Time `gorm:"not null"`
    ExperimentEndTime     time.Time `gorm:"not null"`
    ExperimentCourseName  string    `gorm:"type:varchar(255);not null"`    // 与外键同长度
    FilePath              string    `gorm:"not null"`
    ExperimentCourse      Course    `gorm:"foreignKey:ExperimentCourseName;references:CourseName"`
}

func InitExperiment(db *gorm.DB) error{
	return db.AutoMigrate(&Experiment{})
}

type SubmitRecord struct{
    gorm.Model
    SubmitID uint `gorm:"not null"`
    SubmitName string `gorm:"type:varchar(255);not null"`
    ExperimentName string `gorm:"type:varchar(255);not null"`
    CourseName string `gorm:"type:varchar(255);not null"`
    SubmitStatus uint `gorm:"not null"`//设计三个状态，0是没提交过，1是交了但是错了，2是通过
    StudentID User `gorm:"foreignKey:SubmitID;references:ID"`
    Course Course `gorm:"foreignKey:CourseName;references:CourseName"`
}

func InitSubmitRecord(db *gorm.DB) error{
    return db.AutoMigrate(&SubmitRecord{})
}



type SubmitRecord_pass struct{
    SubmitStatus uint 
}

func AddSubmitRecord(db *gorm.DB, studentID uint, submitName string, experimentName string,courseName string,SubmitStatus int) error {
    record := SubmitRecord{}
    err := db.Where("submit_id = ? AND experiment_name = ?", studentID, experimentName).First(&record).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        // 插入新记录
        record = SubmitRecord{
            SubmitID:       studentID,
            SubmitName:     submitName,
            CourseName:     courseName,
            ExperimentName: experimentName,
            SubmitStatus:   0,
        }
        return db.Create(&record).Error
    } else if err != nil {
        return err
    }
    if SubmitStatus == 1{
    return db.Model(&SubmitRecord{}).
        Where("submit_id = ? AND experiment_name = ?", studentID, experimentName).
        Updates(&SubmitRecord_pass{
            SubmitStatus: 2,
        }).Error
        }else if SubmitStatus == 0{
        return db.Model(&SubmitRecord{}).
        Where("submit_id = ? AND experiment_name = ?", studentID, experimentName).
        Updates(&SubmitRecord_pass{
            SubmitStatus: 1,
        }).Error
    }
    return nil
}
type SubmitInfo struct {
	SubmitID uint `json:"submit_id"`
	SubmitName string `json:"submit_name"`
	SubmitStatus uint `json:"submit_status"`
}	
func GetExperimentResult(db *gorm.DB, experimentName string, courseName string) ([]SubmitInfo, error) {
	var records []SubmitRecord
	err := db.Where("experiment_name = ? AND course_name = ?", experimentName, courseName).
		Find(&records).Error
    var result []SubmitInfo
    for _,record:=range records{
        result=append(result,SubmitInfo{
            SubmitID:record.SubmitID,
            SubmitName:record.SubmitName,
            SubmitStatus:record.SubmitStatus,
        })
}
    return result, err
}
