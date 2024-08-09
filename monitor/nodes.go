package monitor

import (
	"blog/services/dto"
	"fmt"
	"strings"
)

type Monitor struct {
}

type Block struct {
	Height  int64  `json:"height"`
	Miner   string `json:"miner"`
	Message string `json:"cid"`
}

func ReadToBean(content string) dto.MinerStatus {
	fileName := "lotusminer.txt"
	writeRel := WriteToFile(fileName, content)
	var ms dto.MinerStatus
	if writeRel > 0 {
		lines, err := ReadLines(fileName)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return ms
		}
		m := map[string]string{}
		for _, v := range lines {
			v = strings.TrimSpace(v)
			kv := strings.Split(v, ":")
			if _, ok := m[kv[0]]; !ok && len(kv) > 1 {
				kv[1] = strings.TrimSpace(kv[1])
				if kv[0] == "Miner" {
					mm := strings.Split(kv[1], " (")
					m[kv[0]] = mm[0]
					m["SectorSize"] = mm[1][0:6]
				} else {
					m[kv[0]] = strings.ReplaceAll(strings.ReplaceAll(kv[1], "[", ""), "]", "")
				}
			}
		}
		ms = dto.MinerStatus{
			StartTime:    m["StartTime"],
			Chain:        m["Chain"],
			Miner:        m["Miner"],
			Power:        strings.Split(strings.Split(m["Power"], "/")[0], " ")[0],
			Raw:          strings.Split(strings.Split(m["Raw"], "/")[0], " ")[0],
			Balance:      strings.Split(m["Miner Balance"], " ")[0],
			Pledge:       strings.Split(m["Pledge"], " ")[0],
			Vesting:      strings.Split(m["Vesting"], " ")[0],
			Available:    strings.Split(m["Available"], " ")[0],
			Beneficiary:  m["Beneficiary"],
			SectorsTotal: m["Total"],
			SectorSize:   m["SectorSize"],
		}
		fmt.Printf("MinerStatus: %T\n", ms)
	}
	return ms
}
