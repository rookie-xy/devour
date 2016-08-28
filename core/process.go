
/*
 * Copyright (C) 2016 Meng Shi
 */


package core


import (
    . "os"

    . "devour/types"
)


const (
    DvrProcessSingle = 0;
    DvrProcessMaster = 1;
    DvrProcessWorker = 2;
)


var DvrPid int;
var DvrDaemonized int;
var DvrProcess int;


func DvrGetPid() int {
    return Getpid();
}


func DvrMasterProcessCycle(cycle *DvrCycleType) {
    return;
}


func DvrSingleProcessCycel(cycle *DvrCycleType) {
    return;
}
