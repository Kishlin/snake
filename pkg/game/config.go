package game

const MaxSpeed = 9

type Config struct {
	Speed          int
	WallsAreDeadly bool
}

func (config *Config) Init() {
	config.Speed = 4
	config.WallsAreDeadly = false
}

func (config *Config) IncreaseSpeed() {
	if config.Speed < MaxSpeed {
		config.Speed += 1
	} else {
		config.Speed = 1
	}
}

func (config *Config) DecreaseSpeed() {
	if config.Speed > 1 {
		config.Speed -= 1
	} else {
		config.Speed = MaxSpeed
	}
}

func (config *Config) ToggleWallsAreDeadly() {
	config.WallsAreDeadly = !config.WallsAreDeadly
}
