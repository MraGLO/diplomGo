package model

type PricingRecord struct {
	ID                  int    `json:"ID"`
	GroupID             int    `json:"groupID"`
	FirstHalfYear       int    `json:"firstHalfYear"`
	SecondHalfYear      int    `json:"secondHalfYear"`
	Theory              int    `json:"theory"`
	Practice            int    `json:"practice"`
	LPZ1                int    `json:"lpz1"`
	LPZ2                int    `json:"lpz2"`
	Consultation        int    `json:"consultation"`
	CourseProject       int    `json:"courseProject"`
	HoursFirstSemester  string `json:"hoursFirstSemester"`
	HoursSecondSemester string `json:"hoursSecondSemester"`
	Total               int    `json:"total"`
	PricingTableID      int    `json:"pricingTableID"`
	SubjectID           int    `json:"subjectID"`
}
