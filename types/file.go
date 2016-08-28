
/*
 * Copyright (C) 2016 Meng Shi
 */


package types


import (
    . "os"
)


type DvrFileType struct {
    Fd         *File;
    Name       *DvrStrType;
    Stat       *DvrFileStatType;

    Offset      int64;
    SysOffset   int64;

    Log        *DvrErrlogType;
    StatVaild   uint;
};
