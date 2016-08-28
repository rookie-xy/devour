/*
 * Copyright (C) 2016 Meng Shi
 */


package modules


import (
	"fmt"
    . "unsafe"

    . "devour/types"
)


var DvrNetworkCore = DvrStrType{ DvrSizeof("Network") - 1, "Network" };


var (
    DvrConnections       = DvrStrType{ DvrSizeof("Connections") - 1,      "Connections"      };
    DvrMultiAccept       = DvrStrType{ DvrSizeof("MultiAccept") - 1,      "MultiAccept"      };
    DvrUse               = DvrStrType{ DvrSizeof("Use") - 1,              "Use"              };
    DvrAcceptMutex       = DvrStrType{ DvrSizeof("AcceptMutex") - 1,      "AcceptMutex"      };
    DvrAcceptMutexDelay  = DvrStrType{ DvrSizeof("AcceptMutexDelay") - 1, "AcceptMutexDelay" };
    DvrDebugConnection   = DvrStrType{ DvrSizeof("DebugConnection") - 1,  "DebugConnection"  };
)


var DvrNetworkCoreCommands = []DvrCommandType{

    { DvrConnections,
      DvrNetworkConf|DvrConfTake1,
      DvrNetworkConnections,
      0,
      0,
      nil },

    { DvrMultiAccept,
      DvrNetworkConf|DvrConfTake1,
      DvrConfSetFlagSlot,
      0,
      Offsetof(DvrNcf.MultiAccept),
      nil },

    { DvrUse,
      DvrNetworkConf|DvrConfTake1,
      DvrNetworkUse,
      0,
      0,
      nil },

    { DvrAcceptMutex,
      DvrNetworkConf|DvrConfTake1,
      DvrConfSetFlagSlot,
      0,
      Offsetof(DvrNcf.AcceptMutex),
      nil },

    { DvrAcceptMutexDelay,
      DvrNetworkConf|DvrConfTake1,
      DvrConfSetMsecSlot,
      0,
      Offsetof(DvrNcf.AcceptMutexDelay),
      nil },

    { DvrDebugConnection,
      DvrNetworkConf|DvrConfTake1,
      DvrNetworkDebugConnection,
      0,
      0,
      nil },

      DvrNilCommand,
};


var DvrNetworkCoreModuleCtx = DvrCoreModuleType{
    DvrNetworkCore,
    DvrNetworkCoreModuleCreateConf,
    DvrNetworkCoreModuleInitConf,
};


var DvrNetworkCoreModule = DvrMoudleType{
    0,
    0,
    &DvrNetworkCoreModuleCtx,
    DvrNetworkCoreCommands,
    DVR_NETWORK_MODULE,
    nil,
    nil,
};


func DvrNetworkCoreModuleCreateConf(cycle *DvrCycleType) DvrVoidType {
    ncf := DvrNcf;

    ncf.Connections = DvrConfUnsetUint;
    ncf.Use = DvrConfUnsetUint;
    ncf.MultiAccept = DvrConfUnsetBool;
    ncf.AcceptMutex = DvrConfUnsetBool;
    ncf.AcceptMutexDelay = DvrConfUnsetInt;
    ncf.Name = "";

    fmt.Println("Network Core Module Create Conf finish");
    return DvrVoidType(&ncf);
}


func DvrNetworkCoreModuleInitConf(cycle *DvrCycleType, conf *DvrVoidType) string {
fmt.Println("Network Core Module Init Conf finish");
    return DvrConfOk;
}


func DvrNetworkConnections(cf *DvrConfType, cmd *DvrCommandType, conf *DvrVoidType) string {
    fmt.Println("Netwrok Core Moudle Network Connections finish");
    return DvrConfOk;
}


func DvrNetworkUse(cf *DvrConfType, cmd *DvrCommandType, conf *DvrVoidType) string {
fmt.Println("Netwrok Core Moudle Network Use finish");
    return DvrConfOk;
}


func DvrNetworkDebugConnection(cf *DvrConfType, cmd *DvrCommandType, conf *DvrVoidType) string {
    name := cf.Args.Elts[1];
    fmt.Printf("Netwrok Core Moudle Network Debug Connection finish:%s\n", name.Data);
    return DvrConfOk;
}
