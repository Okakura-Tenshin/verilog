package models

import(
	"gorm.io/gorm"
)

type Course struct {
    gorm.Model
    CourseName string `gorm:"type:varchar(255);not null;unique"` // 添加唯一约束和类型
    Teachers   string `gorm:"not null"`
}


type SelectCourse struct {
	gorm.Model
	CourseName    string `gorm:"not null"`
	StudentID     uint `gorm:"not null"`
	StudentStatus bool
	Student       User   `gorm:"foreignKey:StudentID;references:ID"` // 添加关联
}

func CreateCourse(db *gorm.DB,course *Course) error{
	return db.Create(course).Error
}


func InitCourse(db *gorm.DB) error{
	return db.AutoMigrate(&Course{},&SelectCourse{})
}

type StudentInfo struct {
	StudentID uint `json:"student_id"`
	Username  string `json:"username"`
}	

func GetStudentsByCourse(db *gorm.DB, courseName string) ([]StudentInfo, error) {
	var students []SelectCourse
	err := db.Preload("Student"). // 预加载关联的User数据
		Where("course_name = ? AND student_status = ?", courseName, true).
		Find(&students).Error
	
	// 格式化返回数据，只返回需要的字段
	var result []StudentInfo
	for _, student := range students {
		result = append(result, StudentInfo{
			StudentID: student.StudentID,
			Username:  student.Student.Username,
		})
	}
	return result, err
}

// 定义一个新的结构体，只包含需要的字段
type CourseResponse struct {
    CourseName string
}

func GetCoursesByTeacher(db *gorm.DB, username string) ([]CourseResponse, error) {
    var courses []CourseResponse
    err := db.Model(&Course{}).
        Select("course_name").
        Where("teachers = ?", username).
        Find(&courses).Error
    return courses, err
}

func GetCourseByName(db *gorm.DB,courseName string) error{
	return db.Model(&Course{}).
	Where("course_name = ?", courseName).
	First(&Course{}).Error
}

func SQLApproveJoinCourse(db *gorm.DB,courseName string,studentID uint) error{
	return db.Model(&SelectCourse{}).
	Where("course_name = ? AND student_id = ?", courseName, studentID).
	Update("student_status", true).Error
}

func GetPendingStudents(db *gorm.DB,courseName string) ([]StudentInfo, error){
	var students []SelectCourse
	err := db.Preload("Student").
	Where("course_name = ? AND student_status = ?", courseName, false).
	Find(&students).Error
	
	var result []StudentInfo
	for _, student := range students {
		result = append(result, StudentInfo{
			StudentID: student.StudentID,
			Username:  student.Student.Username,
		})
	}
	return result, err	
}

func JoinCourse(db *gorm.DB, courseName string, studentID uint) error {
	selectCourse := SelectCourse{
		CourseName:    courseName,
		StudentID:     studentID,
		StudentStatus: false, // 等价于 status = 0
	}
	return db.Create(&selectCourse).Error
}

