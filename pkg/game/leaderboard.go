package game

import (
	"errors"
)

type Storage interface {
	Init() error
	Write(entry LeaderboardEntry) error
	ReadTop() ([]LeaderboardEntry, error)
}

const EntryVersion = 1

type LeaderboardEntry struct {
	Score          int
	SpeedConfig    int
	WallsAreDeadly bool
	Timestamp      int64
	Version        int
}

type Leaderboard struct {
	storage *Storage
}

func (lb *Leaderboard) Init(storage *Storage) {
	lb.storage = storage
}

func (lb *Leaderboard) Add(entry LeaderboardEntry) error {
	err := (*lb.storage).Write(entry)
	if err != nil {
		return errors.New("writing leaderboard entry: " + err.Error())
	}

	return nil
}

func (lb *Leaderboard) GetTop() ([]LeaderboardEntry, error) {
	leaderboard, err := (*lb.storage).ReadTop()
	if err != nil {
		return nil, errors.New("reading leaderboard: " + err.Error())
	}

	return leaderboard, nil
}
