
/*
 * Copyright (C) 2016 Meng Shi
 */


package modules


import (
      "fmt"
    . "unsafe"

    . "devour/types"
)


var (
    DvrCore            = DvrStrType{ DvrSizeof("Core") - 1,            "Core"            };
    DvrUser            = DvrStrType{ DvrSizeof("User") - 1,            "User"            };
    DvrDaemon          = DvrStrType{ DvrSizeof("Daemon") - 1,          "Daemon"          };
    DvrMasterProcess   = DvrStrType{ DvrSizeof("Master") - 1,          "Master"          };
    DvrWorkerProcesses = DvrStrType{ DvrSizeof("WorkerProcesses") - 1, "WorkerProcesses" };
    DvrPidFile         = DvrStrType{ DvrSizeof("Pid") - 1,             "Pid"             };
)


var DvrCoreCommands = []DvrCommandType{

    { DvrUser,
      DvrMainConf|DvrDirectConf|DvrConfTake12,
      DvrSetUser,
      0,
      0,
      nil },

    { DvrDaemon,
      DvrMainConf|DvrDirectConf|DvrConfTake1,
      DvrConfSetFlagSlot,
      0,
      Offsetof(DvrCcf.Daemon),
      nil },

    { DvrMasterProcess,
      DvrMainConf|DvrDirectConf|DvrConfTake1,
      DvrConfSetFlagSlot,
      0,
      Offsetof(DvrCcf.Master),
      nil },

    { DvrWorkerProcesses,
      DvrMainConf|DvrDirectConf|DvrConfTake1,
      DvrConfSetNumSlot,
      0,
      Offsetof(DvrCcf.WorkerProcesses),
      nil },

    { DvrPidFile,
      DvrMainConf|DvrDirectConf|DvrConfTake1,
      DvrConfSetStrSlot,
      0,
      Offsetof(DvrCcf.Pid),
      nil },

      DvrNilCommand,
};


var DvrCoreModuleCtx = DvrCoreModuleType{
    DvrCore,
    DvrCoreModuleCreateConf,
    DvrCoreModuleInitConf,
};


var DvrCoreModule = DvrMoudleType{
    0,
    0,
    &DvrCoreModuleCtx,
    DvrCoreCommands,
    DVR_CORE_MODULE,
    nil,
    nil,
};


func DvrCoreModuleCreateConf(cycle *DvrCycleType) DvrVoidType {
    ccf := DvrCoreConfType{};
    /*
    if ccf == nil {
        return DvrVoidType(&DvrConfError)
    }
    */

    ccf.Daemon = DvrConfUnsetBool;
    ccf.Master = DvrConfUnsetBool;
    ccf.WorkerProcesses = DvrConfUnsetInt;

    ccf.User = DvrConfUnsetInt;
    ccf.Group = DvrConfUnsetInt;

    fmt.Println("Core Module Create Config Finish");

    return DvrVoidType(&ccf);
}


func DvrCoreModuleInitConf(cycle *DvrCycleType, conf *DvrVoidType) string {
    fmt.Println("Core Module Init Config Finish");
    return DvrConfOk;
}


func DvrSetUser(cf *DvrConfType, cmd *DvrCommandType, conf *DvrVoidType) string {
    value := cf.Args.Elts[1];
    flag := string(value.Data.([]byte));
    fmt.Println("Core Module Set User Finish", flag);
    return DvrConfOk;
}
