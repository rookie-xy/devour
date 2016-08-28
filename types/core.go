
/*
 * Copyright (C) 2016 Meng Shi
 */


package types


type DvrInitMoudleFunc func(cycle *DvrCycleType) uint;
type DvrInitProcessFunc func(cycle *DvrCycleType) uint;
type DvrCreateConfFunc func(cycle *DvrCycleType) DvrVoidType;
type DvrInitConfFunc func(cycle *DvrCycleType, conf *DvrVoidType) string;
type DvrSetFunc func(cf *DvrConfType, cmd *DvrCommandType, conf *DvrVoidType) string;


type DvrCommandType struct {
    Name    DvrStrType;
    Type    uint;
    Set     DvrSetFunc;
    Conf    int;
    Offset  uintptr;
    Post    DvrInterfType;
};


type DvrCoreModuleType struct {
    Name         DvrStrType;
    CreateConf   DvrCreateConfFunc;
    InitConf     DvrInitConfFunc;
};


type DvrMoudleType struct {
    CtxIndex      uint;
    Index         uint;
    Ctx          *DvrCoreModuleType;
    Commands      []DvrCommandType;
    Type          uint;
    InitMoudle   *DvrInitMoudleFunc;
    InitProcess  *DvrInitProcessFunc;
};
