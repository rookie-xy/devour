
/*
 * Copyright (C) 2016 Meng Shi
 */


package core


import (
    . "devour/types"
)


func DvrCreateBuffer(size int) *DvrBufType {
    b := &DvrBufType{};

    if b == nil {
        return nil;
    }

    b.Start = 0;
    b.Pos = b.Start;
    b.Last = b.Start;
    b.End = b.Last + size;
    b.Buf = make([]byte, size);

    return b;
}
