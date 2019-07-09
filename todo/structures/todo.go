package structures

type MicroMessage struct {
	Message string    `json:"message"`
	To      string    `json:"to"`
	From    string    `json:"from"`
	TimeToLifeSec int `json:"timeToLifeSec"`
}
