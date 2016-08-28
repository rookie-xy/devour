
/*
 * Copyright (C) 2016 Meng Shi
 */


package types


type DvrCycleType struct {
    ConfFile   DvrStrType;
    Root      *DvrStrType;
    Log       *DvrErrlogType;
    OldCycle  *DvrCycleType;
    ConfCtx   []*DvrVoidType;
};


type DvrCoreConfType struct {
    Daemon           bool;
    Master           bool;
    WorkerProcesses  int;
    User             int;
    Group            int;
    Pid              DvrStrType;
    NewPid           DvrStrType;
};
