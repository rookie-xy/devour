
/*
 * Copyright (C) 2016 Meng Shi
 */


package core


import (
//     "fmt"
   . "os"
   . "devour/types"
)


const (
    DvrFileInvalid  = -1;
    DvrFileError    = -1;
    DvrFileFlag     = O_CREATE|O_APPEND|O_RDWR;
    DvrFileOpenMode = 0660;
)


var (
    DvrOpenFileName = "os.Openfile()";
    DvrFdStatName   = "fd.Stat()";
    DvrCloseFileName= "os.Close()";
)


func DvrFdStat(fd *File, stat *DvrFileStatType) error {
    var err error;

    stat.FileInfo, err = fd.Stat();
    if err != nil {
	return err;
    }

    return nil;
}


func DvrFileSize(stat *DvrFileStatType) int64 {
    return stat.FileInfo.Size();
}


func DvrCloseFile() int {
    return 0;
}


func DvrReadFile(file *DvrFileType, buf []byte, size int, offset int64) int {
    var n int;

    if file.SysOffset != offset {
	if _, err := file.Fd.Seek(offset, 0); err != nil {
            DvrErrlogPrint(DvrLogCrit, file.Log, 0, "lseek() failed");
            return DvrError;
	}

	file.SysOffset = offset;
    }

    n, err := file.Fd.Read(buf);
    if err != nil {
        DvrErrlogPrint(DvrLogCrit, file.Log, 0, "Read() failed");
	return DvrError;
    }

    file.SysOffset += int64(n);
    file.Offset += int64(n);

    return n;
}
