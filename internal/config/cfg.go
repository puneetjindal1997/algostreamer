// Copyright (C) 2022 AlgoNode Org.
//
// algostreamer is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// algostreamer is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with algostreamer.  If not, see <https://www.gnu.org/licenses/>.

package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/algonode/algostreamer/constant"
	"github.com/algonode/algostreamer/internal/algod"
	"github.com/algonode/algostreamer/internal/rdb"
	"github.com/algonode/algostreamer/internal/rego"
)

var cfgFile = flag.String("f", "config.jsonc", "config file")
var firstRound = flag.Int64("r", -1, "first round to start [-1 = latest]")
var lastRound = flag.Int64("l", -1, "last round to read [-1 = no limit]")
var simpleFlag = flag.Bool("s", true, "simple mode - just sending blocks in JSON format to stdout")

type SinksCfg struct {
	Redis *rdb.RedisConfig `json:"redis"`
}

type SteramerConfig struct {
	Algod  *algod.AlgoConfig `json:"algod"`
	Sinks  SinksCfg          `json:"sinks"`
	Rego   *rego.OpaConfig   `json:"opa"`
	Stdout bool              `json:"stdout"`
}

var defaultConfig = SteramerConfig{}

// loadConfig loads the configuration from the specified file, merging into the default configuration.
func LoadConfig() (cfg SteramerConfig, err error) {
	var algoNode []*algod.AlgoNodeConfig
	var algodVar algod.AlgoConfig
	var algoRedis rdb.RedisConfig
	var OPA rego.OpaConfig
	flag.Parse()
	cfg = defaultConfig
	algoNode = append(algoNode, &algod.AlgoNodeConfig{Address: os.Getenv("NodeHost"), Id: os.Getenv("NodeType"), Token: "", LastBlockKey: os.Getenv("LastBlockPublic")})
	algoNode = append(algoNode, &algod.AlgoNodeConfig{Address: os.Getenv("PrivateHost"), Id: os.Getenv("NodeType"), Token: os.Getenv("Token"), LastBlockKey: os.Getenv("LastBlockPrivate")})
	// algoNode = append(algoNode, &algod.AlgoNodeConfig{Address: constant.NodeHost, Id: constant.NodeType, Token: ""})
	// algoNode = append(algoNode, &algod.AlgoNodeConfig{Address: "http://3.15.228.205:28280", Id: "public-node", Token: "EXRPIUPWOFLWMOTLKVWTOTERBYITVSKUOULLFXKNZEOGYCTNQKEIWVGRKFCMMGFJ"})
	algodVar.ANodes = algoNode
	cfg.Algod = &algodVar
	algoRedis.Addr = os.Getenv("RedisHost") + ":" + os.Getenv("RedisPort")
	// algoRedis.Addr = constant.RedisHost + ":" + constant.RedisPort
	algoRedis.DB = 0
	algoRedis.Password = ""
	algoRedis.Username = ""
	cfg.Sinks.Redis = &algoRedis
	OPA.MyID = constant.MyId
	OPA.Rules.Block = constant.RuleBlock
	cfg.Rego = &OPA
	// err = utils.LoadJSONCFromFile(*cfgFile, &cfg)

	if cfg.Algod == nil {
		return cfg, fmt.Errorf("[CFG] Missing algod config")
	}
	if len(cfg.Algod.ANodes) == 0 {
		return cfg, fmt.Errorf("[CFG] Configure at least one node")
	}
	cfg.Algod.FRound = *firstRound
	cfg.Algod.LRound = *lastRound
	cfg.Stdout = *simpleFlag

	return cfg, err
}
