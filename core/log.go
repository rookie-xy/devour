
/*
 * Copyright (C) 2016 Meng Shi
 */


package core


import (
    . "os"
    . "fmt"
      "time"

    . "devour/types"
);


var DvrStderr = &DvrOpenFileType{};
var DvrErrlog = DvrErrlogType{};
var DvrLogPid *int = &DvrPid;
var DvrLogTid int = 0;


var ErrLogLevel = [...]string{ "stderr", "emerg", "alert", "crit", "error",
                               "warn", "notice", "info", "debug" };


func DvrLogInitStderr() *DvrErrlogType {
    DvrStderr.File = Stdout;
    DvrErrlog.File = DvrStderr;
    DvrErrlog.Level = DvrLogErr;

    return &DvrErrlog;
}


func DvrErrlogPrint(level int, log *DvrErrlogType, err DvrErrType,
		                       fmt string, args ...interface{}) {

    if log.Level < level {
        return;
    }

    var errStr string;

    //if log.File.Fd == DvrFileInvalid {
    //    return;
    //}

    errStr += Sprintf("%s", time.Now().Format(Time[DvrDate]));
    errStr += Sprintf(" [%s] ", ErrLogLevel[level]);
    errStr += Sprintf("%d#%d: ", *DvrLogPid, DvrLogTid);

    if log.Data != nil {
        errStr += Sprintf("*%u ", log.Data);
    }

    errStr += Sprintf(fmt, args...);

    /* process system error
     * if err {
     *
     * } else {
     *
     * }
     */

    if log.Handler != nil {
        errStr += log.Handler(log.Data, errStr, len(errStr))
    }

    errStr += Sprintf("%c", LF);

    DvrErrlogWrite(log, []byte(errStr), len(errStr));
}


func DvrErrlogPrintCore(level int, log *DvrErrlogType, err DvrErrType,
		                           fmt string, args ...interface{}) {
}


func DvrErrlogWrite(log *DvrErrlogType, errstr []byte, len int) {
    n, err := log.File.Write(errstr);
    if len != n {
        if err != nil {
	    Println(err.Error());
        }

	Println("len error");
    }
}
