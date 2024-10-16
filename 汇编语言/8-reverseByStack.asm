ASSUME cs:code,ds:data,ss:stack ;头文件
data SEGMENT                             ;数据段
  arr  DW 1111H,2222H,3333H,4444H,5555H
  res  db 5 dup (1)                      ;反转数组
data ENDS

stack SEGMENT           ;栈段
        db 100 DUP (0)
stack ENDS

code SEGMENT                ;代码段
  start:
        mov  ax,data
        mov  ds,ax
        mov  ax,stack
        mov  ss,ax

        mov  si,0
        mov  cx,5           ;cx计数器
  for:  
        mov  ax,ds:arr[si]
        push ax
        add  si,2           ;每次循环移动一个字节
        loop for

        mov  di,0
        mov  cx,5           ;cx计数器
  for2: 
        pop  ds:res[di]
        add  di,2
        loop for2

        mov  ax,4c00H
        int  21H

code ENDS
end start