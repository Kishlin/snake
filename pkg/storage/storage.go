package storage

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/kishlin/snake/v2/pkg/game"
)

const Filename = "leaderboard.data"
const MaxEntries = 10

type Storage struct {
	rootDir string
}

func (s *Storage) Init() error {
	ex, err := os.Executable()
	if err != nil {
		return errors.New(fmt.Sprintf("getting executable path: %v", err))
	}

	s.rootDir = filepath.Dir(ex)

	return nil
}

func (s *Storage) Write(entry game.LeaderboardEntry) error {
	entries, err := s.ReadTop()
	if err != nil {
		return errors.New(fmt.Sprintf("reading top entries: %v", err))
	}

	entries = append(entries, entry)

	sort.Slice(
		entries,
		func(i, j int) bool {
			return entries[i].Score > entries[j].Score
		},
	)

	if len(entries) > MaxEntries {
		entries = entries[:MaxEntries]
	}

	compressed, err := compressLeaderboard(entries)
	if err != nil {
		return errors.New(fmt.Sprintf("compressing leaderboard: %v", err))
	}

	err = os.WriteFile(filepath.Join(s.rootDir, Filename), compressed, 0644)
	if err != nil {
		return errors.New(fmt.Sprintf("writing leaderboard: %v", err))
	}

	return nil
}

func (s *Storage) ReadTop() ([]game.LeaderboardEntry, error) {
	fileContent, err := os.ReadFile(filepath.Join(s.rootDir, Filename))
	if err != nil {
		if os.IsNotExist(err) {
			err := os.WriteFile(filepath.Join(s.rootDir, Filename), []byte{}, 0644)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("creating file: %v", err))
			}

			return []game.LeaderboardEntry{}, nil
		}

		return nil, errors.New(fmt.Sprintf("reading file: %v", err))
	}

	if len(fileContent) == 0 {
		return []game.LeaderboardEntry{}, nil
	}

	entries, err := uncompressLeaderboard(fileContent)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("uncompressing leaderboard: %v", err))
	}

	return entries, nil
}

func compressLeaderboard(entries []game.LeaderboardEntry) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(entries)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func uncompressLeaderboard(data []byte) ([]game.LeaderboardEntry, error) {
	var entries []game.LeaderboardEntry

	dec := gob.NewDecoder(bytes.NewBuffer(data))

	err := dec.Decode(&entries)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("decoding: %v", err))
	}

	return entries, nil
}
