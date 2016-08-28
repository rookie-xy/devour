
/*
 * Copyright (C) 2016 Meng Shi
 */


package core


import (
    . "devour/types"
)


func DvrCreateArray(n int) *DvrArrayType {
    a := &DvrArrayType{};

    a.Elts = make([]*DvrStrType, n);
    a.Nelts = 0;

    return a;
}
