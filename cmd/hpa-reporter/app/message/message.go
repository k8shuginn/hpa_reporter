package message

const (
	LevelWarning  = "warning"
	LevelCritical = "critical"
)

// Data is message data
type Data struct {
	Time            string
	Level           string
	Name            string
	Namespace       string
	CurrentReplicas int32
	MaxReplicas     int32
}
