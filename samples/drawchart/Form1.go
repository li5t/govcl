// 由GOVCL UI设计器自动生成，不要编辑
package main

import (
    "github.com/ying32/govcl/vcl"
)

type TForm1 struct {
    *vcl.TForm
    ListBox1 *vcl.TListBox

    //::private::
    TForm1Fields
}

var Form1 *TForm1




// 以字节形式加载
// vcl.Application.CreateForm(&Form1)

func NewForm1(owner vcl.IComponent) (root *TForm1)  {
    vcl.CreateResFrame(owner, &root)
    return
}

var (
    form1Bytes = []byte {
        0x54, 0x50, 0x46, 0x30, 0x05, 0x54, 0x46, 0x6F, 0x72, 0x6D, 0x05, 0x46, 
        0x6F, 0x72, 0x6D, 0x31, 0x04, 0x4C, 0x65, 0x66, 0x74, 0x02, 0x00, 0x03, 
        0x54, 0x6F, 0x70, 0x02, 0x00, 0x07, 0x43, 0x61, 0x70, 0x74, 0x69, 0x6F, 
        0x6E, 0x06, 0x31, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x20, 0x66, 0x6F, 
        0x72, 0x20, 0x47, 0x6F, 0x20, 0x2D, 0x3E, 0x20, 0x68, 0x74, 0x74, 0x70, 
        0x73, 0x3A, 0x2F, 0x2F, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2E, 0x63, 
        0x6F, 0x6D, 0x2F, 0x76, 0x64, 0x6F, 0x62, 0x6C, 0x65, 0x72, 0x2F, 0x63, 
        0x68, 0x61, 0x72, 0x74, 0x0C, 0x43, 0x6C, 0x69, 0x65, 0x6E, 0x74, 0x48, 
        0x65, 0x69, 0x67, 0x68, 0x74, 0x03, 0x82, 0x02, 0x0B, 0x43, 0x6C, 0x69, 
        0x65, 0x6E, 0x74, 0x57, 0x69, 0x64, 0x74, 0x68, 0x03, 0x37, 0x05, 0x05, 
        0x43, 0x6F, 0x6C, 0x6F, 0x72, 0x07, 0x09, 0x63, 0x6C, 0x42, 0x74, 0x6E, 
        0x46, 0x61, 0x63, 0x65, 0x0E, 0x44, 0x6F, 0x75, 0x62, 0x6C, 0x65, 0x42, 
        0x75, 0x66, 0x66, 0x65, 0x72, 0x65, 0x64, 0x09, 0x0C, 0x46, 0x6F, 0x6E, 
        0x74, 0x2E, 0x43, 0x68, 0x61, 0x72, 0x73, 0x65, 0x74, 0x07, 0x0F, 0x44, 
        0x45, 0x46, 0x41, 0x55, 0x4C, 0x54, 0x5F, 0x43, 0x48, 0x41, 0x52, 0x53, 
        0x45, 0x54, 0x0A, 0x46, 0x6F, 0x6E, 0x74, 0x2E, 0x43, 0x6F, 0x6C, 0x6F, 
        0x72, 0x07, 0x0C, 0x63, 0x6C, 0x57, 0x69, 0x6E, 0x64, 0x6F, 0x77, 0x54, 
        0x65, 0x78, 0x74, 0x0B, 0x46, 0x6F, 0x6E, 0x74, 0x2E, 0x48, 0x65, 0x69, 
        0x67, 0x68, 0x74, 0x02, 0xF3, 0x09, 0x46, 0x6F, 0x6E, 0x74, 0x2E, 0x4E, 
        0x61, 0x6D, 0x65, 0x06, 0x06, 0x54, 0x61, 0x68, 0x6F, 0x6D, 0x61, 0x0A, 
        0x46, 0x6F, 0x6E, 0x74, 0x2E, 0x53, 0x74, 0x79, 0x6C, 0x65, 0x0B, 0x00, 
        0x0E, 0x4F, 0x6C, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4F, 0x72, 
        0x64, 0x65, 0x72, 0x08, 0x08, 0x50, 0x6F, 0x73, 0x69, 0x74, 0x69, 0x6F, 
        0x6E, 0x07, 0x0A, 0x70, 0x6F, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6E, 0x65, 
        0x64, 0x06, 0x53, 0x63, 0x61, 0x6C, 0x65, 0x64, 0x08, 0x0D, 0x50, 0x69, 
        0x78, 0x65, 0x6C, 0x73, 0x50, 0x65, 0x72, 0x49, 0x6E, 0x63, 0x68, 0x02, 
        0x60, 0x0A, 0x54, 0x65, 0x78, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 
        0x02, 0x10, 0x00, 0x08, 0x54, 0x4C, 0x69, 0x73, 0x74, 0x42, 0x6F, 0x78, 
        0x08, 0x4C, 0x69, 0x73, 0x74, 0x42, 0x6F, 0x78, 0x31, 0x04, 0x4C, 0x65, 
        0x66, 0x74, 0x03, 0x6D, 0x04, 0x03, 0x54, 0x6F, 0x70, 0x02, 0x00, 0x05, 
        0x57, 0x69, 0x64, 0x74, 0x68, 0x03, 0xCA, 0x00, 0x06, 0x48, 0x65, 0x69, 
        0x67, 0x68, 0x74, 0x03, 0x82, 0x02, 0x05, 0x41, 0x6C, 0x69, 0x67, 0x6E, 
        0x07, 0x07, 0x61, 0x6C, 0x52, 0x69, 0x67, 0x68, 0x74, 0x0D, 0x49, 0x74, 
        0x65, 0x6D, 0x73, 0x2E, 0x53, 0x74, 0x72, 0x69, 0x6E, 0x67, 0x73, 0x01, 
        0x06, 0x09, 0x6B, 0x65, 0x79, 0x53, 0x74, 0x79, 0x6C, 0x65, 0x73, 0x06, 
        0x0B, 0x73, 0x63, 0x61, 0x74, 0x74, 0x65, 0x72, 0x54, 0x69, 0x63, 0x73, 
        0x06, 0x0C, 0x73, 0x63, 0x61, 0x74, 0x74, 0x65, 0x72, 0x43, 0x68, 0x61, 
        0x72, 0x74, 0x06, 0x0D, 0x66, 0x75, 0x6E, 0x63, 0x74, 0x69, 0x6F, 0x6E, 
        0x50, 0x6C, 0x6F, 0x74, 0x73, 0x06, 0x09, 0x61, 0x75, 0x74, 0x6F, 0x73, 
        0x63, 0x61, 0x6C, 0x65, 0x06, 0x08, 0x62, 0x6F, 0x78, 0x43, 0x68, 0x61, 
        0x72, 0x74, 0x06, 0x07, 0x6B, 0x65, 0x72, 0x6E, 0x65, 0x6C, 0x73, 0x06, 
        0x12, 0x68, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x68, 0x69, 
        0x73, 0x74, 0x46, 0x61, 0x6C, 0x73, 0x65, 0x06, 0x11, 0x68, 0x69, 0x73, 
        0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x68, 0x69, 0x73, 0x74, 0x54, 0x72, 
        0x75, 0x65, 0x06, 0x13, 0x68, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x72, 
        0x74, 0x73, 0x68, 0x69, 0x73, 0x74, 0x46, 0x61, 0x6C, 0x73, 0x65, 0x06, 
        0x12, 0x68, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x68, 
        0x69, 0x73, 0x74, 0x54, 0x72, 0x75, 0x65, 0x06, 0x13, 0x68, 0x69, 0x73, 
        0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x6F, 0x68, 0x69, 0x73, 0x74, 0x46, 
        0x61, 0x6C, 0x73, 0x65, 0x06, 0x12, 0x68, 0x69, 0x73, 0x74, 0x43, 0x68, 
        0x61, 0x72, 0x74, 0x6F, 0x68, 0x69, 0x73, 0x74, 0x54, 0x72, 0x75, 0x65, 
        0x06, 0x08, 0x62, 0x61, 0x72, 0x43, 0x68, 0x61, 0x72, 0x74, 0x06, 0x13, 
        0x63, 0x61, 0x74, 0x65, 0x67, 0x6F, 0x72, 0x69, 0x63, 0x61, 0x6C, 0x42, 
        0x61, 0x72, 0x43, 0x68, 0x61, 0x72, 0x74, 0x06, 0x07, 0x6C, 0x6F, 0x67, 
        0x41, 0x78, 0x69, 0x73, 0x06, 0x08, 0x70, 0x69, 0x65, 0x43, 0x68, 0x61, 
        0x72, 0x74, 0x06, 0x0C, 0x74, 0x65, 0x73, 0x74, 0x47, 0x72, 0x61, 0x70, 
        0x68, 0x69, 0x63, 0x73, 0x06, 0x06, 0x62, 0x65, 0x73, 0x74, 0x4F, 0x66, 
        0x06, 0x0B, 0x6D, 0x69, 0x65, 0x74, 0x65, 0x6E, 0x43, 0x68, 0x61, 0x72, 
        0x74, 0x00, 0x08, 0x54, 0x61, 0x62, 0x4F, 0x72, 0x64, 0x65, 0x72, 0x02, 
        0x00, 0x0C, 0x45, 0x78, 0x70, 0x6C, 0x69, 0x63, 0x69, 0x74, 0x4C, 0x65, 
        0x66, 0x74, 0x03, 0xB7, 0x02, 0x0B, 0x45, 0x78, 0x70, 0x6C, 0x69, 0x63, 
        0x69, 0x74, 0x54, 0x6F, 0x70, 0x02, 0x0A, 0x0E, 0x45, 0x78, 0x70, 0x6C, 
        0x69, 0x63, 0x69, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x03, 0x2D, 
        0x02, 0x00, 0x00, 0x00}
)

// 注册Form资源
var _ = vcl.RegisterFormResource(Form1, &form1Bytes)
