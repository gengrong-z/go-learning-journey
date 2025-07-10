package battlelog

import "fmt"

type LogEntry struct {
	Round  int
	Actor  string
	Action string
	Target string
	Effect string
}

type BattleLog struct {
	Entries []LogEntry
}

func NewBattleLog() *BattleLog {
	return &BattleLog{make([]LogEntry, 0)}
}

func (b *BattleLog) Add(entry LogEntry) {
	b.Entries = append(b.Entries, entry)
}

func (b *BattleLog) Print() {
	currentRound := 0
	for _, entry := range b.Entries {
		if currentRound != entry.Round {
			currentRound = entry.Round
			fmt.Printf("🌀 Round %d\n", currentRound)
		}
		fmt.Printf("  %s %s → %s (%s)\n", entry.Actor, entry.Action, entry.Target, entry.Effect)
	}
}
