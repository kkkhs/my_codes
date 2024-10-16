ASSUME cs:code,ds:data,ss:stack ;头文件
data SEGMENT    ;数据段

data ENDS

stack SEGMENT    ;栈段

stack ENDS

code SEGMENT          ;代码段
  start:
        mov ax,data
        mov ds,ax


        mov ax,4c00H
        int 21H

code ENDS
end start