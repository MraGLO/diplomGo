package model

type PricingRecord struct {
	ID                  int `json:"id"`
	TeacherID           int `json:"teacherid"`
	GroupID             int `json:"groupid"`
	FirstHalfYear       int `json:"firsthalfyear"`
	SecondHalfYear      int `json:"secondhalfyear"`
	Theory              int `json:"theory"`
	Practice            int `json:"practice"`
	LPZ1                int `json:"lpz1"`
	LPZ2                int `json:"lpz2"`
	Consultation        int `json:"consultation"`
	CourseProject       int `json:"courseproject"`
	HoursFirstSemester  int `json:"hoursfirstsemester"`
	HoursSecondSemester int `json:"hourssecondsemester"`
	Total               int `json:"total"`
}
