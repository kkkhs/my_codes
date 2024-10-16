ASSUME cs:code,ds:data,ss:stack ;头文件
data SEGMENT                 ;数据段
  arr  DW 1H,1H,100 dup (0)
data ENDS

stack SEGMENT           ;栈段
        db 100 DUP (0)
stack ENDS

code SEGMENT                  ;代码段
  start:
        mov  ax,data
        mov  ds,ax
        mov  ax,stack
        mov  ss,ax

        mov  bx,4             ;从第3个字节开始
        mov  cx,10            ;cx计数器
  for:  
        mov  dx,0
        add  dx,ds:arr[bx-2]
        add  dx,ds:arr[bx-4]
        mov  ds:arr[bx],dx

        add  bx,2             ;每次移动1个字节
        loop for

        mov  ax,4c00H
        int  21H

code ENDS
end start