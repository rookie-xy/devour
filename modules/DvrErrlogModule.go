/*
 * Copyright (C) 2016 Meng Shi
 */


package modules


import (
	"fmt"
//    . "unsafe"

    . "devour/types"
)


var DvrErrlogFile = DvrStrType{ DvrSizeof("ErrorLog") - 1, "ErrorLog" };


var DvrErrlogCommands = []DvrCommandType{

    { DvrErrlogFile,
      DvrMainConf|DvrConfMore1,
      DvrSetErrlog,
      0,
      0,
      nil },

      DvrNilCommand,
};


var DvrErrlogModuleCtx = DvrCoreModuleType{
    DvrErrlogFile,
    nil,
    nil,
};


var DvrErrlogModule = DvrMoudleType{
    0,
    0,
    &DvrErrlogModuleCtx,
    DvrErrlogCommands,
    DVR_CORE_MODULE,
    nil,
    nil,
};


func DvrSetErrlog(cf *DvrConfType, cmd *DvrCommandType, conf *DvrVoidType) string {
    fmt.Println("Conf Module Set Error log Finish");
    return DvrConfOk;
}
