package farmsage

type Animal struct {
	AnimalID    int    `json:"id"`
	Name        string `json:"name"`
	Sex         string `json:"sex"`
	Age         int    `json:"age"`
	MedicalInfo string `json:"medicalinfo"`
}

type Activity struct {
	ActivityID   int    `json:"id"`
	AnimalID     int    `json:"animalid"`
	ActivityType string `json:"activitytype"`
	StartTime    string `json:"starttime"`
	EndTime      string `json:"endtime"`
	Latitude     int    `json:"latitude"`
	Longitude    int    `json:"longitude"`
}

type FeedingSchedule struct {
	ScheduleID  int    `json:"id"`
	AnimalID    int    `json:"animalid"`
	FeedID      int    `json:"feedid"`
	FeedingTime string `json:"feedingtime"`
}

type Feed struct {
	FeedID   int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type Biometrics struct {
	BiometricID   int     `json:"id"`
	AnimalID      int     `json:"animalid"`
	Pulse         int     `json:"pulse"`
	Temperature   float64 `json:"temperature"`
	Weight        float64 `json:"weight"`
	BreathingRate int     `json:"breathingrate"`
}

type Farm struct {
	FarmID   int    `json:"id"`
	AnimalID int    `json:"animalid"`
	UserID   int    `json:"userid"`
	Name     string `json:"name"`
}
