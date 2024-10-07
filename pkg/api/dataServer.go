package api

import (
	sync "sync"

	random "github.com/bbitere/gameapi.git/pkg/random"
)


type DataServer struct{

	LogicMutex					sync.Mutex
	GameLogic				*CrashGameLogic
	Random					random.IRandom
	//dbCtx					*db.DatabaseContext
}

var API *DataServer = nil

func (This *DataServer) constr() *DataServer{

	This.Random    = (new (random.Mokup_Random)).Constr();
	This.GameLogic = (new (CrashGameLogic)).Constr();
	
	return This;
}

func DataServer_init() *DataServer{

	var server = (new (DataServer)).constr();
	API = server;

	server.GameLogic.initGame(server)

	return server;
}

