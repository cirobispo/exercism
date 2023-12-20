package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Player struct {
	name string
	wins int
	draws int
	loss int
}

type GameResult int 
const (
	GRLoss GameResult = 0
	GRDraw GameResult = 1
	GRWin GameResult = 3
	GRNone GameResult = -1
)

func textToGameResult(gr string) GameResult {
	switch strings.ToLower(gr) {
	case "win": return GRWin
	case "loss": return GRLoss
	case "draw": return GRDraw
	default:
		return GRNone
	}
}

func updateScore(players map[string]Player, playerNameA, playerNameB string, gr GameResult) {
	getPlayer:=func(playerName string) Player {
		p, found:=players[playerName]
		if !found {
			p=Player{ name: playerName}
			players[playerName]=p
		}
		
		return p
	}

	updatePlayer:=func(player Player) {
		players[player.name]=player
	}

	playerA:=getPlayer(playerNameA)
	playerB:=getPlayer(playerNameB)

	switch gr {
		case GRWin:
			playerA.wins++
			playerB.loss++
		case GRLoss:
			playerA.loss++
			playerB.wins++
		default:
			playerA.draws++
			playerB.draws++
	}

	updatePlayer(playerA)
	updatePlayer(playerB)
}

func readScore(reader io.Reader) (map[string]Player, error) {
	players:=make(map[string]Player)

	var err error
	scanner:=bufio.NewScanner(reader)
	for scanner.Scan() {
		line:=scanner.Text()
		if len(line) > 0 {
			data:=strings.Split(line, ";")
			if ss:=len(data); ss == 1 && data[0][0] == '#' {
				continue
			} else if (ss < 3 || (ss == 3 && textToGameResult(data[2]) == GRNone)) {
				err=errors.New("wrong data")
				break
			} else {
				updateScore(players, data[0], data[1], textToGameResult(data[2]))
			}
		}
	}

	return players, err
}

func writeScore(players []Player, writer io.Writer) {
	fillSpace:=func(text string, size int, onLeft bool) string {
		spaces:=strings.Repeat(" ", size - len(text))
		if onLeft {
			return spaces + text
		}

		return text + spaces
	}
	header:=fmt.Sprintf("%s | MP |  W |  D |  L |  P\n", fillSpace("Team", 30, false))
	fmt.Fprint(writer, header)
	for i:=range players {
		p:=players[i]
		line:=fmt.Sprintf("%s | %s | %s | %s | %s | %s\n", 
			fillSpace(p.name, 30, false),
			fillSpace(fmt.Sprintf("%d", p.wins + p.draws + p.loss), 2, true),
			fillSpace(fmt.Sprintf("%d", p.wins), 2, true),
			fillSpace(fmt.Sprintf("%d", p.draws), 2, true),
			fillSpace(fmt.Sprintf("%d", p.loss), 2, true),
			fillSpace(fmt.Sprintf("%d", p.wins * 3 + p.draws), 2, true),
		) 
		fmt.Fprint(writer, line)
	}
}

func sortPlayers(players map[string]Player) []Player {
	sp:=make([]Player, 0, len(players))
	for k := range players {
		sp = append(sp, players[k])
	}	

	sorttingData:=func(i, j int) bool {
		getScore:=func(p Player) int {
			return (p.wins * 3) + p.draws
		}

		sa, sb:=getScore(sp[i]), getScore(sp[j])
		if sa == sb {
			return strings.Compare(sp[j].name, sp[i].name) == 1
		}

		return sa > sb
	}
	sort.Slice(sp, sorttingData)

	return sp
}

func Tally(reader io.Reader, writer io.Writer) error {
	var players map[string]Player
	var err error
	if players, err=readScore(reader); err == nil {
		data:=sortPlayers(players)
		writeScore(data, writer)	
	} else {
		fmt.Print(players)
	}

	return err
}
