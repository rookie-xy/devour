
/*
 * Copyright (C) 2016 Meng Shi
 */


package network


import (
	"fmt"
//. "unsafe"
    . "devour/types"
)


var DvrNetworkMaxModule uint;


func DvrNetworkBlock(cf *DvrConfType, cmd *DvrCommandType, conf *DvrVoidType) string {
	/*
    var m      uint;
    //var log   *DvrErrlogType;
    var rv     DvrVoidType;
//    var conf   DvrConfType;

    DvrModules := cf.Modules.([]*DvrMoudleType);

    DvrNetworkMaxModule = 0;

    for m = 0; DvrModules[m] != nil; m++  {
	if DvrModules[m].Type != DVR_NETWORK_MODULE {
	    continue;
	}

        DvrNetworkMaxModule++;
        DvrModules[m].Index = DvrNetworkMaxModule;
    }

    for m = 0; DvrModules[m] != nil; m++  {
	if DvrModules[m].Type != DVR_NETWORK_MODULE {
	    continue;
	}

	module := DvrModules[m].Ctx;

	if module.CreateConf != nil {
	    rv = module.CreateConf(cf.Cycle);

	    if *(*string)(Pointer(uintptr(rv))) == DvrConfError {
	        return DvrConfError;
	    }
        }

	cf.Ctx[DvrModules[m].Index] = &rv;
    }

    cf.ModuleType = DVR_NETWORK_MODULE;
    cf.CommandType = DvrNetworkConf;
    rc := cf.ParseHandler(cf, nil);

    if rc != DvrConfOk {
	return rc;
    }

    for m = 0; DvrModules[m] != nil; m++  {
	if DvrModules[m].Type != DVR_NETWORK_MODULE {
	    continue;
	}

	module := DvrModules[m].Ctx;

	if module.InitConf != nil {
	    rv := module.InitConf(cf.Cycle, cf.Ctx[DvrModules[m].Index]);
	    if rv != DvrConfOk {
		return DvrConfError;
	    }
        }
    }
*/
    fmt.Println("Network Module Block Finish");

    return DvrConfOk;
}
