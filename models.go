package farmsage

type Animal struct {
	AnimalID    int    `json:"id" db:"AnimalID"`
	AnimalName  string `json:"animal_name" binding:"required" db:"AnimalName"`
	Number      int    `json:"number" db:"Number"`
	DateOfBirth string `json:"date_of_birth" db:"DateOfBirth"`
	Sex         string `json:"sex" db:"Sex"`
	Age         int    `json:"age" db:"Age"`
	MedicalInfo string `json:"medical_info" db:"MedicalInfo"`
}

type Activity struct {
	ActivityID   int    `json:"id"`
	AnimalID     int    `json:"animal_id"`
	ActivityType string `json:"activity_type"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	Latitude     int    `json:"latitude"`
	Longitude    int    `json:"longitude"`
}

type FeedingSchedule struct {
	ScheduleID   int    `json:"id" db:"ScheduleID"`
	AnimalID     int    `json:"animal_id" db:"AnimalID"`
	FeedID       int    `json:"feed_id" db:"FeedID"`
	AnimalName   string `db:"AnimalName"`
	AnimalNumber int    `db:"Number"`
	FeedName     string `db:"FeedName"`
	FeedingTime  string `json:"feeding_time" db:"FeedingTime"`
}

type Feed struct {
	FeedID   int    `json:"id" db:"FeedID"`
	FeedName string `json:"FeedName" binding:"required" db:"FeedName"`
	Quantity int    `json:"Quantity" db:"Quantity"`
}

type Biometrics struct {
	BiometricID   int     `json:"id"`
	AnimalID      int     `json:"animal_id"`
	Pulse         int     `json:"pulse"`
	Temperature   float64 `json:"temperature"`
	Weight        float64 `json:"weight"`
	BreathingRate int     `json:"breathing_rate"`
}

type Farm struct {
	FarmID   int    `json:"FarmID" db:"FarmID"`
	AnimalID int    `json:"AnimalID"`
	UserID   int    `json:"UserID"`
	FarmName string `json:"FarmName" binding:"required" db:"FarmName"`
}
