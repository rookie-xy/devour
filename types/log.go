
/*
 * Copyright (C) 2016 Meng Shi
 */


package types


var (
    DvrLogStderr  =  0;
    DvrLogEmerg   =  1;
    DvrLogAlert   =  2;
    DvrLogCrit    =  3;
    DvrLogErr     =  4;
    DvrLogWarn    =  5;
    DvrLogNotice  =  6;
    DvrLogInfo    =  7;
    DvrLogDebug   =  8;
)


type DvrErrlogHandleFunc func(ctx DvrVoidType, buf string, len int) string;

type DvrErrlogType struct {
    Level    int;
    File    *DvrOpenFileType;
    Data     DvrVoidType;
    Handler  DvrErrlogHandleFunc;
}
