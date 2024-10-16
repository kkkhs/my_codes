ASSUME DS:DATA,SS:STACK,CS:CODE
DATA SEGMENT
      arr  DB 32H,34H,47H,31H,1H,21H,3H,55H,87H
DATA ENDS

STACK SEGMENT
            DB 10 DUP (0)
STACK ENDS

CODE SEGMENT                          ;功能：冒泡排序
      start:
            mov  ax,data
            mov  ds,ax

            mov  bx,0
            mov  cx,8
      for1: 
            mov  dx,cx                ;存储外层循环的
            mov  si,8
            mov  cx,8
            sub  cx,bx
      for2: 
            mov  ah,ds:arr[si-1]
            mov  al,ds:arr[si]
            cmp  al,ah
            jnb  s
            mov  ds:arr[si-1],al      ;小于:交换
            mov  ds:arr[si],ah
      s:    
            dec  si
            loop for2                 ;二重循环结束
            mov  cx,dx                ;从dx取出外层循环的cx
            inc  bx
            loop for1

            mov  ax,4c00H
            int  21H

CODE ENDS
end start