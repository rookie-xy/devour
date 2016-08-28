/*
 * Copyright (C) 2016 Meng Shi
 */


package modules


import (
//    . "unsafe"

    . "devour/types"
    . "devour/http"
)


var DvrHttp = DvrStrType{ DvrSizeof("Http") - 1, "Http" };


var DvrHttpCommands = []DvrCommandType{

    { DvrHttp,
      DvrMainConf|DvrConfBlock|DvrConfNoargs,
      DvrHttpBlock,
      0,
      0,
      nil },

      DvrNilCommand,
};


var DvrHttpModuleCtx = DvrCoreModuleType{
    DvrHttp,
    nil,
    nil,
};


var DvrHttpModule = DvrMoudleType{
    0,
    0,
    &DvrHttpModuleCtx,
    DvrHttpCommands,
    DVR_CORE_MODULE,
    nil,
    nil,
};
