
/*
 * Copyright (C) 2016 Meng Shi
 */


package main


import (
      "os"
      "fmt"

    . "devour/types"
    . "devour/core"
    . "devour/autoconf"
    . "devour/modules"
)


var (
    DvrMaxModule  uint;
    DvrMaxSockets int;
)


func DvrGetOption(cycle *DvrCycleType, argc int, argv []string) int {
    var i int;

    for i = 1; i < argc; i++ {

	if argv[i][0] != '-' {
            DvrErrlogPrint(DvrLogEmerg, cycle.Log, 0,
                           "invalid option: \"%s\"", argv[i]);
	    return DvrError;
	}

        switch argv[i][1] {

        case 'c':
	    if argv[i + 1] == "" {
                DvrErrlogPrint(DvrLogEmerg, cycle.Log, 0,
                               "the option: \"%s\" requires file name",
			       argv[i]);

	        return DvrError;
	    }

	    cycle.ConfFile.Data = argv[i + 1];
	    cycle.ConfFile.Len  = len(cycle.ConfFile.Data.(string)) + 1;

            i++;

            break;

        case 't':
	    DvrTestConfig = true;
	    break;

        default:
            DvrErrlogPrint(DvrLogEmerg, cycle.Log, 0,
                           "invalid option: \"%s\"", argv[i]);
            break;
        }
    }

    if cycle.ConfFile.Data == nil {
        cycle.ConfFile.Data = DvrConfPath;
        cycle.ConfFile.Len  = len(DvrConfPath) + 1;
    }

    if DvrConfFullName(cycle, &cycle.ConfFile) == DvrError {
	return DvrError;
    }

    return DvrOk;
}


func DvrSaveOption(cycle *DvrCycleType, argc int, argv []string) int {
    return DvrOk;
}


func main() {
    var i           int;
    var initCycle   DvrCycleType;
    var cycle      *DvrCycleType;
    var log        *DvrErrlogType;
    var ccf        *DvrCoreConfType;

    DvrMaxSockets = -1;

    DvrTimeInit();

    DvrPid = DvrGetPid()

    if log = DvrLogInitStderr(); log == nil {
        return;
    }

    initCycle = DvrCycleType{};
    initCycle.Log = log;
    DvrCycle = &initCycle;

    argc := len(os.Args);

    if DvrGetOption(&initCycle, argc, os.Args) == DvrError {
        return;
    }

    if DvrSaveOption(&initCycle, argc, os.Args) == DvrError {
        return;
    }

    if DvrTestConfig {
	log.Level = DvrLogInfo;
    }

    if DvrOsInit(log) == DvrError {
	return;
    }

    DvrMaxModule = 0;
    for i = 0; DvrModules[i] != nil; i++  {
        DvrMaxModule++;
        DvrModules[i].Index = DvrMaxModule;
    }

    cycle = DvrInitCycle(&initCycle);
    if cycle == nil {
	if DvrTestConfig {
	    DvrErrlogPrint(DvrLogEmerg, log, 0,
	                   "the configure file %s test failed",
			   initCycle.ConfFile.Data);
	}

	return;
    }

    if DvrTestConfig {
        DvrErrlogPrint(DvrLogInfo, log, 0,
                       "the configure file %s was tested successfully",
		       initCycle.ConfFile.Data);
	return;
    }

    DvrOsStatus(cycle.Log);

    DvrCycle = cycle;

    ccf = (*DvrCoreConfType)(*DvrGetConf(cycle.ConfCtx, DvrCoreModule));

    if ccf.Master {
        DvrProcess = DvrProcessMaster;
    }

    fmt.Println(ccf.Daemon, ccf.Master, DvrCcf.Daemon);
    if ccf.Daemon {
	if DvrOsDaemon(log) == DvrError {
	    return;
	}

	DvrDaemonized = 1;
    }

    fmt.Printf("rrrrrrrrrrrrrrrrrrrrrrr:%d, %s\n", ccf.Pid.Len, ccf.Pid.Data);

    if DvrProcess == DvrProcessMaster {
	DvrMasterProcessCycle(cycle);

    } else {
	DvrSingleProcessCycel(cycle);
    }

    return;
}
