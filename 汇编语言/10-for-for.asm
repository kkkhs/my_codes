ASSUME DS:DATA,SS:STACK,CS:CODE
DATA SEGMENT
  str  DB 'aaaaabbbbbccccc',
       DB 'aaaaabbbbbccccc',
       DB 'aaaaabbbbbccccc',
       DB 'aaaaabbbbbccccc','$'
DATA ENDS

STACK SEGMENT
        DB 10 DUP (0)
STACK ENDS

CODE SEGMENT                    ;功能：将每一行前五个转大写
  start:
        mov  ax,data
        mov  ds,ax

        mov  bx,0
        mov  cx,4
  for1: 
        mov  dx,cx              ;存储外层循环的cx到dx
        mov  si,0               ;二重循环
        mov  cx,5
  for2: 
        mov  al,ds:str[bx][si]
        and  al,1011111B
        mov  ds:str[bx][si],al

        inc  si
        loop for2               ;二重循环结束
        mov  cx,dx              ;从dx取出外层循环的cx
        add  bx,16
        loop for1

  ;mov dx,offset str
        lea  dx,str
        mov  ah,9               ;将dx的内容显示到屏幕
        int  21H

        mov  ah,4cH
        int  21H

CODE ENDS
end start