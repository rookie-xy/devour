
/*
 * Copyright (C) 2016 Meng Shi
 */


package core


import (
	"fmt"
. "unsafe"
    . "devour/types"
    . "devour/autoconf"
)


var (
    DvrCycle  *DvrCycleType;
)

var DvrTestConfig bool;


func DvrInitCycle(oldCycle *DvrCycleType) *DvrCycleType {
    var i      uint;
    var cycle *DvrCycleType;
    var log   *DvrErrlogType;
    var rv     DvrVoidType;
    var conf   DvrConfType;

    log = oldCycle.Log;

    cycle = &DvrCycleType{}
    cycle.Log = log;
    cycle.OldCycle  = oldCycle;
    cycle.ConfFile  = oldCycle.ConfFile;
    cycle.Root = &DvrNilString;
    cycle.Root.Data = DvrPrefix;
    cycle.Root.Len  = len(DvrPrefix);
    cycle.ConfCtx = make([]*DvrVoidType, 1024);

    for i = 0; DvrModules[i] != nil; i++  {
	if DvrModules[i].Type != DVR_CORE_MODULE {
	    continue;
	}

	module := DvrModules[i].Ctx;

	if module.CreateConf != nil {
	    rv = module.CreateConf(cycle);

	    if *(*string)(Pointer(uintptr(rv))) == DvrConfError {
	        return nil;
	    }
        }

	cycle.ConfCtx[DvrModules[i].Index] = &rv;
    }

    conf = DvrConfType{};
    conf.Args = DvrCreateArray(10);

    conf.Ctx = cycle.ConfCtx;
    conf.Cycle = cycle;
    conf.Log = log;
    conf.ModuleType = DVR_CORE_MODULE;
    conf.CommandType = DvrMainConf;
    conf.Modules = DvrModules;
    conf.ParseHandler = DvrConfParse;

    if DvrConfParse(&conf, &cycle.ConfFile) != DvrConfOk {
	return nil;
    }

    fmt.Println("mengshimengshi");

    if DvrTestConfig {
        DvrErrlogPrint(DvrLogInfo, log, 0,
                       "the configuration file %s syntax is ok",
                       cycle.ConfFile.Data);
    }

    for i = 0; DvrModules[i] != nil; i++  {
	if DvrModules[i].Type != DVR_CORE_MODULE {
	    continue;
	}

	module := DvrModules[i].Ctx;

	if module.InitConf != nil {
	    if module.InitConf(cycle, cycle.ConfCtx[DvrModules[i].Index]) == DvrConfError {
		return nil;
	    }
        }
    }

    return cycle;
}
