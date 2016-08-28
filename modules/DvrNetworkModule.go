/*
 * Copyright (C) 2016 Meng Shi
 */


package modules


import (
//    . "unsafe"

    . "devour/types"
    . "devour/network"
)


var DvrNetwork = DvrStrType{ DvrSizeof("Network") - 1, "Network" };


var DvrNetworkCommands = []DvrCommandType{

    { DvrNetwork,
      DvrMainConf|DvrConfBlock|DvrConfNoargs,
      DvrNetworkBlock,
      0,
      0,
      nil },

      DvrNilCommand,
};


var DvrNetworkModuleCtx = DvrCoreModuleType{
    DvrNetwork,
    nil,
    nil,
};


var DvrNetworkModule = DvrMoudleType{
    0,
    0,
    &DvrNetworkModuleCtx,
    DvrNetworkCommands,
    DVR_CORE_MODULE,
    nil,
    nil,
};
