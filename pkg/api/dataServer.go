package api

import (
	sync "sync"

	random "github.com/bbitere/gameapi.git/pkg/random"
)


type DataServer struct{

	Mutex					sync.Mutex
	GameLogic				*CrashGameLogic
	Random					random.IRandom
	//dbCtx					*db.DatabaseContext
}

var API *DataServer = nil

func (This *DataServer) constr() *DataServer{

	This.GameLogic = (new (CrashGameLogic)).Constr();
	This.Random    = (new (random.Mokup_Random)).Constr();
	return This;
}

func DataServer_init(){

	API = (new (DataServer)).constr();

}

