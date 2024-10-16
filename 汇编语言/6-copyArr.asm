ASSUME cs:code,ds:data,ss:stack ;头文件
data SEGMENT                   ;数据段
  arr  db 1,2,3,4,10,20,30,40
  res  db 8 dup (1)            ;拷贝数组
data ENDS

stack SEGMENT           ;栈段
        db 100 DUP (0)
stack ENDS

code SEGMENT                ;代码段
  start:
        mov  ax,data
        mov  ds,ax

        mov  bx,0
        mov  cx,8           ;cx计数器
  for:  
        mov  al,ds:arr[bx]
        mov  ds:res[bx],al  ;res本身就是一个地址
        inc  bx
        loop for

        mov  ax,4c00H
        int  21H

code ENDS
end start